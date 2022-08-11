package stack

import "fmt"

type Stack struct {
	Array  []int
	Cursor int
}

func (s *Stack) StackInit(n int) []int {
	s.Array = make([]int, n)
	s.Cursor = 0
	fmt.Println(s.Array)
	return s.Array
}

func (s *Stack) Push(inputNumber int) []int {
	c := s.Cursor
	if c >= len(s.Array) {
		fmt.Println("stack is full")
	} else {
		s.Array[c] = inputNumber
	}
	s.PushCursor()
	fmt.Println("inputNumber", inputNumber)
	return s.Array
}

func (s *Stack) PushCursor() int {
	if len(s.Array) == s.Cursor {
		fmt.Println("stack is full")
	} else {
		s.Cursor++
	}
	return s.Cursor
}

func (s *Stack) PopCursor() int {
	if s.Cursor == 0 {
		fmt.Println("stack is empty")
	} else {
		s.Cursor--
	}
	return s.Cursor
}

func (s *Stack) Pop() []int {
	c := s.Cursor - 1
	s.Array[c] = 0
	s.PopCursor()
	fmt.Println("pop", c)
	return s.Array
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
