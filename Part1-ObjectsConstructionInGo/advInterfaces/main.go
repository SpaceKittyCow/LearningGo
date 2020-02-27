package main

import (
	"LearningGo/Part1-ObjectsConstructionInGo/advInterfaces/animals"
	"fmt"
)

func main() {
	catMove, DogMove, OrcaMove := AnimalsMovement()
	fmt.Printf("CatMove: %d, DogMove %d, OrcaMove %d", catMove, DogMove, OrcaMove)

	catPaws, DogPaws := PetsPaws()
	fmt.Printf(" CatPaws: %s, DogPaws %s", catPaws, DogPaws)

	OrcaFins := OrcaFins()
	fmt.Printf(" OrcaFins %d", OrcaFins)
}

func AnimalsMovement() (int, int, int) {
	cats, dogs, orca := animals.InitAllAnimals(5)
	return cats.Move(), dogs.Move(), orca.Move()
}

func PetsPaws() (string, string) {
	cats, dogs, _ := animals.InitAllAnimals(5)
	return cats.PawShape(), dogs.PawShape()
}

func OrcaFins() int {
	_, _, orca := animals.InitAllAnimals(5)
	return orca.FinCount()
}

func AnimalMovement(animal animals.Animals) int {
	return animal.Move()
}
