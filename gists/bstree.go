package main

import (
	"fmt"
)


type dqNode struct {
	val  *btNode
	prev *dqNode
	next *dqNode
}

// dequeue implements the double-linkly queue
type dqueue struct {
	head *dqNode
	tail *dqNode
	size int
}

func newDqueue() *dqueue {
	return &dqueue{}
}

func (d *dqueue) push(v *btNode) {
	n := &dqNode{
		val: v,
	}

	d.size++
	if d.head == nil {
		d.head = n
		d.tail = n
		return
	}

	d.tail.next = n
	n.prev = d.tail
	d.tail = n
}

func (d *dqueue) pop() (*btNode, bool) {
	if d.size == 0 {
		return nil, false
	}

	v := d.head.val
	d.head = d.head.next

	if d.size == 1 {
		d.tail = nil
	}

	d.size--
	return v, true
}

func (d *dqueue) empty() bool {
	return d.size == 0
}

func (d *dqueue) front() *btNode {
	if d.head == nil {
		return nil
	}

	return d.head.val
}


type stack struct {
	nodes []*btNode
}

func newStack() *stack {
	return &stack{}
}

func (s *stack) push( n *btNode) {
	s.nodes = append(s.nodes, n)
}

func (s *stack) pop() *btNode {
	if len(s.nodes) <= 0 {
		return nil
	}
	n := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]
	return n
}

func (s *stack) peek() *btNode {
	if len(s.nodes) <= 0 {
		return nil
	}
	return s.nodes[len(s.nodes)-1]
}

func (s *stack)empty() bool {
	return len(s.nodes) == 0
}

type btNode struct {
	val   int
	left  *btNode
	right *btNode
}

type BSTree struct {
	root *btNode
}

func newBSTree() *BSTree {
	return &BSTree{}
}

func (t *BSTree) insertNode(root, node *btNode) bool {

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

func (t *BSTree) inOrderTraverse(f func(*btNode)) {
	if t.root == nil {
		return
	}

	t.inOrderTraverseHelper(t.root, f)

}

func (t *BSTree) inOrderTraverseHelper(root *btNode, f func(*btNode)) {
	if root == nil {
		return
	}

	t.inOrderTraverseHelper(root.left, f)
	f(root)
	t.inOrderTraverseHelper(root.right, f)
}

func (t *BSTree) min() *btNode {
	if t.root == nil {
		return nil
	}

	node := t.root
	for node.left != nil {
		node = node.left
	}
	return node
}


func (t *BSTree) bfs() {
	if t.root == nil {
		return
	}

	dq := newDqueue()

	dq.push(t.root)

	for !dq.empty() {
		n := dq.front()
		if n.left != nil {
			dq.push(n.left)
		}

		if n.right != nil {
			dq.push(n.right)
		}

		fmt.Print(n.val, "->")
		dq.pop()
	}
	fmt.Println()
}


func ( t *BSTree)preOrderTraverse(f func(*btNode)) {
	if t.root == nil {
		return
	}
	t.preOrderTraverseHelper(f)
}

func (t *BSTree) preOrderTraverseHelper(f func(*btNode)) {

	s := newStack()

	s.push(t.root)

	for !s.empty() {
		n := s.pop()
		fmt.Print(n.val, "->")
		if n.right != nil {
			s.push(n.right)
		}
		if n.left != nil {
			s.push(n.left)
		}

	}
	fmt.Println()
}
func (t *BSTree) max() *btNode {
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
func stringify(n *btNode, level int) {
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

func (t *BSTree) searchHelper(n *btNode, val int) bool {
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

	node := &btNode{
		val: v,
	}

	if t.root == nil {
		t.root = node
		return true
	}

	return t.insertNode(t.root, node)

}

func (t *BSTree) remove(val int) {
	t.root = t.removeHelper(t.root, val) //update the root
}

func (t *BSTree) removeHelper(n *btNode, val int) *btNode {
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

	t.bfs()
	t.preOrderTraverse(nil)
	//t.print()

	fmt.Println(t.min().val, t.max().val)

	fmt.Println(t.search(10), t.search(12))

	t.remove(8)
	t.inOrderTraverse(func(n *btNode) {
		fmt.Print(n.val, "\t")
	})
	fmt.Println()

}
