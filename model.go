package main

import (
	"github.com/mitchellh/go-homedir"
	"go.bug.st/serial"
	"os"
	"fyne.io/fyne/v2"
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
	ChosenPort		  serial.Port
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
func (m *Model) Upload(w *fyne.Window) (int, error) {

	//m.ChosenPort, err := serial.Open(m.Config.Port, &m.Config.Settings)
	if m.ChosenPort == nil {
		DisplayError(w, "Port is not open")
		return 0, errors.New("no open port");
	}

	//if err != nil {
	//	return 0, err
	//}

	//TODO close the port
	str, err := m.ReadUploadFile()
	if err != nil {
		return 0, err
	}
	
	return m.ChosenPort.Write([]byte(str))
}
