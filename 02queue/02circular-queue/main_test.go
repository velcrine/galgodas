package main

import (
	"reflect"
	"testing"
)

func TestLove(t *testing.T) {
	q := NewQueue(3)

	// no element there, dequeue will fail
	_, err := q.Dequeue()
	if err == nil {
		t.Fatal("dequeue in empty queue should fail")
	}

	// all 3 enqueue will pass
	err = q.Enqueue(1)
	if err != nil {
		t.Fatal("enqueue must work if there is space enough")
	}
	err = q.Enqueue(2)
	if err != nil {
		t.Fatal("enqueue must work if there is space enough")
	}
	err = q.Enqueue(3)
	if err != nil {
		t.Fatal("enqueue must work if there is space enough")
	}

	// max size reached
	err = q.Enqueue(4)
	if err == nil {
		t.Fatal("enqueue must 'not' work if there is space enough")
	}

	// dequeue will succeed
	_, err = q.Dequeue()
	if err != nil {
		t.Fatal("dequeue must work if there are elements in queue")
	}

	// enqueue will run successfully as queue is still not empty
	err = q.Enqueue(4)
	if err != nil {
		t.Fatal("enqueue must work if there is space enough")
	}

	// 2, 3. 4
	q.Display()

	// enqueue will fail as rear +1 == front
	err = q.Enqueue(5)
	if err == nil {
		t.Fatal("enqueue must 'not' work if there is space enough")
	}

	// 2 dequeue to remove 2 and 3
	_, err = q.Dequeue()
	if err != nil {
		t.Fatal("dequeue must work if there are elements enough")
	}
	_, err = q.Dequeue()
	if err != nil {
		t.Fatal("dequeue must work if there are elements enough")
	}

	// enqueue will work as queue is reinitialized
	err = q.Enqueue(5)
	if err != nil {
		t.Fatal("enqueue must work if there is space enough")
	}
	err = q.Enqueue(6)
	if err != nil {
		t.Fatal("enqueue must work if there is space enough")
	}
	expectedQueue := [3]int{4, 5, 6}
	if reflect.DeepEqual(q.Arr, expectedQueue) {
		t.Fatalf("queue state is not as expected: is %v, expected %v", q.Arr, expectedQueue)
	}
	q.Display()
}
