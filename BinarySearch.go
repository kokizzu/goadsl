package main
import (
	"fmt"
	
)

type Interface interface{
	Len() int
	More(i int, j interface{}) bool
	Less(i int, j interface{}) bool
}

type arrInt []int
func (f arrInt) Len() int { return len(f) }
func (f arrInt) More(i int, j interface{}) bool { return f[i] > j.(int) }
func (f arrInt) Less(i int, j interface{}) bool { return f[i] < j.(int) }

type arrString []string
func (f arrString) Len() int { return len(f) }
func (f arrString) More(i int,j interface{}) bool {return f[i] > j.(string) }
func (f arrString) Less(i int,j interface{}) bool {return f[i] < j.(string) }


func binarySearch(a Interface, value interface{}) int {  //O(log n)
	low := 0
	high := a.Len() - 1
	for low <= high {
		mid := (low + high) / 2
		if a.More(mid,value) {
			high = mid - 1
		} else if a.Less(mid,value) {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main(){

	list := make(arrInt, 10)
	for i := 0; i < 10; i++ {
		list[i] = i+1
		
	}
	
	posisi:=(binarySearch(list, 1))
	if posisi == -1 {
		fmt.Println("data yang dicari tidak ditemukan")
	}else {
		fmt.Println("ada di ada di array elemen ke", posisi)
	}
	posisi =(binarySearch(list, 10))
	if posisi == -1 {
		fmt.Println("data yang dicari tidak ditemukan")
	}else {
		fmt.Println("ada di ada di array elemen ke", posisi)
	}
	posisi =(binarySearch(list, 5))
	if posisi == -1 {
		fmt.Println("data yang dicari tidak ditemukan")
	}else {
		fmt.Println("ada di ada di array elemen ke", posisi)
	}
	posisi =(binarySearch(list, 13))
	if posisi == -1 {
		fmt.Println("data yang dicari tidak ditemukan")
	}else {
		fmt.Println("ada di ada di array elemen ke", posisi)
	}
	

	
	posisi=0
	Ss := make(arrString, 6)
	Ss[0], Ss[1], Ss[2], Ss[3], Ss[4], Ss[5] = "ab", "abc", "abcd", "abcdf", "abcdefg", "abcdefgh"
	
	
	posisi =(binarySearch(Ss, "ab" ))
	if posisi == -1 {
		fmt.Println("kata yang anda cari tidak ditemukan")
	}else {
		fmt.Println("ada di array elemen ke", posisi)
	}
	posisi =(binarySearch(Ss, "abcdefgh" ))
	if posisi == -1 {
		fmt.Println("kata yang anda cari tidak ditemukan")
	}else {
		fmt.Println("ada di array elemen ke", posisi)
	}
	posisi =(binarySearch(Ss, "abcd" ))
	if posisi == -1 {
		fmt.Println("kata yang anda cari tidak ditemukan")
	}else {
		fmt.Println("ada di array elemen ke", posisi)
	}
	posisi =(binarySearch(Ss, "a" ))
	if posisi == -1 {
		fmt.Println("kata yang anda cari tidak ditemukan")
	}else {
		fmt.Println("ada di array elemen ke", posisi)
	}
	
	
	
	
}
