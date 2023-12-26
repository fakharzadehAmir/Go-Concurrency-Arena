package main

import (
	"context"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup
var chopsticks [5]chan interface{}

func pickupChopstick(ctx context.Context, philosopher int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			log.Printf("philosopher (%v) is thinking\n", philosopher)
			time.Sleep(time.Second)

			if philosopher%2 == 0 {
				<-chopsticks[philosopher]
				<-chopsticks[(philosopher+1)%5]
			} else {
				<-chopsticks[(philosopher+1)%5]
				<-chopsticks[philosopher]
			}

			log.Printf("philosopher (%v) is eating\n", philosopher)
			time.Sleep(time.Second)

			if philosopher%2 == 0 {
				chopsticks[(philosopher+1)%5] <- struct{}{}
				chopsticks[philosopher] <- struct{}{}
			} else {
				chopsticks[philosopher] <- struct{}{}
				chopsticks[(philosopher+1)%5] <- struct{}{}
			}

			log.Printf("philosopher (%v) finished eating\n", philosopher)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for i := 0; i < 5; i++ {
		chopsticks[i] = make(chan interface{}, 1)
		chopsticks[i] <- struct{}{}
		defer close(chopsticks[i])
	}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go pickupChopstick(ctx, i)
	}

	wg.Wait()
}
