package main

import (
	"errors"
	"fmt"
	"os"
	//"time"

	"github.com/mitchellh/go-homedir"
	"go.bug.st/serial"
)

// this struct stores the configuration details that will be passed to the upload/download backend
type SerialConfig struct {
	Port     string
	Settings serial.Mode
}

// The model of MVC, stores internal state
type Model struct {
	Config             SerialConfig
	Port               serial.Port
	UploadFilepath     string
	UploadFileBuffer   string
	DownloadFilepath   string
	DownloadFileBuffer []byte
	DownloadBufLen     int
}

// sets the upload filepath
func (self *Model) SetUploadFilepath(path string) error {
	filepath, err := homedir.Expand(path)
	if err != nil {
		return err
	}

	self.UploadFilepath = filepath
	return nil
}

func (self *Model) SetDownloadFilepath(path string) error {
	filepath, err := homedir.Expand(path)
	if err != nil {
		return err
	}

	self.DownloadFilepath = filepath
	return nil
}

func (self *Model) WriteDownloadFile() error {
	err := os.WriteFile(self.DownloadFilepath, self.DownloadFileBuffer[0:self.DownloadBufLen], os.ModePerm)
	if err != nil {
		return err
	}
	//os.Chmod()
	return nil
}

// reads the value of the upload filepath
func (self *Model) ReadUploadFile() error {
	bytes, err := os.ReadFile(self.UploadFilepath)
	if err != nil {
		return err
	}
	self.UploadFileBuffer = string(bytes)
	return nil
}

// initalizes with proper defaults
func DefaultModel() Model {
	return Model{
		Config: SerialConfig{
			Port: "",
			Settings: serial.Mode{
				BaudRate: 4800,
				//BaudRate: 230400,
				Parity:   serial.EvenParity,
				DataBits: 8,
				StopBits: serial.OneStopBit,
			},
		},
		UploadFilepath:     "",
		DownloadFilepath:   "",
		UploadFileBuffer:   "",
		DownloadFileBuffer: make([]byte, 100000),
		DownloadBufLen:     0,
		Port:               nil,
	}
}

// uploads the file in the model to the port in the model
func (m *Model) Upload() (int, error) {
	// m.Port, err := serial.Open(m.Config.Port, &m.Config.Settings)
	if m.Port == nil {
		return 0, errors.New("no open port")
	}

	err := m.ReadUploadFile()
	if err != nil {
		return 0, err
	}

	return m.Port.Write([]byte(m.UploadFileBuffer))
}

func (m *Model) Listen(callback func()) {
	go func() {
		for {
			if m.Port == nil {
				//fmt.Println("port is nil")
				//time.Sleep(1 * time.Second)
			} else {
				
				//t, _ := time.ParseDuration("500ms")
				//_ = m.Port.SetReadTimeout(t)
				
				//m.Port.Close()
				//var err error
				//m.Port, err = serial.Open(m.Config.Port, &m.Config.Settings)	
				//if err != nil {
					//continue
				//}
				
				//fmt.Println(m.Config.Port)
				n, err := m.Port.Read(m.DownloadFileBuffer)
				
				//fmt.Print("read in lines: ")
				//fmt.Println(n)
				if err == nil && n > 0 {
					m.DownloadBufLen += n
					callback()
				} else if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()
}

func (m *Model) Save() (int, error) {
	// m.Port, err := serial.Open(m.Config.Port, &m.Config.Settings)
	if m.Port == nil {
		return 0, errors.New("no open port")
	}

	err := m.ReadUploadFile()
	if err != nil {
		return 0, err
	}

	return m.Port.Write([]byte(m.UploadFileBuffer))
}
