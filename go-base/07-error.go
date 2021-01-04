package main

import (
	"fmt"
	"errors"
	"math"
)

func doSqrt(x float64) (float64, error){
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	} else {
		return math.Sqrt(x), nil
	}
}

func main(){
	result, error := doSqrt(-8)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(result)
	}
}

