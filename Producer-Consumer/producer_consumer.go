package main

import (
	"context"
	"math/rand"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var item chan interface{}
var wg sync.WaitGroup
var ctx context.Context
var cancel context.CancelFunc

func Producer() {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Warn("Producer stopped")
			return
		default:
			sending_item := rand.Intn(100)
			item <- sending_item
			log.Infof("Produced item: %v", sending_item)
		}
		// Introduce a brief pause to make the production rate visible in the output
		time.Sleep(time.Millisecond * 5)
	}
}

func Consumer() {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Warn("Consumer stopped")
			return
		case received_item := <-item:
			log.Infof("Consumed item: %v", received_item)
		}
		// Introduce a longer pause to make the consumption process more visible
		time.Sleep(time.Millisecond * 700)
	}
}

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	item = make(chan interface{}, 5)
	ctx, cancel = context.WithCancel(context.Background())
	defer close(item)
	wg.Add(2)
	go Producer()
	go Consumer()
	wg.Wait()
	cancel()
}
