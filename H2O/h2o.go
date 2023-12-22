package main

import (
	"context"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var wg sync.WaitGroup
var ctx context.Context
var cancel context.CancelFunc

type H2O struct {
	sync.Mutex
	oxygenCount   int
	hydrogenCount int
}

func (h2o *H2O) ProduceHydrogen() {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Warnln("Hydrogen production stopped")
			return
		default:
			h2o.Lock()
			h2o.hydrogenCount++
			log.Infoln("Hydrogen has been built")
			h2o.Unlock()
			// Introduce a brief pause to make the production rate visible in the output
			time.Sleep(time.Second * 1 / 50)
		}
	}
}

func (h2o *H2O) ProduceOxygen() {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Warnln("Oxygen production stopped")
			return
		default:
			h2o.Lock()
			h2o.oxygenCount++
			if h2o.hydrogenCount >= 2 && h2o.oxygenCount >= 1 {
				log.Infoln("H2O was built!")
				h2o.hydrogenCount -= 2
				h2o.oxygenCount -= 1
			}
			h2o.Unlock()
		}
	}
}

func main() {
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	h2o := H2O{}

	wg.Add(2)
	go h2o.ProduceHydrogen()
	go h2o.ProduceOxygen()
	wg.Wait()
}
