package main

import (
	"go.bug.st/serial"
)

func UploadBytes(config SerialConfig, bytes []byte) (int, error) {
	port, err := serial.Open(config.Port, &config.Settings)
	if err != nil {
		return 0, err 
	}

	return port.Write(bytes)
}
