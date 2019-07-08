package main

import (
	"fmt"
	. "github.com/rezilience/go-trampoline/trampoline"
)

type (
	Integer struct {
		val int
	}
)

func sumBelowRec(args ...interface{}) Response {
	n := args[0].(Integer)
	var sum Integer
	if len(args) == 1 {
		sum = Integer{0}
	} else {
		sum = args[1].(Integer)
	}
	if n.val == 0 {
		return Response{
			Done:   true,
			Result: sum.val,
		}
	}
	return Response{
		Result: sum.val + n.val,
		Fn: func() Response {
			return sumBelowRec(Integer{n.val - 1}, Integer{sum.val + n.val})
		},
	}
}

func main() {
	// Example 1: sumBelow takes a number n and calculates the sum of all numbers from 1 to n recursively
	sumBelow := Trampoline(sumBelowRec)
	fmt.Println(sumBelow(Integer{10}))
}
