// Package linkedlist contains linked lists and operations on them
package linkedlist

// Linked list interfaces
type Node interface {
    NodeValue() interface{}
}

type LinkedList interface {
    // LLDeleteFirstMatch attempts to delete the first occurrence of a value,
    // returning true if a match was found and value deleted, false otherwise
    LLDeleteFirstMatch(interface{}) bool
    LLDeleteAllOccurrences(interface{}) bool
    LLFind(interface{}) *Node
    LLValueAt(int) *Node
    LLAppend(interface{})
    LLPrepend(interface{})
    LLInsertAfter(*Node, interface{})
    LLInsertBefore(*Node, interface{})
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

func (l *SinglyLinkedList) LLDeleteFirstMatch(v interface{}) bool {
    var previous *SinglyLinkedListNode
    previous = nil
    current := l.Head

    for match := current.NodeValue(); current != nil;  {
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
            current = current.Next
        }
    }
    return false
}
