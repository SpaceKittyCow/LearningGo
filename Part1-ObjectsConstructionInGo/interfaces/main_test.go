package main

import "testing"

type AnimalMock struct {
        move     int
        noise    string
}

type AnimalsMock interface {
        Noise() string
        Move() int
}

func (a AnimalMock) Move() int {
        return a.move
}

func (a AnimalMock) Noise() string {
        return a.noise
}
func TestAnimalsMovement(t *testing.T){
	animalMock := AnimalMock{move:10, noise:"test"}	
	test1, test2 := AnimalsMovement(animalMock, animalMock)
	if test1 != 10 {
		t.Error("test 1 didn't return 5")
	}
	
	if test2 != 10 {
		t.Error("test 1 didn't return 5")
	}
}
