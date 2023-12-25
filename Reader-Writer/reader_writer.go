package main

import (
	"context"
	"crypto/rand"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type IWriter interface {
	Write()
}

type IReader interface {
	Read()
}

var myArray = make([]byte, 0)
var wg sync.WaitGroup
var readerWg sync.WaitGroup
var writerWg sync.WaitGroup
var readerWriterMutex sync.RWMutex
var readerCountMutex sync.Mutex
var ctx context.Context
var cancel context.CancelFunc
var readerCount int

type Writer struct {
	writerId uint
}

func (w *Writer) Write() {
	defer writerWg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Warnln("time out for writing in array")
			return
		default:
			readerWriterMutex.Lock()
			newByte := make([]byte, 1)
			_, err := rand.Read(newByte)
			if err != nil {
				log.Errorf("error writing byte: %v", err)
				readerWriterMutex.Unlock()
				continue
			}
			myArray = append(myArray, newByte...)
			log.Infof("writer(%v) writes %v \n", w.writerId, newByte[0])
			readerWriterMutex.Unlock()

			// defining sleep for better visualization
			for newByte[0] < 5 && newByte[0] > 1 {
				rand.Read(newByte)
			}
			if newByte[0] != 0 {
				time.Sleep(time.Second * 1 / time.Duration(newByte[0]))

			}
		}
	}
}

type Reader struct {
	readerId uint
}

func (r Reader) Read() {
	defer readerWg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Warnln("time out for reading from array")
			return
		default:
			readerWriterMutex.RLock()

			readerCountMutex.Lock()
			readerCount++
			readerCountMutex.Unlock()

			readerCountMutex.Lock()
			log.Infof("reader count: %v, reader(%v), %v \n", readerCount, r.readerId, myArray)
			readerCountMutex.Unlock()

			readerCountMutex.Lock()
			if readerCount > 0 {
				readerCount--
			}
			readerCountMutex.Unlock()
			readerWriterMutex.RUnlock()

		}
	}
}

func main() {
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	writerFunc := func() {
		defer wg.Done()
		writerWg.Add(5)
		for i := 0; i < 5; i++ {
			writerInstance := Writer{writerId: uint(i)}
			go writerInstance.Write()
		}
		writerWg.Wait()
	}

	readerFunc := func() {
		defer wg.Done()
		readerWg.Add(3)
		for i := 0; i < 3; i++ {
			readerInstance := Reader{readerId: uint(i)}
			go readerInstance.Read()
		}
		readerWg.Wait()
	}

	wg.Add(2)
	go writerFunc()
	go readerFunc()
	wg.Wait()
}
