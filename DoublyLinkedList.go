/*
    Doubly Linked List Interface
*/

package main

import "fmt"

type Any interface{}

type Node struct {
    Data Any
    NextLink *Node
    PrevLink *Node
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

func (n *Node) GetPrev() *Node {
    return n.PrevLink
}

func (n *Node) SetPrev(prev *Node) {
    n.PrevLink = prev
}

func (n *Node) HasPrev() bool {
    if n.PrevLink == nil { return false }
    return true
}

type DLL struct {
    Head *Node
    Tail *Node
    count int
}

func NewDLL() *DLL {
    return &DLL{}
}

func (d *DLL) Count() int {
    return d.count
}

// Get Nth Node: O(N)
func (d *DLL) GetNth(n int) (nthnode *Node) {
    if (n >= d.count) {
        nthnode = nil
    } else if n <= (d.count-1-n) {
        nthnode = d.Head
        pos := 0
        for (nthnode != d.Tail) && (pos != n) {
            nthnode = nthnode.GetNext()
            pos++
        }
    } else {
        nthnode = d.Tail
        pos := d.count-1
        for (nthnode != d.Head) && (pos != n) {
            nthnode = nthnode.GetPrev()
            pos--
        }
    }
    return
}

// Insert to Head: O(1)
func (d *DLL) InsertFirst(newnode *Node) int {
    if d.Head != nil {
        d.Head.SetPrev(newnode)
        newnode.SetNext(d.Head)
    } else {
        d.Tail = newnode
    }
    d.Head = newnode
    d.count++
    return 0
}

// Insert to Tail: O(1)
func (d *DLL) InsertLast(newnode *Node) int {
    if d.Tail != nil {
        d.Tail.SetNext(newnode)
        newnode.SetPrev(d.Tail)
    } else {
        d.Head = newnode
    }
    d.Tail = newnode
    d.count++
    return d.count-1
}

// Insert to Nth Node: O(N)
func (d *DLL) InsertNth(newnode *Node, n int) (pos int) {
    if d.Head == nil {
        d.Head = newnode
        d.Tail = newnode
        pos = 0
    } else if n == 0 {
        d.Head.SetPrev(newnode)
        newnode.SetNext(d.Head)
        d.Head = newnode
        pos = n
    } else if n > (d.count-1) {
        d.Tail.SetNext(newnode)
        newnode.SetPrev(d.Tail)
        d.Tail = newnode
        pos = d.count-1
    } else {
        current := d.GetNth(n)
        prev := current.GetPrev()
        prev.SetNext(newnode)
        newnode.SetPrev(prev)
        newnode.SetNext(current)
        current.SetPrev(newnode)
        pos = n
    }
    d.count++
    return
}

// Delete Head: O(1)
func (d *DLL) DeleteFirst() {
    if d.count != 0 {
        if d.Head == d.Tail {
            d.Head = nil
            d.Tail = nil
        } else {
            d.Head = d.Head.GetNext()
            d.Head.SetPrev(nil)
        }
        d.count--
    }
}

// Delete Tail: O(1)
func (d *DLL) DeleteLast() {
    if d.count != 0 {
        if d.Tail == d.Head {
            d.Tail = nil
            d.Head = nil
        } else {
            d.Tail = d.Tail.GetPrev()
            d.Tail.SetNext(nil)
        }
        d.count--
    }
}

// Delete Nth Node: O(N)
func (d *DLL) DeleteNth(n int) {
    if d.count != 0 {
        if d.Head == d.Tail {
            d.Head = nil
            d.Tail = nil
        } else if n == 0 {
            d.Head = d.Head.GetNext()
            d.Head.SetPrev(nil)
        } else if n >= (d.count-1) {
            d.Tail = d.Tail.GetPrev()
            d.Tail.SetNext(nil)
        } else {
            current := d.GetNth(n)
            prev, next := current.GetPrev(), current.GetNext()
            prev.SetNext(next)
            next.SetPrev(prev)
        }
        d.count--
    }
}

// Delete All: O(1)
func (d *DLL) DeleteAll() {
    d.Head = nil
    d.Tail = nil
    d.count = 0
}

// Move Nth Node to Head: O(N)
func (d *DLL) MoveNthToFirst(n int) {
    if (d.count != 0) && (n < d.count) && (n != 0) {
        current := d.GetNth(n)
        if current == d.Tail {
            d.Tail = current.GetPrev()
            d.Tail.SetNext(nil)
        } else {
            prev, next := current.GetPrev(), current.GetNext()
            prev.SetNext(next)
            next.SetPrev(prev)
        }
        d.Head.SetPrev(current)
        current.SetNext(d.Head)
        d.Head = current
        d.Head.SetPrev(nil)
    }
}

// Move Nth Node to Tail: O(N)
func (d *DLL) MoveNthToLast(n int) {
    if (d.count != 0) && (n < d.count) && (n != (d.count-1)) {
        current := d.GetNth(n)
        if current == d.Head {
            d.Head = current.GetNext()
            d.Head.SetPrev(nil)
        } else {
            prev, next := current.GetPrev(), current.GetNext()
            prev.SetNext(next)
            next.SetPrev(prev)
        }
        d.Tail.SetNext(current)
        current.SetPrev(d.Tail)
        d.Tail = current
        d.Tail.SetNext(nil)
    }
}

// Move Nth Node to Nth Position: O(N)
func (d *DLL) MoveNthToNth(src, dest int) {
    if (d.count != 0) && (src < d.count) && (dest < d.count) && (src != dest) {
        srcnode := d.GetNth(src)
        destnode := d.GetNth(dest)
        if srcnode == d.Head {
            d.Head = srcnode.GetNext()
            d.Head.SetPrev(nil)
        } else if srcnode == d.Tail {
            d.Tail = srcnode.GetPrev()
            d.Tail.SetNext(nil)
        } else {
            srcprev, srcnext := srcnode.GetPrev(), srcnode.GetNext()
            srcprev.SetNext(srcnext)
            srcnext.SetPrev(srcprev)
        }
        if src < dest {
            if destnode == d.Tail {
                d.Tail = srcnode
                d.Tail.SetNext(nil)
            } else {
                destnext := destnode.GetNext()
                srcnode.SetNext(destnext)
                destnext.SetPrev(srcnode)
            }
            srcnode.SetPrev(destnode)
            destnode.SetNext(srcnode)
        } else {
            if destnode == d.Head {
                d.Head = srcnode
                d.Head.SetPrev(nil)
            } else {
                destprev := destnode.GetPrev()
                srcnode.SetPrev(destprev)
                destprev.SetNext(srcnode)
            }
            srcnode.SetNext(destnode)
            destnode.SetPrev(srcnode)
        }
    }
}

// Print From Head to Tail: O(N)
func (d *DLL) PrintHeadToTail() {
    if d.Head == nil {
        fmt.Println("| Empty |")
        return
    }
    current := d.Head
    pos := 0
    fmt.Print("|")
    for current != nil {
        fmt.Printf(" %d: %#v |", pos, current.Data)
        current = current.GetNext()
        pos++
    }
    fmt.Println()
}

// Print From Tail to Head: O(N)
func (d *DLL) PrintTailToHead() {
    if d.Tail == nil {
        fmt.Println("| Empty |")
        return
    }
    current := d.Tail
    pos := d.count-1
    fmt.Print("|")
    for current != nil {
        fmt.Printf(" %d: %#v |", pos, current.Data)
        current = current.GetPrev()
        pos--
    }
    fmt.Println()
}

func main() {
    d := NewDLL()
    fmt.Println("Kondisi Awal DLL:")
    d.PrintHeadToTail()                     // empty

    fmt.Println()
    fmt.Println("Masukan string world dari depan:")
    d.InsertFirst(&Node{"world", nil, nil}) // "world"
    d.PrintHeadToTail()


    fmt.Println()
    fmt.Println("Masukan string hello dari depan:")
    d.InsertFirst(&Node{"hello", nil, nil}) // "hello" "world"
    d.PrintHeadToTail()


    fmt.Println()
    fmt.Println("Masukan string ! dari belakang:")
    d.InsertLast(&Node{"!", nil, nil})      // "hello" "world" "!"
    d.PrintHeadToTail()


    fmt.Println()
    fmt.Println("Masukan integer(13) ke posisi kedua(1,ingat index mulai dari 0):")
    d.InsertNth(&Node{13, nil, nil}, 1)     // "hello" 13 "world" "!"
    d.PrintHeadToTail()


    fmt.Println()
    fmt.Println("Masukan float(3.14) ke posisi empat(3,ingat index mulai dari 0):")
    d.InsertNth(&Node{3.14, nil, nil},3)    // "hello" 13 "world" 3.14 "!"
    d.PrintHeadToTail()


    fmt.Println()
    fmt.Println("Cetak DLL dari depan ke belakang (head to tail):")
    d.PrintHeadToTail()                     // "hello" 13 "world" 3.14 "!"


    fmt.Println()
    fmt.Println("Cetak DLL dari belakang ke depan (tail to head):")
    d.PrintTailToHead()                     // "!" 3.14 "world" 13 "hello"


    fmt.Println()
    fmt.Println("Hapus yang paling depan(head):")
    d.DeleteFirst()                         // 13 "world" 3.14 "!"
    d.PrintHeadToTail()


    fmt.Println()
    fmt.Println("Hapus yang paling belakang(tail):")
    d.DeleteLast()                          // 13 "world" 3.14
    d.PrintHeadToTail()


    fmt.Println()
    fmt.Println("Hapus yang posisi kedua(1,ingat index mulai dari 0):")
    d.DeleteNth(1)                          // 13 3.14
    d.PrintHeadToTail()

    fmt.Println()
    fmt.Println("Hapus semuanya:")
    d.DeleteAll()                           // empty
    d.PrintHeadToTail()

    fmt.Println()
    fmt.Println("Masukkan 5 integer dari belakang:")
    for i := 0; i < 5; i++ {
        d.InsertLast(&Node{i, nil, nil})
        d.PrintHeadToTail()
    }                                       // 0 1 2 3 4
    fmt.Println()
    fmt.Println("Kondisi DLL sekarang:")
    d.PrintHeadToTail()

    fmt.Println()
    fmt.Println("Pindahkan posisi 4(3, index dari 0) ke depan")
    d.MoveNthToFirst(3)                     // 3 0 1 2 4
    d.PrintHeadToTail()

    fmt.Println()
    fmt.Println("Pindahkan posisi 3(2, index dari 0) ke belakang")
    d.MoveNthToLast(2)                      // 3 0 2 4 1
    d.PrintHeadToTail()

    fmt.Println()
    fmt.Println("Pindahkan posisi 1(0, index dari 0) ke posisi 3(2)")
    d.MoveNthToNth(0,2)                     // 0 2 3 4 1
    d.PrintHeadToTail()

    fmt.Println()
    fmt.Println("Pindahkan posisi 3(2, index dari 0) ke posisi 5(4)")
    d.MoveNthToNth(2,4)                     // 0 2 4 1 3
    d.PrintHeadToTail()

    fmt.Println()
    fmt.Println("Pindahkan posisi 5(4, index dari 0) ke posisi 1(0)")
    d.MoveNthToNth(4,1)                     // 0 3 2 4 1
    d.PrintHeadToTail()

    fmt.Println()
    fmt.Println("Pindahkan posisi 4(3, index dari 0) ke posisi 1(0)")
    d.MoveNthToNth(3,0)                     // 4 0 3 2 1
    d.PrintHeadToTail()

    fmt.Println()
    fmt.Println("Kondisi akhir DLL:")
    d.PrintHeadToTail()
}
