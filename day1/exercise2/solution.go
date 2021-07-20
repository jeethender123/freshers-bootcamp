package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func IsOperator(c uint8) bool {
	return strings.ContainsAny(string(c), "+ & - & * & /")
}

func IsOperand(c uint8) bool {
	return c >= '0' && c <= '9'
}

func GetOperatorWeight(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return -1
}

func HasHigherPrecedence(op1 string, op2 string) bool {
	op1Weight := GetOperatorWeight(op1)
	op2Weight := GetOperatorWeight(op2)
	return op1Weight >= op2Weight
}

func ToPostfix(s string) string {

	var stack Stack

	postfix := ""

	length := len(s)

	for i := 0; i < length; i++ {

		char := string(s[i])
		// Skip whitespaces
		if char == " " {
			continue
		}

		if char == "(" {
			stack.Push(char)
		} else if char == ")" {
			for !stack.Empty() {
				str, _ := stack.Top().(string)
				if str == "(" {
					break
				}
				postfix += str
				stack.Pop()
			}
			stack.Pop()
		} else if !IsOperator(s[i]) {
			// If character is not an operator
			// Keep in mind it's just an operand
			j := i
			number := ""
			for ; j < length && IsOperand(s[j]); j++ {
				number = number + string(s[j])
			}
			postfix +=  number
			i = j - 1
		} else {
			// If character is operator, pop two elements from stack,
			// perform operation and push the result back.
			for !stack.Empty() {
				top, _ := stack.Top().(string)
				if top == "(" || !HasHigherPrecedence(top, char) {
					break
				}
				postfix +=  top
				stack.Pop()
			}
			stack.Push(char)
		}
	}

	for !stack.Empty() {
		str, _ := stack.Pop().(string)
		postfix +=  str
	}

	return strings.TrimSpace(postfix)
}

func ToPrefix(s string) string{
	var stack1 Stack
	//prefix := ""
	length := len(s)

	for i :=0;i<length ;i++{
		if IsOperator(s[i]){
			op1 := stack1.Top().(string)
			stack1.Pop()
			op2 := stack1.Top().(string)
			stack1.Pop()

			temp := string(s[i])+" "+ op2 +" "+ op1
			stack1.Push(temp)

		}else {
			//fmt.Print("hello")
			stack1.Push(string(s[i]))
		}
	}

	ans := ""
	for !stack1.Empty(){
		ans += stack1.Top().(string)
		stack1.Pop()
	}

	return ans
}


func ReadFromInput() (string, error) {

	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')

	return strings.TrimSpace(s), err
}

func main() {

	fmt.Print("Enter infix expression: ")
	infixString, err := ReadFromInput()

	if err != nil {
		fmt.Println("Error when scanning input:", err.Error())
		return
	}
	var poststring = ToPostfix(infixString)
	fmt.Println("This is your postfix notation:", ToPostfix(infixString))
	fmt.Println("This is your postfix notation:", ToPrefix(poststring))



	return
}

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{}
	next  *Element
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Top() interface{} {
	return s.top.value
}

func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}