package datastructures

type BinaryTree interface {
    GetTreeNodeValue() interface{}
    // []TreeNode slice is a pointer to the underlying array, so we don't have
    // to return a reference to the slice itself
    GetTreeChildren() []*BinaryTree
    GetLeftChild() *BinaryTree
    GetRightChild() *BinaryTree
    SetLeftChild(*BinaryTree)
    SetRightChild(*BinaryTree)
}

type GenericBinaryTree struct {
    V interface{}
    LeftChild *GenericBinaryTree
    RightChild *GenericBinaryTree
}

func (t *GenericBinaryTree) GetTreeNodeValue() interface{} {
    return t.V
}

func (t *GenericBinaryTree) GetTreeChildren() []*GenericBinaryTree {
    childrenSlice := make([]*GenericBinaryTree, 2, 2)
    childrenSlice[0] = t.LeftChild
    childrenSlice[1] = t.RightChild
    return childrenSlice
}

func (t *GenericBinaryTree) GetLeftChild() *GenericBinaryTree {
    return t.LeftChild
}

func (t *GenericBinaryTree) GetRightChild() *GenericBinaryTree {
    return t.RightChild
}

func (t *GenericBinaryTree) SetLeftChild(c *GenericBinaryTree) {
    t.LeftChild = c
}

func (t *GenericBinaryTree) SetRightChild(c *GenericBinaryTree) {
    t.RightChild = c
}

// Perform a depth-first search of the tree, starting with the left side.
// Return a pointer to the subtree whose root is the value we're searching for.
func (t *GenericBinaryTree) DepthFirstSearch(v interface{}) *GenericBinaryTree {
    if t.V == v {
        return t
    }
    if t.LeftChild != nil {
        result := t.LeftChild.DepthFirstSearch(v)
        if result != nil {
            return result
        }
    }
    if t.RightChild != nil {
        return t.RightChild.DepthFirstSearch(v)
    }
    return nil
}

// TODO: BFS with queue, rotate, sorted tree, balanced tree
