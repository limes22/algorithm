package stack

import "fmt"

type Stack struct {
	array  []int
	cursor int
}

func (s *Stack) StackInit(n int) []int {
	s.array = make([]int, n)
	s.cursor = 0
	fmt.Println(s.array)
	return s.array
}

func (s *Stack) Push(inputNumber int) []int {
	c := s.cursor
	if c >= len(s.array) {
		fmt.Println("stack is full")
	} else {
		s.array[c] = inputNumber
	}
	s.PushCursor()
	fmt.Println("inputNumber", inputNumber)
	return s.array
}

func (s *Stack) PushCursor() int {
	if len(s.array) == s.cursor {
		fmt.Println("stack is full")
	} else {
		s.cursor++
	}
	return s.cursor
}

func (s *Stack) PopCursor() int {
	if s.cursor == 0 {
		fmt.Println("stack is empty")
	} else {
		s.cursor--
	}
	return s.cursor
}

func (s *Stack) Pop() []int {
	c := s.cursor - 1
	s.array[c] = 0
	s.PopCursor()
	fmt.Println("pop", c)
	return s.array
}

func main() {
	var s Stack
	s.StackInit(3)
	fmt.Println(s.Push(3))
	fmt.Println(s.Push(4))
	fmt.Println(s.Push(5))
	fmt.Println(s.Push(4))
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
