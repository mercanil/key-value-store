package model

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
)

var Store = LockableMap{
	data: make(map[string]string),
}

type LockableMap struct {
	data map[string]string
	sync.RWMutex
}

func Put(key, val string) error {
	Store.Lock()
	defer Store.Unlock()
	Store.data[key] = val
	return nil
}

func Get(key string) (string, error) {
	val, ok := Store.data[key]
	if !ok {
		return "", ErrNoSuchKey
	}
	return val, nil

}

func Delete(key string) error {
	Store.Lock()
	defer Store.Unlock()
	delete(Store.data, key)
	return nil
}

var (
	ErrNoSuchKey = errors.New("No such a key")
)

type EventType byte

const (
	EventDelete EventType = 1
	EventPut    EventType = 2
)

type Event struct {
	Sequence  uint64
	EventType EventType
	Key       string
	Value     string
}

type TransactionLogger interface {
	WritePut(key, value string)
	WriteDelete(key string)
}

type FileTransactionLogger struct {
	events       chan Event
	errors       chan error
	lastSequence uint64
	file         *os.File
}

func (f FileTransactionLogger) WritePut(key, value string) {
	f.events <- Event{
		EventType: EventPut,
		Key:       key,
		Value:     value,
	}
}
func (f FileTransactionLogger) Err() chan error {
	return f.errors
}

func (f FileTransactionLogger) WriteDelete(key string) {
	f.events <- Event{
		EventType: EventDelete,
		Key:       key,
	}

}

func NewFileTransactionLogger(fileName string) (TransactionLogger, error) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal("Please give a valid file")
		return nil, errors.New("invalid file")
	}
	return &FileTransactionLogger{
		events:       make(chan Event, 16),
		errors:       make(chan error, 1),
		lastSequence: 0,
		file:         file,
	}, nil
}

func (f *FileTransactionLogger) Run() {
	go func() {
		for event := range f.events {
			f.lastSequence++
			fmt.Fprintf(f.file, "%d\t%d\t%s\t%s\n",
				f.lastSequence, event.EventType, event.Key, event.Value)
		}
	}()
}
