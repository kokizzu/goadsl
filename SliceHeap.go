package main

import "fmt"

type Interface interface {
    Less(a interface{}) bool
}

type Int int
func (i Int) Less(a interface{}) bool { return i < a.(Int) }

type String string
func (s String) Less(a interface{}) bool { return s < a.(String) }

type Heap struct {
    data []Interface
}
// O(1)
func NewHeap() *Heap {
    return &Heap{make([]Interface,0,0)}
}

// O(log n)
func (h *Heap) Push(data Interface) {
    h.data = append(h.data,data)
    cur := len(h.data)-1
    par := (cur-1)/2
    for (cur != 0) && (h.data[par].Less(h.data[cur])) {
        h.data[cur], h.data[par] = h.data[par], h.data[cur]
        cur = par
        par = (cur-1)/2
    }
}


// O(log n)
func (h *Heap) Pop() (data Interface, ok bool) {
    if len(h.data) == 0 { return Int(0), false }
    data, ok = h.data[0], true
    h.data[0] = h.data[len(h.data)-1]
    h.data = h.data[:len(h.data)-1]
    cur, next := 0, 0
    if left := (cur*2)+1; (left < len(h.data)) && (h.data[next].Less(h.data[left])) { next = left }
    if right := (cur*2)+2; (right < len(h.data)) && (h.data[next].Less(h.data[right])) { next = right }
    for cur != next {
        h.data[cur], h.data[next] = h.data[next], h.data[cur]
        cur = next
        if left := (cur*2)+1; (left < len(h.data)) && (h.data[next].Less(h.data[left])) { next = left }
        if right := (cur*2)+2; (right < len(h.data)) && (h.data[next].Less(h.data[right])) { next = right }
    }
    return
}


func main() {
    h := NewHeap()
    h.Push(Int(6))
    h.Push(Int(5))
    h.Push(Int(3))
    h.Push(Int(1))
    h.Push(Int(8))
    h.Push(Int(7))
    h.Push(Int(2))
    h.Push(Int(4))
    for i := 0; i < 8; i++ {
        fmt.Print(h.data[i]," ")
    }
    fmt.Println()
    for i := 0; i < 8; i++ {
        if u, ok := h.Pop(); ok { fmt.Println(u) }
    }
	
    hs := NewHeap()
    hs.Push(String("baca"))
    hs.Push(String("tidur"))
    hs.Push(String("makan"))
    hs.Push(String("mandi"))
    hs.Push(String("duduk"))
    hs.Push(String("berdiri"))
    hs.Push(String("minum"))
    hs.Push(String("lari"))
    for i := 0; i < 8; i++ {
        fmt.Print(hs.data[i]," ")
    }
    fmt.Println()
    for i := 0; i < 8; i++ {
        if u, ok := hs.Pop(); ok { fmt.Println(u) }
    }
}
