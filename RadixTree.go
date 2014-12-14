//Radix Tree

package main
import "fmt"

type Interface interface{}

type Node struct {
	Key []byte
	Value Interface
	Link *Node
	Next *Node
	}

//Return a new node
//Big O = O(1)
func NewNode() *Node {
	return &Node{}
	}

//Find the length of the prefix from node's key with b
//Big O = O(N)
func (node *Node) Prefix(b []byte) int {
	for i := 0; i < len(b); i++ {
		if i == len(node.Key) || node.Key[i] != b[i] {
			return i
			}
		}
	return len(b)
	}

//Split the node's key with k as the separation point
//Big O = O(1)
func (node *Node) Split(k int) {
	newnode := NewNode()
	newnode.Key = node.Key[k:]
	newnode.Link = node.Link
	newnode.Value = node.Value
	node.Value = 0
	node.Link = newnode
	node.Key = node.Key[:k]
	}

//Merge the node with it's link
//Big O = O(1)
func (node *Node) Merge() {
	link := node.Link
	node.Key = append(node.Key,link.Key...)
	node.Value = link.Value
	node.Link = link.Link
	}

//Search helper
//Big O = O(N)
func (node *Node) search(b []byte) *Node {
	k := node.Prefix(b)
	if k == 0 {
		if node.Next == nil { return nil }
		return node.Next.search(b)
	} else if k == len(b) {
		return node
	} else if k == len(node.Key) {
		if node.Link == nil { return nil }
		return node.Link.search(b[k:])
	}
	return nil
	}

//Insert helper
//Big O = O(N)
func (node *Node) insert(b []byte, i Interface) *Node {
	k := node.Prefix(b)
	if k == 0 {
		if node.Next == nil {
			node.Next = NewNode()
			node.Next.Key = b
			node.Next.Value = i
	} else {
		node.Next = node.Next.insert(b,i)
		}
	} else if k < len(b) {
		if k < len(node.Key) {
		node.Split(k)
		}
	if node.Link == nil {
		node.Link = NewNode()
		node.Link.Key = b[k:]
		node.Link.Value = i
	} else {
		node.Link = node.Link.insert(b[k:],i)
		}
	} else {
		node.Value = i
		}
	return node
	}

//Delete helper
//Big O = O(N)
func (node *Node) delete(b []byte) *Node {
	k := node.Prefix(b)
	if k == 0 {
	if node.Next != nil {
		node.Next = node.Next.delete(b)
		}
	} else if k == len(b) {
		return node.Next
	} else if k == len(node.Key) {
		if node.Link != nil {
			node.Link = node.Link.delete(b[k:])
			if (node.Link != nil) && (node.Link.Next == nil) {
				node.Merge()
				}
			}
		}
	return node
	}

type RadixTree struct {
	Root *Node
	}

//Return new radix tree
//Big O = O(1)
func NewRadixTree() *RadixTree {
	return &RadixTree{NewNode()}
	}

//Search for str in the Tree
//Big O for Search O(N)
func (radix *RadixTree) Search(str string) (Interface, bool) {
	node := radix.Root.search(append([]byte(str),0))
	if node == nil { return 0, false }
	return node.Value, true
	}

//Insert str into the Tree with i as its value
//Big O for Insert O(N)
func (radix *RadixTree) Insert(str string, i Interface) {
	radix.Root.insert(append([]byte(str),0),i)
	}

//Delete str from the Tree
//Big O for Delete O(N)
func (radix *RadixTree) Delete(str string) {
	radix.Root.delete(append([]byte(str),0))
	}

func main() {
	radix := NewRadixTree()
	radix.Insert("Test",5)
	radix.Insert("Search",0.10)
	radix.Insert("TestSearch","abc")
	fmt.Print("Searching for 'Test': ")
	fmt.Println(radix.Search("Test"))
	fmt.Print("Searching for 'Search': ")
	fmt.Println(radix.Search("Search"))
	fmt.Print("Searching for 'TestSearch': ")
	fmt.Println(radix.Search("TestSearch"))
	radix.Delete("Test")
	fmt.Print("Searching for 'Test': ")
	fmt.Println(radix.Search("Test"))
	fmt.Print("Searching for 'Search': ")
	fmt.Println(radix.Search("Search"))
	fmt.Print("Searching for 'TestSearch': ")
	fmt.Println(radix.Search("TestSearch"))
	radix.Delete("TestSearch")
	fmt.Print("Searching for 'Search': ")
	fmt.Println(radix.Search("Search"))
	fmt.Print("Searching for 'TestSearch': ")
	fmt.Println(radix.Search("TestSearch"))
	fmt.Print("Searching for 'AkuStres': ")
	fmt.Println(radix.Search("AkuStres': ")) //Testcase for non-existing Key
	}
