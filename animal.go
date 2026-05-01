package main

import (
	"fmt"
)

type AnimalFarm struct {
	Animals []Animal
}

type Animal struct {
	weight   int
	consumes int
	produces int
}

func (boom AnimalFarm) TotalConsumes() int {
	total := 0
	for _, animal := range boom.Animals {
		total += animal.consumes
	}
	return total
}

func (boom AnimalFarm) TotalProduces() int {
	total := 0
	for _, animal := range boom.Animals {
		total += animal.produces
	}
	return total
}

func (boom AnimalFarm) TotalProfit() int {
	total := 0
	for _, animal := range boom.Animals {
		total += animal.produces - animal.consumes
	}
	return total
}

func (boom AnimalFarm) TotalWeight() int {
	total := 0
	for _, animal := range boom.Animals {
		total += animal.weight
	}
	return total
}

func SameWeight(boom AnimalFarm) map[int]int {
	// animals:=[]string{}
	var manimals = make(map[int]int)
	for i := 0; i < len(boom.Animals); i++ {
		for j := 0; j < len(boom.Animals); j++ {
			if boom.Animals[i].weight == boom.Animals[j].weight {
				manimals[i] = boom.Animals[i].weight
			}
		}
	}
	var canimals = make(map[int]int)
	for _, value := range manimals {
		count := 0
		for _, value2 := range manimals {
			if value == value2 {
				count++
				canimals[value] = count
			}
		}
	}

	return canimals
}

func main() {
	farm := AnimalFarm{
		Animals: []Animal{
			{weight: 100, consumes: 10, produces: 20},
			{weight: 150, consumes: 15, produces: 30},
			{weight: 100, consumes: 12, produces: 25},
			{weight: 100, consumes: 20, produces: 40},
			{weight: 250, consumes: 20, produces: 40},
			{weight: 250, consumes: 20, produces: 40},
			{weight: 150, consumes: 20, produces: 40},
		},
	}
	fmt.Println("Total consumes:", farm.TotalConsumes())
	fmt.Println("Total produces:", farm.TotalProduces())
	fmt.Println("Total profit:", farm.TotalProfit())
	fmt.Println("Total weight:", farm.TotalWeight())

	fmt.Println("Same weight animals:", SameWeight(farm))

}
