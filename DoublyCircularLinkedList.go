/*
    Doubly Circular Linked List Interface
*/

package main

import "fmt"

type Any interface{}

type Node struct {
    Data Any
    LeftLink *Node
    RightLink *Node
}

// Return data: O(1)
func (n *Node) GetData() Any {
    return n.Data
}

// Set data: O(1)
func (n *Node) SetData(data Any) {
    n.Data = data
}

// Return left node: O(1)
func (n *Node) GetLeft() *Node {
    return n.LeftLink
}

// Set left node: O(1)
func (n *Node) SetLeft(left *Node) {
    n.LeftLink = left
}

// Test whether node has left node: O(1)
func (n *Node) HasLeft() bool {
    if n.LeftLink == nil { return false }
    return true
}

// Return right node: O(1)
func (n *Node) GetRight() *Node {
    return n.RightLink
}

// Set right node: O(1)
func (n *Node) SetRight(right *Node) {
    n.RightLink = right
}

// Test whether node has right node: O(1)
func (n *Node) HasRight() bool {
    if n.RightLink == nil { return false }
    return true
}

type DCLL struct {
    Head *Node
    count int
}

// Create a new Doubly Circular Linked List: O(1)
func NewDCLL() *DCLL {
    return &DCLL{}
}

// Return size: O(1)
func (d *DCLL) Count() int {
    return d.count
}

// Insert as head, set old head as head's left: O(1)
func (d *DCLL) InsertLeft(newnode *Node) {
    if d.Head == nil {
        d.Head = newnode
        newnode.SetLeft(newnode)
        newnode.SetRight(newnode)
        d.count++
        return
    }
    right := d.Head.GetRight()
    newnode.SetRight(right)
    right.SetLeft(newnode)
    newnode.SetLeft(d.Head)
    d.Head.SetRight(newnode)
    d.Head = newnode
    d.count++
}

// Insert as head, set old head as head's right: O(1)
func (d *DCLL) InsertRight(newnode *Node) {
    if d.Head == nil {
        d.Head = newnode
        newnode.SetLeft(newnode)
        newnode.SetRight(newnode)
        d.count++
        return
    }
    left := d.Head.GetLeft()
    newnode.SetLeft(left)
    left.SetRight(newnode)
    newnode.SetRight(d.Head)
    d.Head.SetLeft(newnode)
    d.Head = newnode
    d.count++
}

// Delete head, set head's left as new head: O(1)
func (d *DCLL) DeleteLeft() {
    if d.Head != nil {
        if d.count == 1 {
            d.Head = nil
            d.count--
            return
        }
        left, right := d.Head.GetLeft(), d.Head.GetRight()
        left.SetRight(right)
        right.SetLeft(left)
        d.Head = left
        d.count--
    }
}

// Delete head, set head's right as new head: O(1)
func (d *DCLL) DeleteRight() {
    if d.Head != nil {
        if d.count == 1 {
            d.Head = nil
            d.count--
            return
        }
        left, right := d.Head.GetLeft(), d.Head.GetRight()
        left.SetRight(right)
        right.SetLeft(left)
        d.Head = right
        d.count--
    }
}

// Set head's right as new head: O(1)
func (d *DCLL) RotateLeft() {
    if d.Head != nil {
        d.Head = d.Head.GetRight()
    }
}

// Set head's left as new head: O(1)
func (d *DCLL) RotateRight() {
    if d.Head != nil {
        d.Head = d.Head.GetLeft()
    }
}

// Print to left: O(N)
func (d *DCLL) PrintToLeft() {
    if d.Head == nil {
        fmt.Println("| Empty |")
        return
    }
    current := d.Head
    fmt.Print("|")
    for {
        fmt.Printf(" %#v |", current.Data)
        current = current.GetLeft()
        if current == d.Head { break }
    }
    fmt.Println()
}

// Print to right: O(N)
func (d *DCLL) PrintToRight() {
    if d.Head == nil {
        fmt.Println("| Empty |")
        return
    }
    current := d.Head
    fmt.Print("|")
    for {
        fmt.Printf(" %#v |", current.Data)
        current = current.GetRight()
        if current == d.Head { break }
    }
    fmt.Println()
}

func main() {
    d := NewDCLL()
    fmt.Println("Initial condition:")
    d.PrintToRight()

    fmt.Printf("\nInsert \"hello\" as head and set old head as head's left\n")
    d.InsertLeft(&Node{"hello", nil, nil})
    d.PrintToRight()
    fmt.Printf("\nInsert \"world\" as head and set old head as head's left\n")
    d.InsertLeft(&Node{"world", nil, nil})
    d.PrintToRight()
    fmt.Printf("\nInsert 3.14 as head and set old head as head's right\n")
    d.InsertRight(&Node{3.14, nil, nil})
    d.PrintToRight()
    fmt.Printf("\nDelete head and set head's left as new head\n")
    d.DeleteLeft()
    d.PrintToRight()
    fmt.Printf("\nDelete head and set head's right as new head\n")
    d.DeleteRight()
    d.PrintToRight()
    fmt.Printf("\nDelete head and set head's right as new head\n")
    d.DeleteRight()
    d.PrintToRight()

    for i := 1; i <= 3; i++ {
        fmt.Printf("\nInsert %d as head and set old head as head's right\n",i)
        d.InsertRight(&Node{i, nil, nil})
        d.PrintToRight()
    }

    for i := 4; i <= 6; i++ {
        fmt.Printf("\nInsert %d as head and set old head as head's left\n",i)
        d.InsertLeft(&Node{i, nil, nil})
        d.PrintToRight()
    }

    for i := 0; i < 2; i++ {
        fmt.Printf("\nRotate left:\n")
        d.RotateLeft()
        d.PrintToRight()
    }

    fmt.Printf("\nRotate right:\n")
    d.RotateRight()
    d.PrintToRight()

    fmt.Printf("\nPrint to right:\n")
    d.PrintToRight()

    fmt.Printf("\nPrint to left:\n")
    d.PrintToLeft()
}
