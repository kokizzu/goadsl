package main

import "fmt"
import "math/rand"
import "time"

type Interface interface {
	Len() int
	Less(i, j int) bool
	Equal(i, j int) bool
	Swap(i, j int)
}

type IntArr []int
func (IA IntArr) Len() int { return len(IA) }
func (IA IntArr) Less(x, y int) bool { return IA[x] < IA[y] } 
func (IA IntArr) Equal(x, y int) bool { return IA[x] == IA[y] }
func (IA IntArr) Swap(x int, y int) { IA[x], IA[y] = IA[y], IA[x] }

type StringArr []string
func (SA StringArr) Len() int { return len(SA) }
func (SA StringArr) Less(x, y int) bool { return SA[x] < SA[y] } 
func (SA StringArr) Equal(x, y int) bool { return SA[x] == SA[y] }
func (SA StringArr) Swap(x int, y int) { SA[x], SA[y] = SA[y], SA[x] }

type Float32Arr []float32
func (F32A Float32Arr) Len() int { return len(F32A) }
func (F32A Float32Arr) Less(x, y int) bool { return F32A[x] < F32A[y] } 
func (F32A Float32Arr) Equal(x, y int) bool { return F32A[x] == F32A[y] }
func (F32A Float32Arr) Swap(x int, y int) { F32A[x], F32A[y] = F32A[y], F32A[x] }

type Float64Arr []float64
func (F64A Float64Arr) Len() int { return len(F64A) }
func (F64A Float64Arr) Less(x, y int) bool { return F64A[x] < F64A[y] } 
func (F64A Float64Arr) Equal(x, y int) bool { return F64A[x] == F64A[y] }
func (F64A Float64Arr) Swap(x int, y int) { F64A[x], F64A[y] = F64A[y], F64A[x] }

func CocktailSort(data Interface) {
	last := data.Len() - 1
	for {
		swapped := false
		for i:=0; i<last; i++ {
			if data.Less(i+1, i) {
				data.Swap(i+1, i)
				swapped = true
			}
		}
		if swapped == false {
			return
		}
		
		swapped = false
		for i:=last-1; i>=0; i-- {
			if data.Less(i+1, i) {
				data.Swap(i+1, i)
				swapped = true 
			} 
		}
		if swapped == false {
			return 
		}
	}
}

func main(){
	IntSlice := make(IntArr, 12)
	fmt.Print("Awal : ")
	for i:=0; i<12; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		IntSlice[i] = rand.Intn(50)
		fmt.Print(IntSlice[i], " ")
	}
	fmt.Println()
	CocktailSort(IntSlice)
	fmt.Print("Sesudah Sorting : ")
	for i:=0; i<12; i++ {
		fmt.Print(IntSlice[i], " ")
	}
	fmt.Println("\n")
	
	StringSlice := StringArr{"nama", "saya", "adalah", "alvreda", "ivena"}
	fmt.Print("Awal : ")
	for i:=0; i<5; i++ {
		fmt.Print(StringSlice[i], " ")
	}
	fmt.Println()
	CocktailSort(StringSlice)
	fmt.Print("Sesudah Sorting : ")
	for i:=0; i<5; i++ {
		fmt.Print(StringSlice[i], " ")
	}
	fmt.Println("\n")
	
	Float32Slice := make(Float32Arr, 6)
	fmt.Print("Awal : ")
	for i:=0; i<6; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		Float32Slice[i] = rand.Float32()
		fmt.Print(Float32Slice[i], " ")
	}
	fmt.Println()
	CocktailSort(Float32Slice)
	fmt.Print("Sesudah Sorting : ")
	for i:=0; i<6; i++ {
		fmt.Print(Float32Slice[i], " ")
	}
	fmt.Println("\n")
	
	Float64Slice := make(Float64Arr, 6)
	fmt.Print("Awal : ")
	for i:=0; i<6; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		Float64Slice[i] = rand.Float64()
		fmt.Print(Float64Slice[i], " ")
	}
	fmt.Println()
	CocktailSort(Float64Slice)
	fmt.Print("Sesudah Sorting : ")
	for i:=0; i<6; i++ {
		fmt.Print(Float64Slice[i], " ")
	}
	fmt.Println()
}
