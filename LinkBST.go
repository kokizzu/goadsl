/*
    Binary Search Tree using Linked List
*/

package main

import "fmt"

type Node struct {
    Data int
    Left,Right *Node
}

// Return data: O(1)
func (n *Node) GetData() int {
    return n.Data
}

// Set data: O(1)
func (n *Node) SetData(data int) {
    n.Data = data
}

// Return left child: O(1)
func (n *Node) GetLeft() *Node {
    return n.Left
}

// Set left child: O(1)
func (n *Node) SetLeft(left *Node) {
    n.Left = left
}

// Test whether node has left child: O(1)
func (n *Node) HasLeft() bool {
    if n.Left == nil { return false }
    return true
}

// Return right child: O(1)
func (n *Node) GetRight() *Node {
    return n.Right
}

// Set right child: O(1)
func (n *Node) SetRight(right *Node) {
    n.Right = right
}

// Test whether node has right child: O(1)
func (n *Node) HasRight() bool {
    if n.Right == nil { return false }
    return true
}

type BST struct {
    Root *Node
}

// Create a new Binary Search Tree: O(1)
func NewBST() *BST {
    return &BST{}
}

// Search for a value in BST: O(log n)
func (b *BST) Search(data int) bool {
    if b.Root == nil { return false }
    current := b.Root
    for {
        if data < current.GetData() {
            if current.HasLeft() {
                current = current.GetLeft()
            } else {
                break
            }
        } else if data > current.GetData(){
            if current.HasRight() {
                current = current.GetRight()
            } else {
                break
            }
        } else {
            return true
        }
    }
    return false
}

// Return a minimum value from a subtree in BST: O(log n)
func (b *BST) SearchMin(current *Node) int {
    if current.HasLeft() { return b.SearchMin(current.GetLeft()) }
    return current.GetData()
}

// Insert a new value to BST: O(log n)
func (b *BST) Insert(data int) {
    newnode := &Node{data,nil,nil}
    if b.Root == nil { b.Root = newnode; return; }
    current := b.Root
    for {
        if data < current.GetData() {
            if current.HasLeft() {
                current = current.GetLeft()
            } else {
                current.SetLeft(newnode)
                break
            }
        } else if data > current.GetData() {
            if current.HasRight() {
                current = current.GetRight()
            } else {
                current.SetRight(newnode)
                break
            }
        } else {
            break
        }
    }
}

// Delete a value from BST: O(log n)
func (b *BST) Delete(data int) {
    if b.Root == nil { return }
    if b.Root.GetData() == data {
        auxroot := &Node{0,b.Root,nil}
        b.delete(b.Root,auxroot,data)
        b.Root = auxroot.GetLeft()
    } else {
        b.delete(b.Root,nil,data)
    }
}

func (b *BST) delete(current, parent *Node, data int) {
    if data < current.GetData() {
        if current.GetLeft() == nil { return }
        b.delete(current.GetLeft(),current,data)
    } else if data > current.GetData() {
        if current.GetRight() == nil { return }
        b.delete(current.GetRight(),current,data)
    } else {
        if current.HasLeft() && current.HasRight() {
            current.SetData(b.SearchMin(current.GetRight()))
            b.delete(current.GetRight(),current,current.GetData())
        } else if parent.GetLeft() == current {
            if current.HasLeft() {
                parent.SetLeft(current.GetLeft())
            } else {
                parent.SetLeft(current.GetRight())
            }
        } else if parent.GetRight() == current {
            if current.HasRight() {
                parent.SetRight(current.GetRight())
            } else {
                parent.SetRight(current.GetLeft())
            }
        }
    }
}

// Print BST with pre-order traversal: O(n)
func (b *BST) Preorder() {
    if b.Root == nil { return }
    b.preorder(b.Root)
    fmt.Println()
}

func (b *BST) preorder(current *Node) {
    fmt.Print(current.GetData()," ")
    if current.HasLeft() { b.preorder(current.GetLeft()) }
    if current.HasRight() { b.preorder(current.GetRight()) }
}

// Print BST with in-order traversal: O(n)
func (b *BST) Inorder() {
    if b.Root == nil { return }
    b.inorder(b.Root)
    fmt.Println()
}

func (b *BST) inorder(current *Node) {
    if current.HasLeft() { b.inorder(current.GetLeft()) }
    fmt.Print(current.GetData()," ")
    if current.HasRight() { b.inorder(current.GetRight()) }
}

// Print BST with post-order traversal: O(n)
func (b *BST) Postorder() {
    if b.Root == nil { return }
    b.postorder(b.Root)
    fmt.Println()
}

func (b *BST) postorder(current *Node) {
    if current.HasLeft() { b.postorder(current.GetLeft()) }
    if current.HasRight() { b.postorder(current.GetRight()) }
    fmt.Print(current.GetData()," ")
}

// Print BST with BFS: O(n)
func (b *BST) BFSPrinting() {
    if b.Root == nil { return }
    q := make([]*Node,0,0)
    q = append(q,b.Root)
    for len(q) != 0 {
        front := q[0]
        q = q[1:]
        fmt.Print(front.GetData()," ")
        if front.HasLeft() { q = append(q,front.GetLeft()) }
        if front.HasRight() { q = append(q,front.GetRight()) }
    }
    fmt.Println()
}

func main() {
    b := NewBST()
    fmt.Println("Insert 8, 3, 10, 1, 6, 14, 4, 7, 13")
    b.Insert(8)
    b.Insert(3)
    b.Insert(10)
    b.Insert(1)
    b.Insert(6)
    b.Insert(14)
    b.Insert(4)
    b.Insert(7)
    b.Insert(13)

    fmt.Println()
    fmt.Print("Preorder: ")
    b.Preorder()
    fmt.Print("Inorder: ")
    b.Inorder()
    fmt.Print("Postorder: ")
    b.Postorder()
    fmt.Print("BFS: ")
    b.BFSPrinting()

    fmt.Println()
    for i := 1; i <= 20; i++ {
        fmt.Printf("Searching for %d: ",i)
        if b.Search(i) { fmt.Println("Found") } else { fmt.Println("Not found") }
    }

    fmt.Println()
    fmt.Println("Delete 8, 6, 4, 13")
    b.Delete(8)
    b.Delete(6)
    b.Delete(4)
    b.Delete(13)

    fmt.Println()
    fmt.Print("Preorder: ")
    b.Preorder()
    fmt.Print("Inorder: ")
    b.Inorder()
    fmt.Print("Postorder: ")
    b.Postorder()
    fmt.Print("BFS: ")
    b.BFSPrinting()

    fmt.Println()
    for i := 1; i <= 20; i++ {
        fmt.Printf("Searching for %d: ",i)
        if b.Search(i) { fmt.Println("Found") } else { fmt.Println("Not found") }
    }

}
