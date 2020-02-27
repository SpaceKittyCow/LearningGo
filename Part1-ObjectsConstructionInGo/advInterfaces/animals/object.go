package animals

import "fmt"

type Animal struct {
	speed    int
	distance int
	noise    string
}

type Animals interface {
	Noise() string
	Move() int
}

type Cat struct {
	purr     string
	paws     string
	toeCount int
	Animal
}

type Dog struct {
	wag      string
	paws     string
	toeCount int
	//A        Animal //this will not work
	Animal
}

type Pets interface {
	Noise() string
	Move() int
	PawShape() string
}

type Orca struct {
	fins int
	Animal
}

type WaterAnimals interface {
	Noise() string
	Move() int
	FinCount() int
}

func (a Animal) Move() int {
	return (a.speed * a.distance)
}

func (a Animal) Noise() string {
	return a.noise
}

func (c Cat) Purr() {
	fmt.Println(c.purr)
}

func (c Cat) PawShape() string {
	return fmt.Sprintf("%s,  %d toes", c.paws, c.toeCount)
}

func (d Dog) PawShape() string {
	return fmt.Sprintf("%s,  %d toes", d.paws, d.toeCount)
}

func (d Dog) Wag() {
	fmt.Println(d.wag)
}

func (o Orca) FinCount() int {
	return o.fins
}

//func InitAllAnimals(distance int) (Cat, Dog, Orca) {
//func InitAllAnimals(distance int) (Animals, Animals, Animals) {
func InitAllAnimals(distance int) (Pets, Pets, WaterAnimals) {
	var animal = Animal{speed: 10, distance: distance, noise: "meow"}
	var cat = Cat{purr: "purr", paws: "furry small", toeCount: 5, Animal: animal}

	animal.noise = "bark"
	animal.speed = 20
	var dog = Dog{wag: "wag the tail", paws: "furry big", toeCount: 4, Animal: animal}

	var waterAnimal = Animal{speed: 40, distance: distance, noise: "splash"}
	var orca = &Orca{fins: 2, Animal: waterAnimal}
	return cat, dog, orca
}
