package datastructures

type TreeNode interface {
    TreeNodeValue() interface{}
    // []TreeNode slice is a pointer to the underlying array, so we don't have
    // to return a reference to the slice itself
    TreeChildren() []*TreeNode
}

type GenericTree struct {
    V interface{}
    Children []*GenericTree
}

func (t *GenericTree) TreeNodeValue() interface{} {
    return t.V
}

func (t *GenericTree) TreeChildren() []*GenericTree {
    return t.Children
}

// Note: This can be rewritten with AddNChild
func (t *GenericTree) AddLeftChild(child *GenericTree) {
    t.AddNChild(0, child)
    /* Old implementation - rewrote with AddNChild
    childSlice := make([]*GenericTree, 1)
    childSlice[0] = child
    newChildren := append(childSlice, t.Children...)
    t.Children = newChildren */
}

func (t *GenericTree) AddRightChild(child *GenericTree) {
    // TODO Initialize an empty tree correctly
    newChildren := append(t.Children, child)
    t.Children = newChildren
}

// Insert a child in position i in the tree.
// If the position requested is larger than the current number of children, the
// child is added in the last position.
func (t *GenericTree) AddNChild(i uint, child *GenericTree) {
    if i > uint(len(t.Children)) {
        t.AddRightChild(child)
        return
    }
    // TODO Bounds checking? Also, initialize an empty tree correctly
    previousChildren := t.Children[0:i]
    nextChildren := t.Children[i+1:]
    newChildren := append(previousChildren, child, nextChildren...)
    t.Children = newChildren
}
