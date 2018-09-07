package behavioral

import "fmt"

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator
}

func (o *Operation) Operate(right, left int) int {
	return o.Apply(right, left)
}

type Addition struct{}

func (Addition) Apply(r, l int) int {
	return r + l
}

type Multiplication struct{}

func (Multiplication) Apply(r, l int) int {
	return r * l
}

func behavioral() {
	mult := Operation{Multiplication{}}

	fmt.Println(mult.Operate(3, 5)) // 15
}
