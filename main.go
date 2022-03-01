package main

import "fmt"

func main() {
	var names [3]string

	var distances [5]int

	var data [5]uint8

	var ratios [1]float64

	var alives [4]bool

	var zero [0]uint8

	fmt.Printf("names : %#v\n", names)
	fmt.Printf("distances: %#v\n", distances)
	fmt.Printf("data: %#v\n", data)
	fmt.Printf("ratios: %#.2v\n", ratios)
	fmt.Printf("alives: %#v\n", alives)
	fmt.Printf("zero: %#v\n", zero)

}
