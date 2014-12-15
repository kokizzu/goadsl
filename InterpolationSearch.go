package main

import "fmt"

type Interface interface{
	Len() int
	More(i int, j interface{}) bool
	Less(i int, j interface{}) bool
	Integerize(i int) int
	Convert(a interface{}) int
}

type arrInt []int
func (f arrInt) Len() int { return len(f) }
func (f arrInt) More(i int, j interface{}) bool { return f[i] > j.(int) }
func (f arrInt) Less(i int, j interface{}) bool { return f[i] < j.(int) }
func (f arrInt) Integerize(i int) int { return f[i] }
func (f arrInt) Convert(a interface{}) int { return a.(int) }

type arrString []string
func (f arrString) Len() int { return len(f) }
func (f arrString) More(i int,j interface{}) bool {return f[i] > j.(string) }
func (f arrString) Less(i int,j interface{}) bool {return f[i] < j.(string) }
func (f arrString) Integerize(i int) int {
	if (f[i][0] >= 'a') && (f[i][0] <= 'z') {
		return int(f[i][0])-'a'
	} else if (f[i][0] >= 'A') && (f[i][0] <= 'Z') {
		return int(f[i][0])-'A'
	} else {
		return int(f[i][0])
	}
}
func (f arrString) Convert(a interface{}) int {
	s := a.(string)
	if (s[0] >= 'a') && (s[0] <= 'z') {
		return int(s[0])-'a'
	} else if (s[0] >= 'A') && (s[0] <= 'Z') {
		return int(s[0])-'A'
	} else {
		return int(s[0])
	}
}

//O(log log n)	
func InterpolationSearch(a Interface, search interface{}) int {
	low := 0
	high := a.Len() - 1
	for low <= high {
		mid := low + (high - low) * ((a.Convert(search) - a.Integerize(low)) / (a.Integerize(high) - a.Integerize(low)))
		if  a.More(mid,search) {
			high = mid - 1
		} else if a.Less(mid,search) {
			low = mid + 1
		} else {
		return mid+1
		}
	}
	return -1
}

func main(){

	a := make(arrInt, 10)
	for i:=0; i<10; i++ {
		a[i] = i+1
	}

	fmt.Println("Data :", a)
	var search int
	search = 3
	fmt.Println("Nilai yang ingin dicari :", search)
	pos := InterpolationSearch(a, search)
	if pos == -1 {
		fmt.Println(search, "tidak ditemukan")
	} else {
		fmt.Println(search, "ada di urutan ke", pos)
	}

	s := make(arrString,5)
	s[0] = "bangun"
	s[1] = "duduk"
	s[2] = "makan"
	s[3] = "minum"
	s[4] = "tidur"
	fmt.Println("Data :", s)
	fmt.Println("kata yang ingin dicari : bangun")
	pos = InterpolationSearch(s, "bangun")
	if pos == -1 {
		fmt.Println("tidak ditemukan")
	} else {
		fmt.Println("ada di urutan ke", pos)
	}
	fmt.Println("kata yang ingin dicari : minum")
	pos = InterpolationSearch(s, "minum")
	if pos == -1 {
		fmt.Println("tidak ditemukan")
	} else {
		fmt.Println("ada di urutan ke", pos)
	}
	fmt.Println("kata yang ingin dicari : tidur")
	pos = InterpolationSearch(s, "tidur")
	if pos == -1 {
		fmt.Println("tidak ditemukan")
	} else {
		fmt.Println("ada di urutan ke", pos)
	}
	fmt.Println("kata yang ingin dicari : berdiri")
	pos = InterpolationSearch(s, "berdiri")
	if pos == -1 {
		fmt.Println("tidak ditemukan")
	} else {
		fmt.Println("ada di urutan ke", pos)
	}
}
