package animals

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
	Animal
}

func (a Animal) Move() int {
	return (a.speed * a.distance)
}

func (a Animal) Noise() string {
	return a.noise
}

func InitAllAnimals(distance int) (Animals, Animals) {
	var animal = Animal{speed: 10, distance: distance, noise: "meow"}
	var cat = Cat{purr: "purr", paws: "furry small", toeCount: 5, Animal: animal}

	animal.noise = "bark"
	animal.speed = 20
	var dog = Dog{wag: "wag the tail", paws: "furry big", toeCount: 4, Animal: animal}

	return cat, dog

}
