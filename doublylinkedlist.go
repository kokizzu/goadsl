package main
import (
	"fmt"
	"errors"
	"strings"
)


type Value struct {
	Name string
	MilesAway int
}
type Node struct {
	Value				// Embedded struct
	next, prev  *Node
}
type List struct {
	head, tail *Node
}
func (l *List) First() *Node {
	return l.head
}
func (n *Node) Next() *Node {
	return n.next
}
func (n *Node) Prev() *Node {
	return n.prev
}
// Create new node with value
func (l *List) Push(v Value) *List {
	n := &Node{Value: v}
	if l.head == nil {
		l.head = n		// First node
	} else {
		l.tail.next = n	// Add after prev last node
		n.prev = l.tail // Link back to prev last node
	}
	l.tail = n  		// reset tail to newly added node
	return l
}
func (l *List) Find(name string) *Node {
	found := false
	var ret *Node = nil
	for n := l.First(); n != nil && !found; n = n.Next() {
		if n.Value.Name == name {
			found = true
			ret = n
		}
	}
	return ret
}
func (l *List) Delete(name string) bool {
	success := false
	node2del := l.Find(name)
	if node2del != nil {
		fmt.Println("Delete - FOUND: ", name)
		prev_node := node2del.prev
		next_node := node2del.next
		// Remove this node
		prev_node.next = node2del.next
		next_node.prev = node2del.prev
		success = true
	}
	return success
}
var errEmpty = errors.New("ERROR - List is empty")
// Pop last item from list
func (l *List) Pop() (v Value, err error) {
	if l.tail == nil {
		err = errEmpty
	} else {
		v = l.tail.Value
		l.tail = l.tail.prev
		if l.tail == nil {
			l.head = nil
		}
	}
	return v, err
}

func main() {
	dashes := strings.Repeat("-", 50)
	l := new(List) 
	l.Push(Value{Name: "Indonesia", MilesAway: 0})
	l.Push(Value{Name: "Medan", MilesAway: 1961})
	l.Push(Value{Name: "Kuala Namu", MilesAway: 881})

	processed := make(map[*Node]bool)

	fmt.Println("waktu pertama ")
	for n := l.First(); n != nil; n = n.Next() {
		fmt.Printf("%v\n", n.Value)
		if processed[n] {
			fmt.Printf("%s di proses\n", n.Value)
		}
		processed[n] = true
	}
	fmt.Println(dashes)
	fmt.Println("waktu kedua")
	for n := l.First(); n != nil; n = n.Next() {
		fmt.Printf("%v", n.Value)
		if processed[n] {
			fmt.Println(" diproses")
		} else { fmt.Println() }
		processed[n] = true
	}

	fmt.Println(dashes)
	var found_node *Node
	kota_dicari := "Kuala namu"
	found_node = l.Find(kota_dicari)
	if found_node != nil {
		fmt.Printf("tidak ditemukan: %v\n", kota_dicari)
	} else {
		fmt.Printf("ditemukan: %v\n", kota_dicari)
	}

	kota_dicari = "Pantai labu"
	found_node = l.Find(kota_dicari)
	if found_node == nil {
		fmt.Printf("tidak ditemukan: %v\n", kota_dicari)
	} else {
		fmt.Printf("ditemukan: %v\n", kota_dicari)
	}

	fmt.Println(dashes)
	kota_dihapus := "Siantar"
	successfully_removed_city := l.Delete(	kota_dihapus)
	if successfully_removed_city {
		fmt.Printf("dihapus: %v\n",kota_dihapus)
	} else {
		fmt.Printf("tidak dihapus: %v\n", 	kota_dihapus)
	}

		kota_dihapus= "Pantai labu"
	successfully_removed_city = l.Delete(	kota_dihapus)
	if successfully_removed_city {
		fmt.Printf("dihapus: %v\n", 	kota_dihapus)
	} else {
		fmt.Printf("tidak dihapus: %v\n", 	kota_dihapus)
	}

	fmt.Println(dashes)
	fmt.Println("Nilai")
	for v, err := l.Pop(); err == nil; v, err = l.Pop() {
		fmt.Printf("%v\n", v)
	}
	fmt.Println(l.Pop())  
}