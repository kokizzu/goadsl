/*
    Singly Linked List Interface
*/

package main

import "fmt"

type Any interface{}

type Node struct {
    Data Any
    NextLink *Node
}

func (n *Node) GetData() Any {
    return n.Data
}

func (n *Node) SetData(data Any) {
    n.Data = data
}

func (n *Node) GetNext() *Node {
    return n.NextLink
}

func (n *Node) SetNext(next *Node) {
    n.NextLink = next
}

func (n *Node) HasNext() bool {
    if n.NextLink == nil { return false }
    return true
}

type SLL struct {
    Head *Node
    count int
}

func NewSLL() *SLL {
    return &SLL{}
}

func (s *SLL) Count() int {
    return s.count
}

// Insert to Head: O(1)
func (s *SLL) InsertFirst(newnode *Node) int {
    s.count++
    oldhead := s.Head
    newnode.SetNext(oldhead)
    s.Head = newnode
    return 0
}

// Insert to Tail: O(N)
func (s *SLL) InsertLast(newnode *Node) int {
    s.count++
    if s.Head == nil {
        s.Head = newnode
        return 0
    }
    current := s.Head
    for current.GetNext() != nil {
        current = current.GetNext()
    }
    current.SetNext(newnode)
    return s.count-1
}

// Insert to Nth Node: O(N)
func (s *SLL) InsertNth(newnode *Node, n int) int {
    s.count++
    if s.Head == nil {
        s.Head = newnode
        return 0
    }
    pos := 0
    current := s.Head
    var prev *Node
    for (current.GetNext() != nil) && (pos < n) {
        prev = current
        current = current.GetNext()
        pos++
    }
    if pos == 0 {
        oldhead := current
        newnode.SetNext(oldhead)
        s.Head = newnode
    } else if pos != n {
        current.SetNext(newnode)
    } else {
        prev.SetNext(newnode)
        newnode.SetNext(current)
    }
    return pos
}

// Delete Head: O(1)
func (s *SLL) DeleteFirst() {
    if s.Head != nil {
        s.count--
        s.Head = s.Head.GetNext()
    }
}

// Delete Tail: O(N)
func (s *SLL) DeleteLast() {
    if s.Head != nil {
        s.count--
        current := s.Head
        var prev *Node
        for current.GetNext() != nil {
            prev = current
            current = current.GetNext()
        }
        if prev == nil {
            s.Head = nil
        } else {
            prev.SetNext(nil)
        }
    }
}

// Delete Nth Node: O(N)
func (s *SLL) DeleteNth(n int) {
    if s.Head != nil {
        s.count--
        current := s.Head
        var prev *Node
        pos := 0
        for (current.GetNext() != nil) && (pos < n) {
            prev = current
            current = current.GetNext()
            pos++
        }
        if pos == 0 {
            s.Head = current.GetNext()
        } else if pos != n {
            prev.SetNext(nil)
        } else {
            prev.SetNext(current.GetNext())
        }
    }
}

// Delete All: O(1)
func (s *SLL) DeleteAll() {
    s.count = 0
    s.Head = nil
}

// Move Nth Node to Head: O(N)
func (s *SLL) MoveNthToFirst(n int) {
    if (s.Head != nil) && (n < s.count) {
        current := s.Head
        var prev *Node
        pos := 0
        for (current.GetNext() != nil) && (pos < n) {
            prev = current
            current = current.GetNext()
            pos++
        }
        if pos > 0 {
            next := current.GetNext()
            oldhead := s.Head
            current.SetNext(oldhead)
            s.Head = current
            prev.SetNext(next)
        }
    }
}

// Move Nth Node to Tail: O(N)
func (s *SLL) MoveNthToLast(n int) {
    if (s.Head != nil) && (n < s.count) {
        current := s.Head
        var prev *Node
        pos := 0
        for (current.GetNext() != nil) && (pos < n) {
            prev = current
            current = current.GetNext()
            pos++
        }
        lastnode := current
        for lastnode.GetNext() != nil {
            lastnode = lastnode.GetNext()
        }
        if current.GetNext() != nil {
            next := current.GetNext()
            lastnode.SetNext(current)
            current.SetNext(nil)
            if prev == nil {
                s.Head = next
            } else {
                prev.SetNext(next)
            }
        }
    }
}

// Move Nth Node to Nth Position: O(N)
func (s *SLL) MoveNthToNth(src, dest int) {
    if (s.Head != nil) && (src < s.count) && (dest < s.count) && (src != dest) {
        current := s.Head
        var srcnode, srcprev, destnode, destprev, prev *Node
        pos := 0
        for current != nil {
            if pos == src {
                srcnode = current
                srcprev = prev
            }
            if pos == dest {
                destnode = current
                destprev = prev
            }
            prev = current
            current = current.GetNext()
            pos++
        }
        if src < dest {
            if srcprev == nil {
                s.Head = srcnode.GetNext()
            } else {
                srcprev.SetNext(srcnode.GetNext())
            }
            srcnode.SetNext(destnode.GetNext())
            destnode.SetNext(srcnode)
        } else {
            srcprev.SetNext(srcnode.GetNext())
            srcnode.SetNext(destnode)
            if destprev == nil {
                s.Head = srcnode
            } else {
                destprev.SetNext(srcnode)
            }
        }
    }
}

// Print From Head to Tail: O(N)
func (s *SLL) Print() {
    if s.Head == nil {
        fmt.Println("| Empty |")
        return
    }
    current := s.Head
    pos := 0
    fmt.Print("|")
    for current != nil {
        fmt.Printf(" %d: %#v |", pos, current.Data)
        current = current.GetNext()
        pos++
    }
    fmt.Println()
}

func main() {
    s := NewSLL()
    fmt.Println("Kondisi Awal DLL:")
    s.Print()                           // empty

    fmt.Println()
    fmt.Println("Masukan integer(1) dari depan:")
    s.InsertFirst(&Node{1, nil})        // 1
    s.Print()

    fmt.Println()
    fmt.Println("Masukan string hello dari depan:")
    s.InsertFirst(&Node{"hello", nil})  // "hello" 1
    s.Print()

    fmt.Println()
    fmt.Println("Masukan string # dari belakang:")
    s.InsertLast(&Node{"#", nil})       // "hello" 1 "#"
    s.Print()

    fmt.Println()
    fmt.Println("Masukan float 13.9 dari belakang:")
    s.InsertLast(&Node{13.9, nil})      // "hello" 1 "#" 13.9
    s.Print()

    fmt.Println()
    fmt.Println("Masukan integer(14) ke posisi N yaitu 4(3, index mulai dari 0):")
    s.InsertNth(&Node{14, nil},3)       // "hello" 1 "#" 14 13.9
    s.Print()


    fmt.Println()
    fmt.Println("Hapus Head:")
    s.DeleteFirst()                     // 1 "#" 14 13.9
    s.Print()

    fmt.Println()
    fmt.Println("Hapus Tail:")
    s.DeleteLast()                      // 1 "#" 14
    s.Print()

    fmt.Println()
    fmt.Println("Hapus Posisi N yaitu 2(1, index mulai dari 0):")
    s.DeleteNth(1)                      // 1 14
    s.Print()

    fmt.Println()
    fmt.Println("Hapus semuanya:")
    s.DeleteAll()                       // empty
    s.Print()

    fmt.Println()
    fmt.Println("Masukan 5 integer dari belakang(tail):")
    for i := 0; i < 5; i++ {
        s.InsertLast(&Node{i, nil})
        s.Print()
    }   // 0 1 2 3 4

    fmt.Println()
    fmt.Println("Kondisi saat ini:")
    s.Print()

    fmt.Println()
    fmt.Println("Pindahkan posisi n yaitu 4(3, index mulai dari 0) ke depan:")
    s.MoveNthToFirst(3)                 // 3 0 1 2 4
    s.Print()


    fmt.Println()
    fmt.Println("Pindahkan posisi n yaitu 2(1, index mulai dari 0) ke tail/belakang:")
    s.MoveNthToLast(1)                  // 3 1 2 4 0
    s.Print()


    fmt.Println()
    fmt.Println("Pindahkan posisi n yaitu 3(2, index mulai dari 0) ke posisi n yaitu 5(4, index mulai dari 0):")
    s.MoveNthToNth(2,4)                 // 3 1 4 0 2
    s.Print()


    fmt.Println()
    fmt.Println("Pindahkan posisi n yaitu 4(3, index mulai dari 0) ke posisi n yaitu 2(1, index mulai dari 0):")
    s.MoveNthToNth(3,1)                 // 3 0 1 4 2
    s.Print()
}
