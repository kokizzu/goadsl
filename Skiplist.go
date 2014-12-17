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
          value interface{}
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

// O(N*M)
func (s *skiplist) Search(nilai Interface) (interface{}, bool){ // fungsi yang berisi (s menunjuk ke pointer skiplist) untuk merubah nilai pada string
ketemu:=false
var valuenya, nilainya interface{}
          level := 0 // level bernilai nol
          for ; rand.Int31n(2) == 1; level++ { // lakukan proses tersebut selama sesuai dengan kondisi
                    if level > s.T { // kondisi
                         s.T = level //memasukan nilai level kedalam s.T
                         break // break jika selesai memasukkan nilai
                    }
          }
                    current := s.kpl // current di isi dengan fungsi s.kpl
          for i := s.T; i >= 0; i-- { //lakukan selama memnuhi kondisi
                    for ; current.next[i] != nil; current = current.next[i] { //lakukan selama kondisi ini
                         next := current.next[i]
                         if next.nilai==nilai {
nilainya=next.nilai
                         ketemu=true
                    valuenya=next.value
                         break
                         }
     
                    }
                    if i > level { // kondisi
                         continue
                    }
                           
          }
          if valuenya==nil {
          fmt.Println("No Match Found!")
          } else if valuenya!=nil{
                                   fmt.Println("Found! =", nilainya)
          }
          return valuenya, ketemu
}

// O(N*M)
func (s *skiplist) Delete(nilai Interface) { // fungsi yang berisi (s menunjuk ke pointer skiplist) untuk merubah nilai pada string

          level := 0 // level bernilai nol
          for ; rand.Int31n(2) == 1; level++ { // lakukan proses tersebut selama sesuai dengan kondisi
                    if level > s.T { // kondisi
                         s.T = level //memasukan nilai level kedalam s.T
                         break // break jika selesai memasukkan nilai
                    }
          }
                    current := s.kpl // current di isi dengan fungsi s.kpl
          for i := s.T; i >= 0; i-- { //lakukan selama memnuhi kondisi
                    for ; current.next[i] != nil; current = current.next[i] { //lakukan selama kondisi ini
                         next := current.next[i] //
                         if next.nilai==nilai {
next.nilai=nil
next.value=nil
                         break
                         }

                    }
                    if i > level { // kondisi
                         continue
                    }
   
          }

}

// O(N)
func (s *skiplist) Set(nilai Interface, value interface{}) { // fungsi yang berisi (s menunjuk ke pointer skiplist) untuk merubah nilai pada string
          level := 0 // level bernilai nol
          for ; rand.Int31n(2) == 1; level++ { // lakukan proses tersebut selama sesuai dengan kondisi
                    if level > s.T { // kondisi
                         s.T = level //memasukan nilai level kedalam s.T
                         break // break jika selesai memasukkan nilai
                    }
          }
                   
          node := &skiplistNode{ // node menunjuk alamat dari skiplist
                    nilai: nilai ,//nilai selalu menghasilkan nilai yang sama pada proses selanjutnya
                    next: make([]*skiplistNode, level+1),// membuat array of pointer skiplistNode yang panjangnya level + 1
                    value: value,
                    }
          current := s.kpl // current di isi dengan fungsi s.kpl
                    q := new(skiplistNode)
          q = s.kpl //untuk insert dari setiap huruf string, kita harus memulainya dari root
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
                    q.value=value
                    node.next[i] = current.next[i]
                    current.next[i] = node
                    node.prev = current
                   
          }
  
}

// O(N)o
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
                         if node.nilai!=nil {
                         node.nilai.Print()
                         print("-")
                         i = rate + 1 //
                         }
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
          sl.Set(Int(6), 1)
          sl.Set(Int(5), 2)
               sl.Set(Int(4), 4)
               sl.Set(Int(9), 0)
          sl.draw() // 4-5-6-9-
           
          valuenya,ketemu:=sl.Search(Int(5)) // Found! = 5
          fmt.Println(valuenya) //2
          fmt.Println(ketemu) // true
          valuenya,ketemu=sl.Search(Int(4)) // Found! = 4
          fmt.Println(valuenya) //4
          fmt.Println(ketemu) // true
          valuenya,ketemu=sl.Search(Int(6)) // Found! = 6
          fmt.Println(valuenya) //1
          fmt.Println(ketemu) // true
               valuenya,ketemu=sl.Search(Int(9)) // Found! = 9
          fmt.Println(valuenya) //0
          fmt.Println(ketemu) // true
                   
               sl.Delete(Int(5))
               sl.draw() // 4---6-9-
               valuenya,ketemu=sl.Search(Int(5)) //  No Match Found!
          fmt.Println(valuenya) //<nil>
          fmt.Println(ketemu) // false
          sl.draw() // 4---6-9-
          
               rand.Seed(8)   
               s2 := New()
          s2.Set(String("h"), 1)
          s2.Set(String("c"), 2)
               s2.Set(String("i"), 4)
               s2.Set(String("s"), 0)
          s2.draw() // c-h-i-s-
           
          valuenya,ketemu=s2.Search(String("c")) // Found! = c
          fmt.Println(valuenya) //2
          fmt.Println(ketemu) // true
          valuenya,ketemu=s2.Search(String("i")) // Found! = i
          fmt.Println(valuenya) //4
          fmt.Println(ketemu) // true
          valuenya,ketemu=s2.Search(String("s")) // Found! = s
          fmt.Println(valuenya) //0
          fmt.Println(ketemu) // true
               valuenya,ketemu=s2.Search(String("h")) // Found! = h
          fmt.Println(valuenya) //1
          fmt.Println(ketemu) // true
                   
               s2.Delete(String("i"))
               s2.draw() // c-h---s-
               valuenya,ketemu=s2.Search(String("i")) // No Match Found!
          fmt.Println(valuenya) //<nil>
          fmt.Println(ketemu) // false
          s2.draw() // c-h---s-


    }
