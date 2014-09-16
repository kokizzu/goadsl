package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Any interface{}

type Queue struct {
	data []Any
}

// Create a new queue: O(1)
func NewQueue() *Queue {
	var q Queue
	q.data = make([]Any,0,0)
	return &q
}

// Return size: O(1)
func (q *Queue) Count() int {
	return len(q.data)
}

// Text whether queue is empty: O(1)
func (q *Queue) IsEmpty() bool {
	if len(q.data) == 0 { return true }
	return false
}

// Access first element: O(1)
func (q *Queue) Front() Any {
	if len(q.data) == 0 { return nil }
	return q.data[0]
}

// Access last element: O(1)
func (q *Queue) Back() Any {
	if len(q.data) == 0 { return nil }
	return q.data[len(q.data)-1]
}

// Insert new element into the end of queue: O(1)
func (q *Queue) enqueue(data Any) {
	q.data = append(q.data,data)
}

// Remove first element from queue: O(1)
func (q *Queue) dequeue() Any {
	if len(q.data) == 0 { return nil }
	data := q.data[0]
	q.data = q.data[1:]
	return data
}

// Print first N elements from queue: O(N)
func (q *Queue) PrintN(n int) {
	if len(q.data) == 0 { fmt.Println("| Empty |"); return }
	fmt.Print("|")
	for i := 0; (i < len(q.data)) && (i < n); i++ {
		fmt.Printf(" %d: %#v |", i, q.data[i])
	}
	fmt.Println()
}

func init() {
	rand.Seed(time.Now().Unix())
}

func RandInt(min, max int) int {
	return rand.Intn(max - min) + min
}

func main() {
	q := NewQueue()
	fmt.Println("Enqueue Numbers")
	for i := 0; i < 6; i++ {
		a := RandInt(0,100)
		fmt.Print(a)
		if i == 5 { fmt.Println() } else { fmt.Print(" ") }
		q.enqueue(a)
	}
	q.PrintN(q.Count())
	fmt.Println()
	fmt.Println("Enqueue String")
	s:= "abc"
	fmt.Println(s)
	q.enqueue(s)
	q.PrintN(q.Count())

	fmt.Println()
	fmt.Println("Dequeue All")
	for i := 0; i < 7; i++ {
		fmt.Print(q.dequeue())
		if i == 6 { fmt.Println() } else { fmt.Print(" ") }
	}
	q.PrintN(q.Count())
}
