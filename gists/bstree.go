package main

import (
	"fmt"
)

type node struct {
	val   int
	left  *node
	right *node
}

type BSTree struct {
	root *node
}

func newBSTree() *BSTree {
	return &BSTree{}
}

func (t *BSTree) insertNode(root, node *node) bool {

	if root.val == node.val {
		return false
	}

	if root.val > node.val {
		if root.left == nil {
			root.left = node
			return true
		}
		return t.insertNode(root.left, node)
	}

	if root.right == nil {
		root.right = node
		return true
	}

	return t.insertNode(root.right, node)
}

func (t *BSTree) traverse(f func(*node)) {
	if t.root == nil {
		return
	}

	t.inOrderTraverse(t.root, f)

}

func (t *BSTree) inOrderTraverse(root *node, f func(*node)) {
	if root == nil {
		return
	}

	t.inOrderTraverse(root.left, f)
	f(root)
	t.inOrderTraverse(root.right, f)
}

func (t *BSTree) min() *node {
	if t.root == nil {
		return nil
	}

	node := t.root
	for node.left != nil {
		node = node.left
	}
	return node
}

func (t *BSTree) max() *node {
	if t.root == nil {
		return nil
	}

	node := t.root
	for {
		if node.right == nil {
			return node
		}

		node = node.right
	}
}

func (t *BSTree) print() {
	fmt.Println("------------------------------------------------")
	stringify(t.root, 0)
	fmt.Println("------------------------------------------------")
}

// internal recursive function to print a tree
func stringify(n *node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		fmt.Printf(format+"%d\n", n.val)
		stringify(n.right, level)
	}
}

func (t *BSTree) searchHelper(n *node, val int) bool {
	if n == nil {
		return false
	}

	if n.val == val {
		return true
	}

	if n.val > val {
		return t.searchHelper(n.left, val)
	}

	return t.searchHelper(n.right, val)
}

func (t *BSTree) search(val int) bool {
	if t.root == nil {
		return false
	}

	return t.searchHelper(t.root, val)

}
func (t *BSTree) insert(v int) bool {

	node := &node{
		val: v,
	}

	if t.root == nil {
		t.root = node
		return true
	}

	return t.insertNode(t.root, node)

}

func (t *BSTree) remove(val int) {
	t.root = t.removeHelper(t.root, val)  //update the root
}

func (t *BSTree) removeHelper(n *node, val int) *node {
	if n == nil {
		return nil
	}

	if val < n.val {
		n.left = t.removeHelper(n.left, val)
		return n
	}

	if val > n.val {
		n.right = t.removeHelper(n.right, val)
		return n
	}

	if n.left == nil && n.right == nil {
		return nil
	}

	if n.left == nil {
		return n.right
	}

	if n.right == nil {
		return n.left
	}

	p := n.left

	for p.right != nil {
		p = p.right
	}

	n.val = p.val

	n.left = t.removeHelper(n.left, n.val)

	return n

}

func fillTree(bst *BSTree) {
	bst.insert(8)
	bst.insert(4)
	bst.insert(10)
	bst.insert(2)
	bst.insert(6)
	bst.insert(1)
	bst.insert(3)
	bst.insert(5)
	bst.insert(7)
	bst.insert(9)
}

func main() {
	t := newBSTree()
	fillTree(t)
	t.insert(11)
	//t.print()

	fmt.Println(t.min().val, t.max().val)

	fmt.Println(t.search(10), t.search(12))

	t.remove(8)
	t.traverse(func(n *node) {
		fmt.Print(n.val, "\t")
	})
	fmt.Println()

}
