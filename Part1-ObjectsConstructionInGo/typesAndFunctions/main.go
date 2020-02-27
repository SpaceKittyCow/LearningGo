package main

import "fmt"

type Point struct{ x, y int } // Structs are a collection of types, more accurate description
func main() {
	var p1 = Point{1, 2} //initializing a type into a variable
	p2 := Point{2, 2}    //different way of init any variable
	fmt.Printf("%v", distanceBetween(p1, p2))
}

func distanceBetween(p1 Point, p2 Point) Point {
	return Point{p1.x - p2.x, p1.y - p2.y}
}
