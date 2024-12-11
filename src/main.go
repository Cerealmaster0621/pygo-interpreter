// Purpose: demonstrate scanning/lexical analysis

package main

import (
	"fmt"
	"project/bytecodes/tokenizer"
	"project/tools/stack"
)

func main() {
	newStack := stack.NewStack[int]()
	fmt.Println(newStack)
	newStack.Push(1)
	fmt.Println(newStack)

	stack := stack.Stack[int]{}
	fmt.Println(stack)
	stack.Push(1)
	fmt.Println(stack)

	s := "あ아ㅑa"
	bt := []byte(s)
	fmt.Println(bt)
	fmt.Println(string(bt))

	var inputIntnumber = []byte("123 fdsaf 1323f 12321")
	var locIntnumber = tokenizer.IntnumberRegex.FindIndex(inputIntnumber)
	fmt.Println(string(inputIntnumber[locIntnumber[0]:locIntnumber[1]]))

}
