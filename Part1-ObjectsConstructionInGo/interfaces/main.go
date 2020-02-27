package main

import (
	"fmt"
	"testingWithGo/Part1-ObjectsConstructionInGo/interfaces/animals"
)

func main() {
	cats, dogs := animals.InitAllAnimals(5)
	catMove, DogMove := AnimalsMovement(cats, dogs)
	fmt.Printf("CatMove: %d, DogMove %d ", catMove, DogMove)
}

func AnimalsMovement(cats animals.Animals,dogs animals.Animals) (int, int) {
	return cats.Move(), dogs.Move()
}

