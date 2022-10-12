package main

import (
	"fmt"

	"github.com/alurm/exhaustive-go/isf"
)

func do(it isf.T) {
	isf.Match(it,
		func(it int) {
			fmt.Println("Int:", it)
		},
		func(it string) {
			fmt.Println("String:", it)
		},
		func(it float64) {
			fmt.Println("Float64:", it)
		},
	)
}

func main() {
	var it isf.T
	it = isf.Int(33)
	do(it)
	it = isf.Float64(15.6)
	do(it)
	it = isf.String("Hello")
	do(it)
}
