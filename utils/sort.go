package utils

import "fmt"

func InsertionSort(input []int) int {
	fmt.Printf("sorting %v of length %v\n", input, len(input))
	var loop int = 0
	for i := 0 ; i < len(input); i++ {
		for j := i ; j < len(input) ; j++ {
			loop = loop + 1
			if input[j] < input [i] {
				temp := input[i]
				input[i] = input[j]
				input[j] = temp
			}
		}
	}
	return loop
}

/*
func main() {
	input := []int{5,4,3,2,1}
	loop := insertionSort(input)
	fmt.Printf("looped %v times %v\n", loop, input)
}*/
