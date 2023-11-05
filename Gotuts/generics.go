package main

import (
	"fmt"
)

var pl = fmt.Println

type MyConstraint interface {
	int | float64
}

func getSumgen[T MyConstraint](x T, y T) T {
	return x + y
}

func main() {
	pl("5.1 + 4.5 = ", getSumgen(4.5, 5.1))
}
