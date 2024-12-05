// Purpose: demonstrate scanning/lexical analysis

package main

import (
	"fmt"

	"project/stack"
)

func main() {
	var intStack stack.Stack[int]
	intStack.Push(1)
	fmt.Println(intStack)
	intStack.Push(2)
	fmt.Println(intStack)
	intStack.Pop()
	fmt.Println(intStack)
	intStack.Push(3)
	fmt.Println(intStack)
	intStack.Clear()
	fmt.Println(intStack)
}
