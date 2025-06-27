package stack

import "fmt"

type Stack struct {
	items []rune
}

func (s *Stack) Push(data rune) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() {
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) Top() (rune, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Print() {
	for _, item := range s.items {
		fmt.Print(string(item), " ")
	}
	fmt.Println()
}

type IntStack struct {
	items []int
}

func (s *IntStack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *IntStack) Pop() {
	s.items = s.items[:len(s.items)-1]
}

func (s *IntStack) Top() int {
	return s.items[len(s.items)-1]
}

func (s *IntStack) IsEmpty() bool {
	return len(s.items) == 0
}
