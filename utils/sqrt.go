// file to simple square root calculation, start with half of number while calculating square root using newton's formula
package utils

import "fmt"

const Delta = 0.001

func nextSqrt(cur float64, x int) float64 {
	return (cur - ((cur*cur)-float64(x))/(2*cur))
}

func sqrtWithTries(x int, start float64, tries int) float64 {
	var next float64 = float64(start)
	for i := 0; i < tries; i++ {
		next = nextSqrt(next, x)
	}
	return next
}

func chooseStart(x int) float64 {
	return float64(x) / 2
}

func Sqrt(x int) float64 {
	var cur float64 = sqrtWithTries(x, chooseStart(x), 5) // do 5 minimum tries
	var next float64 = nextSqrt(cur, x)
	for {
		fmt.Printf(".\n")
		if next-cur < Delta {
			return next
		}
		cur = next
		next = nextSqrt(cur, x)
	}

}

/*
func main() {
	var val int
	args, err := fmt.Scanln(&val)
	if(err != nil) {
		fmt.Printf("%v\n",err)
	} else {
		fmt.Printf("read %v args\n", args)
	}
	fmt.Printf("Calculating Square Root of %v\n", val)
	fmt.Printf("sqrt of %v = %v\n", val, sqrt(val))
	
}*/
