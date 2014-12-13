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
type arrInt []int
func (f arrInt) Len() int              	{ return len(f) }
func (f arrInt) Less(i, j int) bool    	{ return f[i] < f[j] }
func (f arrInt) Swap(i, j int)         	{ f[i], f[j] = f[j], f[i] }
 
type arrString []string
func (f arrString) Len() int           		{ return len(f) }
func (f arrString) Swap(i, j int)      		{ f[i], f[j] = f[j], f[i] }
func (f arrString) Less(i, j int) bool 	{return f[i] < f[j] }
 
func Shellsort(a Interface){ //O(n^2) -- fungsi untuk shellsort
    for n:=a.Len()/2 ; n> 0 ; n= (n+1)*5/11 {
        for i:=n ; i<a.Len() ; i++ {
            for j := i; j>=n && a.Less(j, j-n); j-=n {
                    a.Swap(j, j-n)
            }
        }
    }
}

func Ciurashellsort (a Interface){ // O(n log n) -- Ciura's shellsort
	c:= []int{701, 301, 132, 57, 23, 10, 4, 1}
	for n:=0; n<8; n++{
		gap:= c[n]
		for i:=gap ; i<a.Len(); i++{
			for j:=i ; j>=gap && a.Less(j,j-gap); j-=gap{
				a.Swap(j, j-gap)
				}
			}
		}
}

func Knuthshellsort (a Interface){ //O(n^3/2) -- Knuth's Shellsort
	for k:= a.Len()/3 ; k>=0; k--{
		t:=3
		for y:=1; y<k; y++{ t*=3}
		gap:= (t-1)/2
		for i:=gap ; i<a.Len(); i++{
			for j:=i ; j>=gap && a.Less(j,j-gap); j-=gap{
				a.Swap(j, j-gap)
				}
			}
		}
}

func main() {
	rand.Seed(time.Now().Unix())
	x := make(arrInt, 10)
	for i := 0; i < 10; i++ { x[i] = rand.Intn(50)}
	
	fmt.Print("Sebelum Sort: ")
	fmt.Println(x)
	Shellsort(x)
	fmt.Print("sesudah Shellsort: ")
	fmt.Println(x)
	
	Ss := make( arrString, 6)
	Ss[0], Ss[1], Ss[2], Ss[3], Ss[4], Ss[5] = "contoh", "shellsort", "yang", "biasa","asd", "go"
	fmt.Print("Sebelum Sort: ")
	fmt.Println(Ss)
	Shellsort(Ss)
	fmt.Print("sesudah Shellsort: ")
	fmt.Println(Ss)
    fmt.Println()
	
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ { x[i] = rand.Intn(50)}
	fmt.Print("Sebelum CiuraSort: ")
	fmt.Println(x)
	Ciurashellsort(x)
	fmt.Print("sesudah Ciura Shellsort: ")
	fmt.Println(x)
	
	xs := make( arrString, 5)
	xs[0], xs[1], xs[2], xs[3], xs[4] = "contoh","Ciura", "shellsort", "asd", "go"
	fmt.Print("Sebelum CiuraSort: ")
	fmt.Println(xs)
	Ciurashellsort(xs)
	fmt.Print("sesudah Ciura Shellsort: ")
	fmt.Println(xs)
	fmt.Println()
	
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ { x[i] = rand.Intn(50)}
	fmt.Print("Sebelum KnuthSort: ")
	fmt.Println(x)
	Knuthshellsort(x)
	fmt.Print("sesudah Knuth Shellsort: ")
	fmt.Println(x)
	
	ks := make( arrString, 5)
	ks[0], ks[1], ks[2], ks[3], ks[4] = "contoh", "Knuth", "shellsort", "asd", "go"
	fmt.Print("Sebelum KnuthSort: ")
	fmt.Println(ks)
	Knuthshellsort(ks)
	fmt.Print("sesudah Knuth Shellsort: ")
	fmt.Println(ks)
	fmt.Println()
	
}