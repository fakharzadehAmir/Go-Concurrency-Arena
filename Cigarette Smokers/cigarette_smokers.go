package main

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
)

type IngredientType int

const (
	TobaccoType IngredientType = iota
	MatchesType
	PaperType
)

var (
	table []IngredientType
	cond  *sync.Cond
	wg    sync.WaitGroup
)

type Smoker struct {
	id      int
	ingType IngredientType
}

func printIngredient(ingredient IngredientType) string {
	switch ingredient {
	case TobaccoType:
		return "Tobacco"
	case PaperType:
		return "Paper"
	default:
		return "Matches"
	}
}

func agencyTable() {
	defer wg.Done()
	ingredientRotation := -1
	for ingredientRotation < 2 {
		cond.L.Lock()
		ingredientRotation = (ingredientRotation + 1) % 3
		table = []IngredientType{
			IngredientType((ingredientRotation + 1) % 3),
			IngredientType((ingredientRotation + 2) % 3),
		}
		log.Printf("Agent places %s and %s on the table.\n",
			printIngredient(table[0]), printIngredient(table[1]))
		cond.L.Unlock()
		cond.Broadcast()
		// Introduce a brief pause to make the reloading table and making cigarette visible in the output
		time.Sleep(time.Second)
	}
}

func makeCigarette(smoker Smoker) {
	defer wg.Done()
	cond.L.Lock()
	for len(table) != 2 || (len(table) == 2 && slices.Contains(table, smoker.ingType)) {
		cond.Wait()
	}
	log.Infof("Smoker %d with %s makes a cigarette.\n", smoker.id, printIngredient(smoker.ingType))
	table = nil
	cond.L.Unlock()

}

func main() {
	cond = sync.NewCond(&sync.Mutex{})

	wg.Add(4)
	go agencyTable()
	for i := 0; i < 3; i++ {
		go makeCigarette(Smoker{id: i, ingType: IngredientType(i % 3)})

	}
	wg.Wait()

}
