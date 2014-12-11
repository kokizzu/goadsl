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

//O(n/2^n/2)
//Fungsi untuk memisahkan antara kiri dan kanan pivot
func partition(a Interface, first int, last int, pivotIndex int) int {
	a.Swap(first, pivotIndex)
	left := first + 1
	right := last
	for left <= right {
		for left <= last && a.Less(left, first) {
			left++
		}
		for right >= first && a.Less(first, right) {
			right--
		}
		if left <= right {
			a.Swap(left, right)
			left++
			right--
		}
	}
	a.Swap(first, right)
	return right
}

//O(1)
//Fungsi untuk menentukan pivot
func quicksortHelper(a Interface, first int, last int) {
	if first >= last {
		return
	}
	pivotIndex := partition(a, first, last, rand.Intn(last-first+1)+first)
	quicksortHelper(a, first, pivotIndex-1)
	quicksortHelper(a, pivotIndex+1, last)
}

//O(1)
//Fungsi untuk melakukan sort menggunakan quicksort
func quicksort(a Interface) {
	quicksortHelper(a, 0, a.Len()-1)
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
	rand.Seed(time.Now().Unix())
	x := make(SliceOfInt, 10)
	for i := 0; i < 10; i++ {
		x[i] = rand.Intn(20)
	}
	fmt.Print("Before: ")
	for i := 0; i < 10; i++ {
		fmt.Print(x[i], " ")
	}
	fmt.Println()
	quicksort(x)
	fmt.Print("After: ")
	for i := 0; i < 10; i++ {
		fmt.Print(x[i], " ")
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
	quicksort(y)
	fmt.Print("After: ")
	for i := 0; i < 3; i++ {
		fmt.Print(y[i], " ")
	}
	fmt.Println()

}
