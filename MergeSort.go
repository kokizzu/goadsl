package main

import "fmt"
import "math/rand"
import "time"

type Interface interface {
	Len() int
	Less(x, y int) bool
	Equal(x, y int) bool
	Swap(x, y int)
	Get(x int) interface{}
	Set(x int, src Interface, y int)
}

type IntArr []int
func (IA IntArr) Len() int { return len(IA) }
func (IA IntArr) Less(x, y int) bool { return IA[x] < IA[y] } 
func (IA IntArr) Equal(x, y int) bool { return IA[x] == IA[y] }
func (IA IntArr) Swap(x int, y int) { IA[x], IA[y] = IA[y], IA[x] }
func (IA IntArr) Get(x int) interface{} { return IA[x] }
func (IA IntArr) Set(x int, src Interface, y int) { IA[x] = src.Get(y).(int) }

type StringArr []string
func (SA StringArr) Len() int { return len(SA) }
func (SA StringArr) Less(x, y int) bool { return SA[x] < SA[y] } 
func (SA StringArr) Equal(x, y int) bool { return SA[x] == SA[y] }
func (SA StringArr) Swap(x int, y int) { SA[x], SA[y] = SA[y], SA[x] }
func (SA StringArr) Get(x int) interface{} { return SA[x] }
func (SA StringArr) Set(x int, src Interface, y int) { SA[x] = src.Get(y).(string) }

type Float32Arr []float32
func (F32A Float32Arr) Len() int { return len(F32A) }
func (F32A Float32Arr) Less(x, y int) bool { return F32A[x] < F32A[y] } 
func (F32A Float32Arr) Equal(x, y int) bool { return F32A[x] == F32A[y] }
func (F32A Float32Arr) Swap(x int, y int) { F32A[x], F32A[y] = F32A[y], F32A[x] }
func (F32A Float32Arr) Get(x int) interface{} { return F32A[x] }
func (F32A Float32Arr) Set(x int, src Interface, y int) { F32A[x] = src.Get(y).(float32) }

type Float64Arr []float64
func (F64A Float64Arr) Len() int { return len(F64A) }
func (F64A Float64Arr) Less(x, y int) bool { return F64A[x] < F64A[y] } 
func (F64A Float64Arr) Equal(x, y int) bool { return F64A[x] == F64A[y] }
func (F64A Float64Arr) Swap(x int, y int) { F64A[x], F64A[y] = F64A[y], F64A[x] }
func (F64A Float64Arr) Get(x int) interface{} { return F64A[x] }
func (F64A Float64Arr) Set(x int, src Interface, y int) { F64A[x] = src.Get(y).(float64) }

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MergeSort(data, temp Interface) {
	for n:=2; (n/2)<data.Len(); n*=2 {
		for i:=0; i<data.Len(); i+=n {
			left := i
			right := min(i+n, data.Len())
			mid := i+(n/2)
			for j:=left; j<right; j++ {
				temp.Set(j, data, j)
			}
			l := i
			r := mid
			for j:=left; j<right; j++ {
			 	if l < mid && r < right {
			 		if temp.Less(l, r) {
			 			data.Set(j, temp, l)
			 			l += 1
			 		} else {
						data.Set(j, temp, r)
						r += 1
			 		}
			 	} else if l < mid {
			 		data.Set(j, temp, l)
			 		l += 1
			 	} else {
					data.Set(j, temp, r)
					r += 1
			 	}
			}
		}	
	}
}

func main(){	
	IntSlice := make(IntArr, 12)
	ITemp := make(IntArr, IntSlice.Len())
	fmt.Print("Awal : ")
	for i:=0; i<12; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		IntSlice[i] = rand.Intn(50)
		fmt.Print(IntSlice[i], " ")
	}
	fmt.Println()
	MergeSort(IntSlice, ITemp)
	fmt.Print("Sesudah Sorting : ")
	for i:=0; i<12; i++ {
		fmt.Print(IntSlice[i], " ")
	}
	fmt.Println("\n")
	
	StringSlice := StringArr{"nama", "saya", "adalah", "alvreda", "ivena"}
	Stemp := make(StringArr, StringSlice.Len())
	fmt.Print("Awal : ")
	for i:=0; i<5; i++ {
		fmt.Print(StringSlice[i], " ")
	}
	fmt.Println()
	MergeSort(StringSlice, Stemp)
	fmt.Print("Sesudah Sorting : ")
	for i:=0; i<5; i++ {
		fmt.Print(StringSlice[i], " ")
	}
	fmt.Println("\n")
	
	Float32Slice := make(Float32Arr, 6)
	F32temp := make(Float32Arr, Float32Slice.Len())
	fmt.Print("Awal : ")
	for i:=0; i<6; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		Float32Slice[i] = rand.Float32()
		fmt.Print(Float32Slice[i], " ")
	}
	fmt.Println()
	MergeSort(Float32Slice, F32temp)
	fmt.Print("Sesudah Sorting : ")
	for i:=0; i<6; i++ {
		fmt.Print(Float32Slice[i], " ")
	}
	fmt.Println("\n")
	
	Float64Slice := make(Float64Arr, 6)
	F64temp := make(Float64Arr, Float64Slice.Len())
	fmt.Print("Awal : ")
	for i:=0; i<6; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		Float64Slice[i] = rand.Float64()
		fmt.Print(Float64Slice[i], " ")
	}
	fmt.Println()
	MergeSort(Float64Slice, F64temp)
	fmt.Print("Sesudah Sorting : ")
	for i:=0; i<6; i++ {
		fmt.Print(Float64Slice[i], " ")
	}
	fmt.Println()
}
