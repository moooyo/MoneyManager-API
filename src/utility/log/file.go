package log

import (
	"fmt"
	"os"
	"sync"
)

type FileLogger struct {
	File         *os.File
	lock         *sync.Mutex
	buffer       [4096]byte
	writeChannel chan ([4096]byte)
}

func (logger FileLogger) Init() {
	file, err := os.OpenFile("", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("Can not init file logger %s", err.Error()))
	}
	logger.File = file
	logger.lock = new(sync.Mutex)
}
