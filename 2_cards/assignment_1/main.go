package main

import "fmt"

func main() {
	sliceOfInts := []int{}

	for i := 0; i < 11; i++ {
		sliceOfInts = append(sliceOfInts, i)
	}

	for _, value  := range sliceOfInts {
		if value % 2 == 0 {
			fmt.Println(value, "is even")
		} else {
			fmt.Println(value, "is odd")
		}
	}
}