package main

import "fmt"

type node struct {
	val  int
	next *node
}

type list struct {
	head *node
	tail *node
	size int
}

func newList() *list {
	return &list{}
}

func (s *list) push(v int) {

	n := &node{
		val: v,
	}

	s.size++
	//empty list
	if s.head == nil {
		s.head = n
		s.tail = n

	}

	s.tail.next = n
	s.tail = n

}

func (s *list) remove(v int) {

	if s.head.val == v {
		s.size--
		s.head = s.head.next
		return
	}
	cur := s.head
	p := cur.next
	for p != nil && p.val != v {
		cur = p
		p = p.next
	}

	if p == nil {
		return
	}

	cur.next = p.next
	s.size--

}

func (s *list) insert(v, after int) {
	p := s.head

	n := &node{
		val: v,
	}
	s.size++
	for i := 0; i < s.size-1; i++ {
		if p.val == after {
			n.next = p.next
			p.next = n
			return
		}

		p = p.next
	}

	s.tail.next = n
	s.tail = n

}
func (s *list) reverse() {

	if s.size <= 1 {
		return
	}

	prev := s.head
	cur := prev.next
	next := cur.next

	s.tail = s.head

	for {

		cur.next = prev

		if next == nil {
			break
		}

		prev = cur
		cur = next
		next = cur.next

	}

	s.head = cur

}

func (s *list) print() {

	p := s.head

	for i := 0; i < s.size; i++ {
		fmt.Print(p.val, "\t")
		p = p.next
	}

	fmt.Println()
	fmt.Println("head:", s.head.val, "tail:", s.tail.val, "size:", s.size)

}

func (s *list) len() int {
	return s.size
}

func main() {
	s := newList()

	for i := 0; i < 5; i++ {
		s.push(i)
	}

	s.print()
	s.remove(0)
	//s.print()
	s.reverse()
	s.push(20)
	s.print()

	s.insert(55, 3)
	s.print()
}
