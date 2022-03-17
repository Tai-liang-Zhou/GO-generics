package main

import (
	"fmt"
)

type Number interface {
	int8 | int16 | int32 |int64 | float32 | float64
}

func newGenericFunc[age Number](myAge age){
	val := float32(myAge) + 1
	fmt.Println(val)
}


func BobbleSort[N Number](input []N) []N{
	n := len(input)
	swapped := true
	for swapped {
		swapped = false
		for i:= 0;i<n-1;i++{
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				swapped = true
			}
		}
	}
	return input
}

func main() {

	list := []int32{4,3,1,5,6}
	listfloat := []float32{4.3, 7.5, 2.4, 1.5}


	sorted := BobbleSort(list)
	fmt.Println(sorted)

	sorted2 := BobbleSort(listfloat)
	fmt.Println(sorted2)
}
