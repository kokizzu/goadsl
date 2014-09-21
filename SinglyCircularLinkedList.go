/*
    Singly Circular Linked List Interface
*/

package main

import "fmt"

type Any interface{}

type Node struct {
    Data Any
    NextLink *Node
}

// Return data: O(1)
func (n *Node) GetData() Any {
    return n.Data
}

// Set data: O(1)
func (n *Node) SetData(data Any) {
    n.Data = data
}

// Return next node: O(1)
func (n *Node) GetNext() *Node {
    return n.NextLink
}

// Set next node: O(1)
func (n *Node) SetNext(next *Node) {
    n.NextLink = next
}

// Test whether node has next node: O(1)
func (n *Node) HasNext() bool {
    if n.NextLink == nil { return false }
    return true
}

type SCLL struct {
    Head *Node
    count int
}

// Create a new Singly Circular Linked List: O(1)
func NewSCLL() *SCLL {
    return &SCLL{}
}

// Return size: O(1)
func (s *SCLL) Count() int {
    return s.count
}

// Return last node: O(N)
func (s *SCLL) GetTail() *Node {
    if s.Head == nil { return nil }
    current := s.Head
    for current.GetNext() != s.Head {
        current = current.GetNext()
    }
    return current
}

// Insert as head, set old head as head's prev: O(1)
func (s *SCLL) InsertPrev(newnode *Node) {
    if s.Head == nil {
        s.Head = newnode
        newnode.SetNext(newnode)
        s.count++
        return
    }
    next := s.Head.GetNext()
    newnode.SetNext(next)
    s.Head.SetNext(newnode)
    s.Head = newnode
    s.count++
}

// Insert as head, set old head as head's next: O(N)
func (s *SCLL) InsertNext(newnode *Node) {
    if s.Head == nil {
        s.Head = newnode
        newnode.SetNext(newnode)
        s.count++
        return
    }
    tail := s.GetTail()
    tail.SetNext(newnode)
    newnode.SetNext(s.Head)
    s.Head = newnode
    s.count++
}

// Delete head, set head's prev as new head: O(N)
func (s *SCLL) DeletePrev() {
    if s.Head != nil {
        if s.count == 1 {
            s.Head = nil
            s.count--
            return
        }
        tail := s.GetTail()
        tail.SetNext(s.Head.GetNext())
        s.Head = tail
        s.count--
    }
}

// Delete head, set head's next as new head: O(N)
func (s *SCLL) DeleteNext() {
    if s.Head != nil {
        if s.count == 1 {
            s.Head = nil
            s.count--
            return
        }
        tail := s.GetTail()
        tail.SetNext(s.Head.GetNext())
        s.Head = s.Head.GetNext()
        s.count--
    }
}

// Set head's next as new head: O(1)
func (s *SCLL) RotateLeft() {
    if s.Head != nil {
        s.Head = s.Head.GetNext()
    }
}

// Set head's prev as new head: O(N)
func (s *SCLL) RotateRight() {
    if s.Head != nil {
        s.Head = s.GetTail()
    }
}

// Print to left: O(N)
func (s *SCLL) PrintToLeft() {
    if s.Head == nil {
        fmt.Println("| Empty |")
        return
    }
    current := s.Head
    fmt.Printf("| %#v |", current.Data)
    defer fmt.Println()
    current = current.GetNext()
    for current != s.Head {
        defer fmt.Printf(" %#v |", current.Data)
        current = current.GetNext()
    }
}

// Print to right: O(N)
func (s *SCLL) PrintToRight() {
    if s.Head == nil {
        fmt.Println("| Empty |")
        return
    }
    current := s.Head
    fmt.Print("|")
    for {
        fmt.Printf(" %#v |", current.Data)
        current = current.GetNext()
        if current == s.Head { break }
    }
    fmt.Println()
}

func main() {
    s := NewSCLL()
    fmt.Println("Initial condition:")
    s.PrintToRight()

    fmt.Printf("\nInsert \"hello\" as head and set old head as head's prev\n")
    s.InsertPrev(&Node{"hello", nil})
    s.PrintToRight()
    fmt.Printf("\nInsert \"world\" as head and set old head as head's prev\n")
    s.InsertPrev(&Node{"world", nil})
    s.PrintToRight()
    fmt.Printf("\nInsert 3.14 as head and set old head as head's next\n")
    s.InsertNext(&Node{3.14, nil})
    s.PrintToRight()
    fmt.Printf("\nDelete head and set head's prev as new head\n")
    s.DeletePrev()
    s.PrintToRight()
    fmt.Printf("\nDelete head and set head's next as new head\n")
    s.DeleteNext()
    s.PrintToRight()
    fmt.Printf("\nDelete head and set head's next as new head\n")
    s.DeleteNext()
    s.PrintToRight()

    for i := 1; i <= 3; i++ {
        fmt.Printf("\nInsert %d as head and set old head as head's next\n",i)
        s.InsertNext(&Node{i, nil})
        s.PrintToRight()
    }

    for i := 4; i <= 6; i++ {
        fmt.Printf("\nInsert %d as head and set old head as head's prev\n",i)
        s.InsertPrev(&Node{i, nil})
        s.PrintToRight()
    }

    for i := 0; i < 2; i++ {
        fmt.Printf("\nRotate left:\n")
        s.RotateLeft()
        s.PrintToRight()
    }

    fmt.Printf("\nRotate right:\n")
    s.RotateRight()
    s.PrintToRight()

    fmt.Printf("\nPrint to right:\n")
    s.PrintToRight()

    fmt.Printf("\nPrint to left:\n")
    s.PrintToLeft()
}
