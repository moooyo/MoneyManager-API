package log

import (
	"sync"
	"time"
)

type LoggerBuffer[T any] struct {
	lock          *sync.Mutex
	buffer        []T
	bufferChannel chan ([]T)
	writeFunc     func([]T)
	maxBufferSize int
	flushInterval int
}

func (logger *LoggerBuffer[T]) Write(v T) {
	logger.lock.Lock()
	defer logger.lock.Unlock()

	if len(logger.buffer) >= logger.maxBufferSize {
		writedBuffer := logger.buffer
		logger.buffer = make([]T, 0, logger.maxBufferSize)
		logger.bufferChannel <- writedBuffer
	}

	logger.buffer = append(logger.buffer, v)
}

func (logger *LoggerBuffer[T]) Run() {
	go logger.autoFlush()
	go logger.autoWrite()
}

func (logger *LoggerBuffer[T]) autoFlush() {
	// auto flush buffer to buffer channel every 1 second
	go func() {
		task := func() {
			logger.lock.Lock()
			defer logger.lock.Unlock()
			if len(logger.buffer) == 0 {
				return
			}
			writedBuffer := logger.buffer
			logger.buffer = make([]T, 0, logger.maxBufferSize)

			logger.bufferChannel <- writedBuffer
		}

		for {
			time.Sleep(time.Duration(logger.flushInterval) * time.Second)
			task()
		}
	}()
}

func (logger *LoggerBuffer[T]) autoWrite() {
	// auto write buffer from buffer channel
	go func() {
		for {
			buffer := <-logger.bufferChannel
			logger.writeFunc(buffer)
		}
	}()
}
