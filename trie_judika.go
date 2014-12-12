package main

 import "fmt"

 
type node struct {
    data int
     link[26]*node
}
 
type Trie struct {
    Root *node
}

func (T *Trie) initialize(){
T.Root=nil
}
 
 func create_node() *node{
	q:=new(node)
    for x:=0; x<26; x++ {
        q.link[x] = nil
		}
    q.data = -1
    return q
}
 
func (T *Trie) insert_node(key string) {
    length := len(key)
   var index int
   level:=0
    if T.Root == nil {
        T.Root = create_node()
		}
		q:=new(node)
		q=T.Root
 
    for ; level < length ; level++ {
        index = int(key[level] - 'a')
		if index<26 {
        if q.link[index] == nil  {
            q.link[index] = create_node() 
        }
 
        q = q.link[index];
		}
    }

    q.data = level  
}
 
func (T *Trie) search(key string) int{
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

    if q.data != -1 {
	fmt.Println("Found! = ", key)
        return q.data
		} else {
		fmt.Println("No Match Found!")
		return -1 
		}
    return -1
}
 
func main() {
trie:=new(Trie)
trie.initialize()
var menu, kata string
for {
fmt.Println("-----------------------------------------------")
fmt.Println("Selamat datang di program Trie")
fmt.Println("-----------------------------------------------")
fmt.Println("Pilih Menu :")
fmt.Println("1. Insert")
fmt.Println("2. Search")
fmt.Printf("Masukan = ")
fmt.Scanln(&menu)
if menu=="1" {
fmt.Println("Masukan Kata yang akan di insert ")
fmt.Scanln(&kata)
trie.insert_node(kata)
continue
} else if menu=="2" {
fmt.Println("Masukan Kata yang akan di cari ")
fmt.Scanln(&kata)
trie.search(kata)
}
}
  
}