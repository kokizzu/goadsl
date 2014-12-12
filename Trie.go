//Trie implementation in go

package main

import "fmt"

type node struct {
	data  int
	value interface{}
	link  [26]*node
}

type Trie struct {
	Root *node
}

//fungsi inisialisasi, berguna untuk menginisiali sasi bahwa trie yang baru dibuat, rootnya hrs nil
//bigO=1
func (T *Trie) Initialize() {
	T.Root = nil
}

//fungsi create_node() atau membuat node baru, berguna untuk membuat node baru dengan fungsi return node yang ditujukan untuk penggunaan node baru (untuk elemen etc)
//bigO=1
func Create_node() *node {
	//membuat element berisi node baru
	q := new(node)
	//membuat inisialisasi data di node tersebut -1
	q.data = -1
	q.value = nil
	return q
}

//fungsi insert node dengan parameter input string
// fungsi ini akan tetap traversing, (q = q.link [index] bukan q = q.link) sampai kita dapatkan q.link [index] == nil (bukan q.link yang == nil)
// Ketika kita mendapatkan nil, maka bukannya menambahkan hanya 1 node dan membuat titik simpul sebelumnya untuk itu, kita buat  node baru yang banyak
// Sebagai nilai (panjang - level) pada waktu itu.
//BigO=n(length si key)
func (T *Trie) Insert_node(key string, value interface{}) {
	//membuat elemen baru sebagai panjang input (key)
	length := len(key)
	var index int
	level := 0
	//pengecekan jika root masih nil/null, maka root diisi dengan node baru
	if T.Root == nil {
		T.Root = Create_node()
	}
	q := new(node)
	q = T.Root //untuk insert dari setiap huruf string, kita harus memulainya dari root
	for ; level < length; level++ {
		// Pada setiap level, menemukan indeks yang sesuai

		index = int(key[level] - 'a')
		if index < 26 {
			if q.link[index] == nil {
				// Taruh nilai karakter ini kedalam q.link[index]
				// disaat p :=new(node) maka q.link[index]=p
				q.link[index] = Create_node()
			}

			q = q.link[index]
		}
	}

	//Sekarang, karakter terakhir (node) dari string akan berisi nilai dari key ini
	q.data = level
	q.value = value
}

//fungsi search yang berguna untuk mencari dengan input parameter string
//BigO=n (length nya si key)
func (T *Trie) Search(key string) (interface{}, bool) {
	//membuat variable ketemu sebagai boolean, berguna untuk tanda jika kata yang di cari ketemu, jika ketemu maka ketemu = true, dan mengeprint kata yang ketemu
	var ketemu bool
	ketemu = true
	q := new(node)
	q = T.Root
	length := len(key)
	level := 0
	for ; level < length; level++ {
		index := key[level] - 'a'
		if index < 26 {
			if q.link[index] != nil {
				q = q.link[index]
			} else {
				ketemu = false
				break
			}
		}
	}
	/*
	   //setelah loop pencarian selesai, sekarang melihat kondisi
	   //jika data element baru (q) tidak samadengan -1, maka tandanya kata yang kamu cari itu ketemu!
	   //lalu boolean ketemu akan menjadi true
	   if q.data != -1 {
	       ketemu = true
	   }
	*/
	if q.value == nil {
		ketemu = false
	}
	//jika ketemu true, maka print found beserta key yang di cari
	if ketemu == true {
		fmt.Println("Found! = ", key)
		return q.value, ketemu
	} else { //elsenya (jika ketemu false maka mengeprint  No match found
		fmt.Println("No Match Found!")
	}
	return q.value, ketemu
}

//fungsi erase,berfungsi untuk menghapus key dan value yang ada ditrie dengan kondisi tertentu(ditentukan oleh parameter input dari user
//BigO=n(length si key)
func (T *Trie) Erase(key string) {
	var ketemu bool
	ketemu = true
	q := new(node)
	q = T.Root
	length := len(key)
	level := 0
	for ; level < length; level++ {
		index := key[level] - 'a'
		if index < 26 {
			if q.link[index] != nil {
				q = q.link[index]
			} else {
				ketemu = false
				break
			}
		}
	}
	if ketemu == true {
		q.value = nil
	}
}

type coba struct {
	masuk interface{}
}

func main() {
	trie := new(Trie)
	trie.Initialize()
	baru := coba{666}
	//mencoba memasukan string kedalam trie dengan value integer
	trie.Insert_node("halo", 1)
	//mencoba memasukan string kedalam trie dengan value string
	trie.Insert_node("nama", "biar")
	//mencoba memasukan string kedalam trie dengan value struct
	trie.Insert_node("chikuso", baru)
	//mencoba mencari key yang tidak ada di trie
	valuenya, nemu := trie.Search("sialan")
	fmt.Println(nemu)
	fmt.Println(valuenya)
	//mencoba mencari key yang ada di trie
	valuenya, nemu = trie.Search("halo")
	fmt.Println(nemu)
	fmt.Println(valuenya)
	valuenya, nemu = trie.Search("nama")
	fmt.Println(nemu)
	fmt.Println(valuenya)
	valuenya, nemu = trie.Search("chikuso")
	fmt.Println(nemu)
	fmt.Println(valuenya)
	//mencoba menghapus key yang ada di trie
	trie.Erase("halo")
	//lalu kita coba mencari halo, apakah masih ada di trie atau tidak
	valuenya, nemu = trie.Search("halo")
	fmt.Println(nemu)
	fmt.Println(valuenya)
	//kita coba menghapus key yang tidak ada di trie
	trie.Erase("cobaajalah")
	//kita coba erase yang lain
	trie.Erase("chikuso")
	//kita lihat lagi apakah chikuso masih ada?!
	valuenya, nemu = trie.Search("chikuso")
	fmt.Println(nemu)
	fmt.Println(valuenya)
	//ternyata sudah tidak ada lagi pemirsa!
}
