package main

import "fmt"

func main() {
	// create new queue
	q := NewQueue(3)

	// no element there, dequeue will fail
	e, err := q.Dequeue()
	fmt.Println(e, err)

	// all 3 enqueue will pass
	err = q.Enqueue(1)
	fmt.Println(err)
	err = q.Enqueue(2)
	fmt.Println(err)
	err = q.Enqueue(3)
	fmt.Println(err)

	// max size reached
	err = q.Enqueue(4)
	fmt.Println(err)

	// dequeue will succeed
	e, err = q.Dequeue()
	fmt.Println(e, err)

	// enqueue will fail as queue is still not empty
	err = q.Enqueue(4)
	fmt.Println(err)

	// 2 more successful dequeue for clearing entire queue
	e, err = q.Dequeue()
	fmt.Println(e, err)

	e, err = q.Dequeue()
	fmt.Println(e, err)

	// dequeue will fail now
	e, err = q.Dequeue()
	fmt.Println(e, err)

	// enqueue will work as queue is reinitialized
	err = q.Enqueue(4)
	fmt.Println(err)
	err = q.Enqueue(5)
	fmt.Println(err)

	q.Display()
}

// define the queue type with standard queue parameters

type Queue struct {
	// this is the underlying array that will hold the elements of array
	Arr []int
	// Front: where is the first element that will be removed, when dequeued
	Front int8
	// Rear: where is the last element present currently,
	// after which position a new enqueue will insert the element
	Rear int8
	// max capacity
	Size int8
}

func NewQueue(size int8) *Queue {
	// return pointer to a newly created queue
	return &Queue{
		// allot array of input size
		Arr:   make([]int, size),
		Front: -1,
		Rear:  -1,
		Size:  size,
	}
}

// a check before dequeue to safely dequeue

func (q *Queue) IsEmpty() bool {
	// in case queue is empty, both rear and front will be -1
	return q.Front == -1
}

// a check before enqueue to safely enqueue

func (q *Queue) IsFull() bool {
	// if the queue has reached its capacity, rear would be last index
	// i.e. rear + 1 will be the last position, and hence, equal to length of underlying array
	return (q.Rear + 1) == q.Size
	// index to position conversion
}

func (q *Queue) Enqueue(e int) error {
	// safety, enqueue only if there is space to add new element
	if q.IsFull() {
		return fmt.Errorf("queue is full")
	}
	// this is special case,
	// if it's the first element to be inserted, then front must be 0 from -1,
	// to point the first available element's index to be dequeued
	if q.IsEmpty() {
		q.Front = 0
	}
	// insert at next index of rear
	q.Rear++
	q.Arr[q.Rear] = e
	return nil
}

func (q *Queue) Dequeue() (e int, err error) {
	// safety, dequeue only if there is any element left to be dequeued
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	// this is a special case,
	// if there would be only one element left in queue,
	// front and rear will point to same position
	// in that case, after removing that last element,
	// it would be safe to reinitialize the queue

	// also, this method is the only way,
	// queue gets all clear, after being empty once.
	if q.Front == q.Rear {
		// save this element, it will be sent at return
		e = q.Arr[q.Rear]
		// reinitialize queue
		q.Front, q.Rear = -1, -1
		// return
		return
	}
	e = q.Arr[q.Front]
	q.Front++
	return
}

func (q *Queue) Display() {

	if q.IsEmpty() {
		fmt.Println("Queue Empty: nothing to display")
	}
	// iterate from front to rear, to print all elements
	for i := q.Front; i <= q.Rear; i++ {
		fmt.Print(q.Arr[i], " > ")
	}
	fmt.Println()
}
