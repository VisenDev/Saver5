package main

import (
	"github.com/mitchellh/go-homedir"
	"go.bug.st/serial"
	"os"
	"time"
)

// this struct stores the configuration details that will be passed to the upload/download backend
type SerialConfig struct {
	Port     string
	Settings serial.Mode
}

// The model of MVC, stores internal state
type Model struct {
	Config           SerialConfig
	UploadFilepath   string
	DownloadFilepath string
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

// reads the value of the upload filepath
func (self *Model) ReadUploadFile() (string, error) {
	bytes, err := os.ReadFile(self.UploadFilepath)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
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
	}
}

//uploads the file in the model to the port in the model
func (m *Model) Upload() (int, error) {

	port, err := serial.Open(m.Config.Port, &m.Config.Settings)

	if err != nil {
		return 0, err
	}

	time.Sleep(10)

	//TODO close the port
	str, err := m.ReadUploadFile()
	if err != nil {
		return 0, err
	}
	
	return port.Write([]byte(str))
}
