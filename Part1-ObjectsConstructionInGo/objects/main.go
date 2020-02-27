package main

import "fmt"

type Animal struct {
	speed    int
	distance int
	noise    string
}
type Animal2 struct {
	speed    int
	distance int
	noise    string
}
type Cat struct {
	purr string
	Animal
	Animal2
}

func (a Animal) Noise() {
	fmt.Printf(a.noise)
}

func (a Animal2) Move() {
	fmt.Printf("%d", (a.speed * a.distance))
}

func (c Cat) Purr() {
	fmt.Printf(c.purr)
}
func main() {
	var animal = Animal{speed: 10, distance: 4, noise: "meow"}
	var cat = Cat{}
	cat.purr = "purr"
	cat.Animal = animal
	cat.Purr()         //Output: purr
	cat.Noise()        //Output: meow
	cat.Animal.Noise() //Output: meow
	animal.Noise()     //Output: noise
}
