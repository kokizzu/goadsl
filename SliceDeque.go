package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Any interface{}

type Deque struct {
	depan []Any
	belakang []Any
}

// Create a new deque: O(1)
func NewDeque() *Deque {
	var q Deque
	q.depan = make([]Any,0,0)
	q.belakang = make([]Any,0,0)
	return &q
}

// Return size: O(1)
func (q *Deque) Count() int {
	return (len(q.depan)+len(q.belakang))
}

// Text whether deque is empty: O(1)
func (q *Deque) IsEmpty() bool {
	if (len(q.depan)+len(q.belakang)) == 0 { return true }
	return false
}

// Access first element: O(1)
func (q *Deque) Front() Any {
	if len(q.depan) == 0 { return nil }
	return q.depan[0]
}

// Access last element: O(1)
func (q *Deque) Back() Any {
	if len(q.belakang) == 0 { return nil }
	return q.belakang[len(q.belakang)-1]
}

// Insert new element into the end of deque: O(1)
func (q *Deque) enqueueBack(data Any) {
	q.belakang = append(q.belakang,data)
}

// Insert new element into the head of deque: O(1)
func (q *Deque) enqueueFront(data Any) {
	q.depan = append(q.depan,data)
}

// Remove first element from deque: O(1)
func (q *Deque) dequeueFront() Any {
	if len(q.depan) == 0 {
		q.depan = q.belakang
		data := q.depan[0]
		q.depan = q.depan[1:]
		for i:=0;i<len(q.belakang);i++ {
			q.belakang = make([]Any,0,0)
		}
	        return data
	} else {
		data := q.depan[len(q.depan)-1]
		q.depan = q.depan[0:(len(q.depan)-1)]
		return data
	}
}

// Remove last element from deque: O(1)
func (q *Deque) dequeueBack() Any {
	if len(q.belakang) == 0 {
		data := q.depan[0]
		q.depan = q.depan[1:(len(q.depan)-1)]
		return data
	} else {
		data := q.belakang[(len(q.belakang)-1)]
		q.belakang = q.belakang[0:(len(q.belakang)-1)]
		return data
	}
}

// Print elements from deque: O(N)
func (q *Deque) Print() {
	if (len(q.depan)+len(q.belakang)) == 0 { fmt.Println("| Empty |"); return }
	fmt.Print("|")
	var k,i,j int
	k=0
	for i = (len(q.depan)-1); i>=0; i-- {
		fmt.Printf(" %d: %#v |", k, q.depan[i])
		k++
	}
	i=len(q.depan)+1
	for j=0; j<len(q.belakang); j++ {
		fmt.Printf(" %d: %#v |", i, q.belakang[j])
		i++
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
	q := NewDeque()
	fmt.Println("Enqueue Front Numbers *Remember The Last One Become The Head*")
	for i := 0; i < 6; i++ {
		a := RandInt(0,100)
		fmt.Print(a)
		if i == 5 { fmt.Println() } else { fmt.Print(" ") }
		q.enqueueFront(a)
	}
	q.Print()
	fmt.Println()

	fmt.Println("Enqueue Back String")
	s:= "abc"
	fmt.Println(s)
	q.enqueueBack(s)
	q.Print()

	fmt.Println()
	fmt.Println("Dequeue Front")
	for i := 0; i < 6; i++ {
		fmt.Print(q.dequeueFront())
		if i == 5 { fmt.Println() } else { fmt.Print(" ") }
	}
	q.Print()

	fmt.Println()
	fmt.Println("Enqueue Back Numbers")
	for i := 0; i < 6; i++ {
		a := RandInt(0,100)
		fmt.Print(a)
		if i == 5 { fmt.Println() } else { fmt.Print(" ") }
		q.enqueueBack(a)
	}
	q.Print()

	fmt.Println()
	fmt.Println("Dequeue Back All")
	for i := 0; i < 7; i++ {
		fmt.Print(q.dequeueBack())
		if i == 6 { fmt.Println() } else { fmt.Print(" ") }
	}
	q.Print()

	fmt.Println()
	fmt.Println("Enqueue Front Numbers *Remember The Last One Become The Head*")
	for i := 0; i < 6; i++ {
		a := RandInt(0,100)
		fmt.Print(a)
		if i == 5 { fmt.Println() } else { fmt.Print(" ") }
		q.enqueueFront(a)
	}
	q.Print()
	fmt.Println()

	fmt.Println("Dequeue Front All")
	for i := 0; i < 6; i++ {
		fmt.Print(q.dequeueFront())
		if i == 5 { fmt.Println() } else { fmt.Print(" ") }
	}
	q.Print()


}
