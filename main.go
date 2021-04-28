package main

import (
	"fmt"
	"strconv"
)

func main() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}

	value := evalRPN(tokens)

	fmt.Println(tokens)
	fmt.Println(value)
}

type polishStack struct {
	nums []int
}

func Constructor() polishStack {
	return polishStack{}
}
func (this *polishStack) Push(val string) {
	j, err := strconv.Atoi(val)
	if err == nil {
		this.nums = append(this.nums, j)
	}
}
func (this *polishStack) Pop() {
	this.nums = this.nums[:len(this.nums)-1]
}
func (this *polishStack) Top() int {
	return this.nums[len(this.nums)-1]
}
func (this *polishStack) IsEmpty() bool {
	return len(this.nums) == 0
}
func evalRPN(tokens []string) int {
	stack := new(polishStack)
	ret := 0
	for i := 0; i <= len(tokens)-1; i++ {
		//Case statement for math operators
		switch op := tokens[i]; op {
		case "+":
			val1 := stack.Top()
			stack.Pop()
			val2 := stack.Top()
			stack.Pop()
			ret = val2 + val1
			stack.Push(fmt.Sprint(ret))
		case "-":
			val1 := stack.Top()
			stack.Pop()
			val2 := stack.Top()
			stack.Pop()
			ret = val2 - val1
			stack.Push(fmt.Sprint(ret))
		case "*":
			val1 := stack.Top()
			stack.Pop()
			val2 := stack.Top()
			stack.Pop()
			ret = val2 * val1
			stack.Push(fmt.Sprint(ret))
		case "/":
			val1 := stack.Top()
			stack.Pop()
			val2 := stack.Top()
			stack.Pop()
			ret = val2 / val1
			stack.Push(fmt.Sprint(ret))
		}

		if tokens[i] != "+" || tokens[i] != "-" || tokens[i] != "*" || tokens[i] != "/" {
			stack.Push(tokens[i])
		}
	}
	return stack.Top()
}
