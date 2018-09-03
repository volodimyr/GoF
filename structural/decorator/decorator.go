package main

import (
	"log"
)

//Decorator structural pattern allows extending the function of an existing object dynamically without altering its internals.
//Decorators provide a flexible method to extend functionality of objects.

type Calculator func(float64, float64) float64

func LogDecorate(f func(float64, float64) float64) Calculator {
	return func(f1 float64, f2 float64) float64 {
		log.Println("Starting the execution")

		result := f(f1, f2)

		log.Println("Execution is completed with the result", result)

		return result
	}
}

func Multiply(f1, f2 float64) float64 {
	return f1 * f2
}

//usage	- ---- -- -- LogDecorate(Multiply)(3.6, 4.88)
