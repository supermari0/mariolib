// Package linkedlist contains linked lists and operations on them
package linkedlist

// Linked list interfaces
type Node interface {
    NodeValue() interface{}
}

// Probably too many methods, but it's good extra practice
type LinkedList interface {
    // LLDeleteFirstMatch attempts to delete the first occurrence of a value,
    // returning true if a match was found and value deleted, false otherwise
    LLDeleteFirstMatch(interface{}) bool
    // LLDeleteAllOccurrences removes all occurrences of a value, returning
    // true if any matches were found and deleted, false otherwise
    LLDeleteAllOccurrences(interface{}) bool
    // LLFind returns a pointer to the first Node containing the matched value
    // if found, nil otherwise
    LLFind(interface{}) Node
    // LLNodeAt returns the node at the specified index in the list, nil if
    // the list is longer than the provided index
    LLNodeAt(int) *Node
    LLInsertAfter(*Node, interface{})
    LLInsertBefore(*Node, interface{})
    LLAppend(interface{})
    LLPrepend(interface{})
}

// Singly linked list implementation
type SinglyLinkedList struct {
    Head *SinglyLinkedListNode
}

type SinglyLinkedListNode struct {
    V interface{}
    Next *SinglyLinkedListNode
}

func (n *SinglyLinkedListNode) NodeValue() interface{} {
    if n == nil { return nil } // empty list case
    return n.V
}

// TODO Add tests for all functions below.
func (l *SinglyLinkedList) LLDeleteFirstMatch(v interface{}) bool {
    var previous *SinglyLinkedListNode
    current := l.Head

    for match := current.NodeValue(); current != nil; current = current.Next {
        if match == v {
            if previous == nil {
                // list head is a match; move head forward
                l.Head = current.Next
                return true
            }
            previous.Next = current.Next
            return true
        } else {
            previous = current
            // current = current.Next happens as part of loop iteration
        }
    }
    return false
}

func (l *SinglyLinkedList) LLDeleteAllOccurrences(v interface{}) bool {
    anyMatch := false
    for {
        deleted := l.LLDeleteFirstMatch(v)
        if deleted {
            anyMatch = true
        } else {
            break
        }
    }
    return anyMatch
}

// TODO Possibly use modified version of this in LLDeleteFirst
func (l *SinglyLinkedList) LLFind(v interface{}) Node {
    current := l.Head
    for {
        currentVal := current.NodeValue()
        if currentVal == nil {
            return nil
        }
        if currentVal == v {
            return current
        }
        current = current.Next
    }
}

func (l *SinglyLinkedList) LLNodeAt(index int) Node {
    current := l.Head
    for i := 0; i < index; i++ {
        if current == nil {
            return nil
        }
        current = current.Next
    }
    return current
}

// TODO Handle case where node is not part of list?
func (l *SinglyLinkedList) LLInsertAfter(n *SinglyLinkedListNode, v interface{}) {
    newNode := SinglyLinkedListNode{v, n.Next}
    n.Next = &newNode
}

func (l *SinglyLinkedList) LLInsertBefore(n *SinglyLinkedListNode, v interface{}) {
    // basically the same as DeleteFirstMatch with different pointer ops
    var previous *SinglyLinkedListNode
    current := l.Head

    for match := current; current != nil; current = current.Next {
        if match == n {
            newNode := SinglyLinkedListNode{v, n}
            if previous == nil {
                // at head of list, reset head
                l.Head = &newNode
            } else {
                previous.Next = &newNode
            }
            return
        }
    }
}

func (l *SinglyLinkedList) LLAppend(v interface{}) {
    newNode := SinglyLinkedListNode{v, nil}
    var previous *SinglyLinkedListNode
    var current *SinglyLinkedListNode

    // traverse entire list
    for current = l.Head; current != nil; current = current.Next {
        previous = current
    }
    if previous == nil {
        // this is the head of the list
        l.Head = &newNode
        return
    }
    current.Next = &newNode
}

func (l *SinglyLinkedList) LLPrepend(v interface{}) {
    newNode := SinglyLinkedListNode{v, l.Head}
    l.Head = &newNode
}
