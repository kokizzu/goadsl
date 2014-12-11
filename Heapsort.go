package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Interface interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
}

//O(len-1)
//Fungsi untuk melakukan sort menggunakan heapsort
func Heapsort(a Interface) {
	for start := (a.Len() - 2) / 2; start >= 0; start-- {
		check(a, start, a.Len()-1)
	}
	for end := a.Len() - 1; end > 0; end-- {
		a.Swap(0, end)
		Check(a, 0, end-1)
	}
}

//O(len-2)
//Fungsi untuk memindahkan kebelakang
func Check(a Interface, start, end int) {
	for root := start; root*2+1 <= end; {
		child := root*2 + 1
		if child+1 <= end && a.Less(child, child+1) {
			child++
		}
		if !a.Less(root, child) {
			return
		}
		a.Swap(root, child)
		root = child
	}
}

type SliceOfInt []int
type SliceOfString []string

func (a SliceOfInt) Len() int              { return len(a) }
func (a SliceOfInt) Swap(i, j int)         { a[i], a[j] = a[j], a[i] }
func (a SliceOfInt) Less(i, j int) bool    { return a[i] < a[j] }
func (a SliceOfString) Len() int           { return len(a) }
func (a SliceOfString) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SliceOfString) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	a := make(SliceOfInt, 10)
	fmt.Println()
	fmt.Print("Before: ")
	for i := 0; i < 10; i++ {
		a[i] = rand.Intn(20)
		fmt.Print(a[i], " ")
	}
	Heapsort(a)
	fmt.Println()
	fmt.Print("After: ")
	for i := 0; i < 10; i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
	fmt.Println()

	y := make(SliceOfString, 3)

	fmt.Print("Before: ")
	y[0] = "Ruby"
	y[1] = "Emerald"
	y[2] = "Sapphire"
	for i := 0; i < 3; i++ {
		fmt.Print(y[i], " ")
	}

	fmt.Println()
	Heapsort(y)
	fmt.Print("After: ")
	for i := 0; i < 3; i++ {
		fmt.Print(y[i], " ")
	}
	fmt.Println()

}
