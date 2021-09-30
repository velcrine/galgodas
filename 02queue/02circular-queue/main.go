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

	// enqueue will run successfully as queue is still not empty
	err = q.Enqueue(4)
	fmt.Println(err)

	// 4, 2, 3
	q.Display()

	// enqueue will fail as rear +1 == front
	err = q.Enqueue(4)
	fmt.Println(err)

	// 2 dequeue to remove 2 and 3
	e, err = q.Dequeue()
	fmt.Println(e, err)
	e, err = q.Dequeue()
	fmt.Println(e, err)

	// enqueue will work as queue is reinitialized
	err = q.Enqueue(5)
	fmt.Println(err)
	err = q.Enqueue(6)
	fmt.Println(err)

	q.Display()
}

// define the queue type with standard queue parameters

type CircularQueue struct {
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

func NewQueue(size int8) *CircularQueue {
	// return pointer to a newly created queue
	return &CircularQueue{
		// allot array of input size
		Arr:   make([]int, size),
		Front: -1,
		Rear:  -1,
		Size:  size,
	}
}

// a check before dequeue to safely dequeue

func (q *CircularQueue) IsEmpty() bool {
	// in case queue is empty, both rear and front will be -1
	return q.Front == -1
}

// a check before enqueue to safely enqueue

func (q *CircularQueue) IsFull() bool {
	// circular queue is full when, rear comes before front
	//  1  2  3  4  5  6  7  8
	//        |  |
	//     rear  front
	// queue is full as rear + 1 will be front
	// because element at front couldn't be overwritten

	// that's because in circular queue, rear will go back again at index 0
	// if it hits capacity, like below
	//  1  2  3  4  5  6  7  8
	//  |                    |
	//  front                rear
	// now also queue is full
	// explaination:
	// rear will be 7; 7 +1 % 8 will be 0, which is index of front

	return (q.Rear+1)%q.Size == q.Front
}

func (q *CircularQueue) Enqueue(e int) error {
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
	// take care of handling the case
	//  1  2  3  4  5  6  7  8
	//                       |
	//                      rear

	// as this is circular queue, rear can go back to index 0
	q.Rear = (q.Rear + 1) % q.Size
	q.Arr[q.Rear] = e
	return nil
}

func (q *CircularQueue) Dequeue() (e int, err error) {
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
	// take care of handling the case
	//  1  2  3  4  5  6  7  8
	//                       |
	//                      front

	// as this is circular queue, front can go back to index 0
	q.Front = (q.Front + 1) % q.Size
	return
}

func (q *CircularQueue) Display() {

	if q.IsEmpty() {
		fmt.Println("CircularQueue Empty: nothing to display")
	}
	// iterate from front to rear, to print all elements

	// special case in circular queue
	//  1  2  3  4  5  6  7  8
	//        |  |
	//     rear  front
	//           i

	// iterate upto
	//  1  2  3  4  5  6  7  8
	//        |  |           |
	//     rear  front       |
	//                       i

	// i then will go back to 0
	//  1  2  3  4  5  6  7  8
	//  |     |  |
	//  |  rear  front
	//  i

	// i <= q.Rear condition will not work
	// as front can be easily after rear
	// hence, iterate, till you reach to rear
	for i := q.Front; i != q.Rear; i = (i + 1) % q.Size {
		fmt.Print(q.Arr[i], " > ")
	}
	fmt.Println(q.Arr[q.Rear])
}
