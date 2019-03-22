package tests

import (
    "testing"
    ds "github.com/supermari0/mariolib/golang/datastructures"
)

func TestSinglyLinkedListDelete(t *testing.T) {
    node := ds.SinglyLinkedListNode{
        1,
        nil,
    }
    ll := ds.SinglyLinkedList{
        &node,
    }

    deleted := ll.LLDeleteFirstMatch(1)
    if !deleted {
        t.Errorf("Failed deletion in singly linked list")
    }

    deleted = ll.LLDeleteFirstMatch(1)
    if deleted {
        t.Errorf("Delete should not succeed twice on the same element")
    }
}

func TestSinglyLinkedListInsert(t *testing.T) {
    ll := ds.SinglyLinkedList{nil}

    node := ds.SinglyLinkedListNode{1, nil}
    ll.LLAppend(&node)
    // TODO finish test
}
