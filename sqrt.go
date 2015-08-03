package main

import "fmt"

const Delta = 0.001

func nextSqrt(cur float64, x int) float64 {
	return (cur - ((cur * cur) - float64(x)) / (2 * cur))
}

func sqrtWithTries(x int, start float64, tries int) float64 {
	var next float64 = float64(start)
	for i := 0 ; i < tries ; i++ {
		next = nextSqrt(next, x)
	}
	return next
}

func chooseStart(x int) float64 {
	return float64(x)/2
}

func sqrt(x int) float64 {
	var cur float64 = sqrtWithTries(x,chooseStart(x), 5) // do 5 minimum tries
	var next float64 = nextSqrt(cur, x)
	for {
		fmt.Printf(".")
		if next - cur  < Delta {
			return next
		}
		cur = next
		next = nextSqrt(cur, x)
	}

}

func main() {
	sqrtWithTries(10,3,10)
	fmt.Printf("%v",sqrt(10))
	//sqrt(10)
}


