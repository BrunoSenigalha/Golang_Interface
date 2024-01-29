package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	base  float64
	heigh float64
}

type square struct {
	sideLenght float64
}

var option string

func main() {

	fmt.Println("Enter a kind of shape (triangle/square)")
	fmt.Scanln(&option)

	switch option {
	case "triangle":
		t := triangle{}
		fmt.Println("Triangle:")
		fmt.Print("Enter a base value: ")
		fmt.Scanf("%f\n", &t.base)
		fmt.Print("Enter a heigh value: ")
		fmt.Scan(&t.heigh)
		printArea(t)

	case "square":
		s := square{}
		fmt.Println("Square:")
		fmt.Print("Enter a side lenght:")
		fmt.Scan(&s.sideLenght)
		printArea(s)

	default:
		fmt.Println("There were something wrong")
	}
}

func printArea(s shape) {
	fmt.Println("The area is", s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.heigh
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLenght, 2)
}
