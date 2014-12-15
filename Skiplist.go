package main

import (
	"fmt"
	"math/rand"
)
//O(1)
type Interface interface{ 
	Greater(a interface{}) bool // agar dapat menggunakan string dan integer secara bersamaan 
	Print()
}
//O(1)
type Int int
func (i Int) Greater(a interface{}) bool { return i > a.(Int) }// kembali i jika i lebih besar dari a.(Int)
func (i Int) Print() { fmt.Print(int(i)) }// mencetak output integer i
//O(1)
type String string
func (s String) Greater(a interface{}) bool { return s > a.(String) }// kembali ke s jika s lebih besar dari pada a.(String)
func (s String) Print() { fmt.Print(string(s)) } //mencetak output string s

const Tmax = 64
//O(1)
type skiplist struct {
	T int // T itu untuk menunjukan tinggi yang berupa bilangan bulat
	kpl *skiplistNode // kpl untuk menunjukkan head dengan cara memanggil pointer skiplistNode
}
//O(1)
type skiplistNode struct {
	nilai Interface // nilai disini menghasilkan output berupa interface
	next  []*skiplistNode // next disini itu untuk memasukkan output yang selanjutnya ke node yang sesuai
	prev  *skiplistNode // prev itu
}
//O(1)
func New() *skiplist { // fungsi untuk membuat list yang baru dengan cara memanggil pointer skiplist
	kpl := &skiplistNode{ // head menunjuk alamat si node
		nilai: nil,
		next:  make([]*skiplistNode, Tmax),// membuat array of pointer skiplistNode yang panjangnya Tmax
	}
	return &skiplist{ //mengembalikan alamat skiplist
		T:   0, // tinggi disini bernilai 0
		kpl: kpl, // head yang dipilih selalu menghasilkan nilai yang sama pada proses selanjutnya
	}
}
// O(N)
func (s *skiplist) Set(nilai Interface) { // fungsi yang berisi (s menunjuk ke pointer skiplist) untuk merubah nilai pada string
	level := 0 // level bernilai nol
	for ; rand.Int31n(2) == 1; level++ { // lakukan proses tersebut selama sesuai dengan kondisi
		if level > s.T { // kondisi
			s.T = level //memasukan nilai level kedalam s.T
			break // break jika selesai memasukkan nilai
		}
	}
	node := &skiplistNode{ // node menunjuk alamat dari skiplist
		nilai: nilai, // nilai selalu menghasilkan nilai yang sama pada proses selanjutnya
		next:  make([]*skiplistNode, level+1),// membuat array of pointer skiplistNode yang panjangnya level + 1
	}
	current := s.kpl // current di isi dengan fungsi s.kpl
	for i := s.T; i >= 0; i-- { //lakukan selama memnuhi kondisi
		for ; current.next[i] != nil; current = current.next[i] { //lakukan selama kondisi ini
			next := current.next[i] //
			if next.nilai.Greater(nilai) { //
				break
			}
		}
		if i > level { // kondisi
			continue 
		}
		node.next[i] = current.next[i]
		current.next[i] = node
		node.prev = current
	}
}
// O(N)
func (s *skiplist) draw() {
	rate := make(map[Interface]int) //
	for i, node := 0, s.kpl.next[0]; node != nil; node = node.next[0] {
		rate[node.nilai] = i
		i++
	}

	for level := s.T; level >= 0; level-- {
		if s.kpl.next[level] == nil {
			continue
		}
		for i, node := 0, s.kpl.next[level]; node != nil; node = node.next[level] {
			rate := rate[node.nilai]
			for j := 0; j < rate-i; j++ {
				print("--")
			}
			node.nilai.Print()
			print("-")
			i = rate + 1 //
		}
		println("\n")//untuk mencetak input dan setelah itu pindah ke line berikutnya atau enter
	}
	fmt.Println("")
}

//pada fungsi ini akan menyimpan data yang digunakan sebagai input
func main() {
	// change "64" to anything else to see
	rand.Seed(8)
	sl := New()
	sl.Set(String("h"))
	sl.Set(String("c"))
	sl.Set(String("i"))
	sl.Set(String("2"))
	sl.Set(String("0"))
	sl.Set(String("1"))
	sl.Set(String("3"))
	sl.Set(String("A"))
	sl.Set(String("S"))
	sl.Set(String("D"))
	sl.draw()
	s2 := New()
	s2.Set(Int(13))
	s2.Set(Int(4))
	s2.Set(Int(2))
	s2.Set(Int(17))
	s2.Set(Int(101))
	s2.Set(Int(75))
	s2.Set(Int(189))
	s2.Set(Int(0))
	s2.Set(Int(-12))
	s2.Set(Int(54))
	s2.draw()
}
