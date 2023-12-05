package main

import (
	"github.com/mitchellh/go-homedir"
	"go.bug.st/serial"
	"os"
	"errors"
)

// this struct stores the configuration details that will be passed to the upload/download backend
type SerialConfig struct {
	Port     string
	Settings serial.Mode
}

// The model of MVC, stores internal state
type Model struct {
	Config           SerialConfig
	Port		  serial.Port
	UploadFilepath   string
	UploadFileBuffer string
	DownloadFilepath string
	DownloadFileBuffer string
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
	err := os.WriteFile(self.DownloadFilepath, []byte(self.DownloadFileBuffer), 0)
	if err != nil {
		return  err
	}
	return nil;
}

// reads the value of the upload filepath
func (self *Model) ReadUploadFile() error {
	bytes, err := os.ReadFile(self.UploadFilepath)
	if err != nil {
		return  err
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
				Parity:   serial.EvenParity,
				DataBits: 8,
				StopBits: serial.OneStopBit,
			},
		},
		UploadFilepath:   "",
		DownloadFilepath: "",
		UploadFileBuffer: "",
		DownloadFileBuffer: "",
		Port: nil,
	}
}

//uploads the file in the model to the port in the model
func (m *Model) Upload() (int, error) {

	//m.Port, err := serial.Open(m.Config.Port, &m.Config.Settings)
	if m.Port == nil {
		return 0, errors.New("no open port");
	}

	err := m.ReadUploadFile()
	if err != nil {
		return 0, err
	}

	return m.Port.Write([]byte(m.UploadFileBuffer))
}
