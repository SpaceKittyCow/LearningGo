package main // Package Name, main is default, can import others

import "fmt"

var point1 = [2]int{2, 3} // x, y (array of int (statically assigned to 2))
var point2 = []int{4, 5}  // x, y (slice of an array of int (not static amount, can have n amount of data))

//Main function is function that is run when run go run or the built executable  i
func main() {
	fmt.Printf("%v", DistanceBetween()) //prints file out into
} 

// DistanceBetween calculates the distance between two graph points x and y
func DistanceBetween() (x, y) {
	return (point1[0] - point2[0]), (point1[1] - point2[1]) //Can return multiple variables inline
}
