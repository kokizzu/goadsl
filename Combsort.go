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

//O(n^2)
//fungsi untuk mengsort dengan comb sort
func Combsort(data Interface) {
	if data.Len() < 2 {
		return
	}
	for gap := data.Len(); ; {
		if gap > 1 {
			gap = gap * 4 / 5
		}
		swapped := false
		for i := 0; ; {
			if data.Less(i+gap, i) {
				data.Swap(i, i+gap)
				swapped = true
			}
			i++
			if i+gap >= data.Len() {
				break
			}
		}
		if gap == 1 && !swapped {
			break
		}
	}
}

type SliceOfInt []int

type SliceOfString []string

//fungsi Len, Swap, dan Less untuk SliceOfInt dan SliceOfString
func (a SliceOfInt) Len() int              { return len(a) }
func (a SliceOfInt) Swap(i, j int)         { a[i], a[j] = a[j], a[i] }
func (a SliceOfInt) Less(i, j int) bool    { return a[i] < a[j] }
func (a SliceOfString) Less(i, j int) bool { return a[i] < a[j] }
func (a SliceOfString) Len() int           { return len(a) }
func (a SliceOfString) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

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
	sort(x)
	fmt.Print("After: ")
	for i := 0; i < 10; i++ {
		fmt.Print(x[i], " ")
	}
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
	Combsort(y)
	fmt.Print("After: ")
	for i := 0; i < 3; i++ {
		fmt.Print(y[i], " ")
	}
	fmt.Println()

}
