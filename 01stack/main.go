package main

import "fmt"

func main() {
	stack := NewStack(3)

	pop, err := stack.Pop()
	fmt.Println(pop, err)

	err = stack.Push(1)
	fmt.Println(err)

	err = stack.Push(2)
	fmt.Println(err)

	err = stack.Push(3)
	fmt.Println(err)

	// this will fail as capacity is crossed
	err = stack.Push(4)
	fmt.Println(4, err)

	// will return 3
	peek, err := stack.Peek()
	fmt.Println(peek, err)

	// will return 3
	pop, err = stack.Pop()
	fmt.Println(pop, err)

	// this time push should work as capacity left is 1
	err = stack.Push(4)
	fmt.Println(4, err)

	// 1 > 2 > 4
	stack.Display()
}

// define the stack type with standard stack parameters

type Stack struct {
	// this is the underlying array which will store the elements of stack
	Arr []int
	// top : where is the topmost element
	Top int8
	// maximum capacity
	Size int8
}

func NewStack(size int8) *Stack {
	//return pointer to a newly created stack
	return &Stack{
		// allot array of input size
		Arr:  make([]int, size),
		Top:  -1,
		Size: size,
	}
}

// a check before pop, to safely pop

func (s *Stack) IsEmpty() bool {
	// in case stack is empty, the Top will be -1
	if s.Top < 0 {
		return true
	}
	return false
}

// a check before push, to safely push

func (s *Stack) IsFull() bool {
	// if the stack is totally filled, i.e. Top has reached last index
	// Top + 1 will be length of arr
	if s.Top+1 == s.Size {
		return true
	}
	return false
}

func (s *Stack) Push(e int) error {
	// safety, push only when there is space to push
	if s.IsFull() {
		return fmt.Errorf("stack full")
	}
	// increment top
	s.Top++
	// and insert element in array
	s.Arr[s.Top] = e
	return nil
}

func (s *Stack) Pop() (int, error) {
	// safety, pop only when there is something to pop
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack empty")
	}
	// get element at current position Top
	temp := s.Arr[s.Top]
	// reduce the current position of Top
	s.Top--
	return temp, nil
}

func (s *Stack) Peek() (int, error) {
	// safety, pop only when there is something to pop
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack empty")
	}
	// return the element at current position of Top, simply
	return s.Arr[s.Top], nil
}

func (s *Stack) Display() {
	fmt.Print("Stack Top Bottom : ")
	var i int8
	// print element top to bottom, from top element to last one in array
	for i = s.Top; i >= 0; i-- {
		fmt.Print(s.Arr[i], " > ")
	}
	fmt.Println()
}
