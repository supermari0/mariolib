package main

import (
    "fmt"
    ds "github.com/supermari0/mariolib/golang/datastructures"
)

func main() {
    n := ds.SinglyLinkedListNode{
        1,
        nil,
    }
    ll := ds.SinglyLinkedList{
        &n,
    }

    fmt.Printf("%t\n", ll.LLDeleteFirstMatch(1)) // true
    fmt.Printf("%t\n", ll.LLDeleteFirstMatch(1)) // false, already removed
}
