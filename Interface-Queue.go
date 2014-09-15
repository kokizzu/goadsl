package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Any interface{}

type Queue struct {
	data []Any
	head, tail, count int
	size int
}

func NewQueue(size int) *Queue {
	var q Queue
	q.data = make([]Any,size)
	q.head, q.tail, q.count = 0, 0, 0
	q.size = size
	return &q
}

func (q *Queue) Count() int {
	return q.count
}

func (q *Queue) IsEmpty() bool {
	if q.tail == 0 { return true }
	return false
}

func (q *Queue) Front() Any {
	return q.data[q.head]
}

func (q *Queue) Back() Any {
	return q.data[q.tail-1]
}

func (q *Queue) enqueue(data Any) {
	if q.count < q.size {
		q.data[q.tail] = data
		q.tail = (q.tail+1)%q.size
		q.count++
		q.size++
	}
}

func (q *Queue) dequeue() Any {
	if q.count != 0 {
		front := q.data[q.head]
		q.head = (q.head+1)%q.size
		q.count--
		return front
	}
	return nil
}

func (q *Queue) PrintN(n int) {
	if (q.count == 0) { fmt.Println("NIL"); return }
	fmt.Print("|")
	in := q.head
	for i := 0; i < n; i++ {
		fmt.Printf(" %d: %#v |", i, q.data[in])
		in = (in+1)%q.size
		if in == q.head { break }
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
	q := NewQueue(100)
	fmt.Println("Enqueue Angka")
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
