//Trie Implementation in golang by Judika
//Judika Gultom
//HCI A
package main

 import "fmt"

 
type node struct {
    data int
     link[26]*node
}
 
type Trie struct {
    Root *node
}

//fungsi inisialisasi, berguna untuk menginisiali sasi bahwa trie yang baru dibuat, rootnya hrs nil
//bigO=0
func (T *Trie) Initialize(){
T.Root=nil
}
 
//fungsi create_node() atau membuat node baru, berguna untuk membuat node baru dengan fungsi return node yang ditujukan untuk penggunaan node baru (untuk elemen etc)
 //bigO=0
 func Create_node() *node{
 //membuat element berisi node baru
	q:=new(node)
	//membuat inisialisasi data di node tersebut -1  
    q.data = -1
    return q
}
 
 //fungsi insert node dengan parameter input string
// fungsi ini akan tetap traversing, (q = q.link [index] bukan q = q.link) sampai kita dapatkan q.link [index] == nil (bukan q.link yang == nil)
// Ketika kita mendapatkan nil, maka bukannya menambahkan hanya 1 node dan membuat node sebelumnya menunjuk ke situ, kita buat  node baru yang banyak
// sebanyak nilai (panjang - level) pada waktu itu.
 //BigO=n(length si key)
func (T *Trie) Insert_node(key string) { 
//membuat elemen baru sebagai panjang input (key)
    length := len(key)
   var index int
   level:=0
   //pengecekan jika root masih nil/null, maka root diisi dengan node baru 
    if T.Root == nil {
        T.Root = Create_node()
		}
		q:=new(node)
		q=T.Root //untuk insert dari setiap huruf string, kita harus memulainya dari root
 
    for ; level < length ; level++ {
	// Pada setiap level, menemukan indeks yang sesuai
         // Karakter (a-z = 0-26)
        index = int(key[level] - 'a')
		if index<26 {
        if q.link[index] == nil  {
		// Taruh nilai karakter ini kedalam q.link[index]
             // Dan membuat 1 node lagi, dan node ini akan menjadi petunjuk
// disaat p :=new(node) maka q.link[index]=p
            q.link[index] = Create_node() 
        }
 
        q = q.link[index];
		}
    }
//Sekarang, karakter terakhir (node) dari string akan berisi nilai dari key ini
    q.data = level  
}
 //fungsi search yang berguna untuk mencari dengan input parameter string
 //BigO=n (length nya si key)
func (T *Trie) Search(key string) {
//membuat variable ketemu sebagai boolean, berguna untuk tanda jika kata yang di cari ketemu, jika ketemu maka ketemu = true, dan mengeprint kata yang ketemu
var ketemu bool
//inisialisasi ketemu = false untuk mempermudah pertandaan pertemuan saat pencarian berlangsung
ketemu=false
q:=new(node)
q=T.Root
 length :=len(key)
level:=0
    for ;level < length;level++ {
        index := key[level] - 'a';
		if index<26 {
        if q.link[index] != nil {
            q = q.link[index]
			
			} else {
            break
			}
			}
    }
	//setelah loop pencarian selesai, sekarang melihat kondisi
	//jika data element baru (q) tidak samadengan -1, maka tandanya kata yang kamu cari itu ketemu!
	//lalu boolean ketemu akan menjadi true
if q.data != -1 {
ketemu=true
}
//jika ketemu true, maka print found beserta key yang di cari
    if ketemu==true {
	fmt.Println("Found! = ", key)

		} else { //elsenya (jika ketemu false maka mengeprint  No match found
		fmt.Println("No Match Found!")
		}
    
}


 
func main() {
trie:=new(Trie)
trie.Initialize()
//mencoba memasukan string kedalam trie
   trie.Insert_node("halo")
       trie.Insert_node("nama")
trie.Search("halo")
trie.Search("anjing")
trie.Search("halo")
 
}
