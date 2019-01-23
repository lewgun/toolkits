package main

import "fmt"

type stack struct {
	slice []int
}

func newStack( ) *stack {
	return &stack{
	}
}

func (s *stack)push(v int) {
	s.slice = append(s.slice, v)
}

func ( s *stack) pop() int {
	v := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]

	return v
}

func (s *stack) size() int {
	return len(s.slice)
}

func (s *stack) empty() bool {
	return len(s.slice) == 0
}

func (s *stack) print() {
	for _, v := range s.slice {
		fmt.Print(v, "\t")
	}
	fmt.Println()

}

type node struct {

	val int
	prev *node
	next *node
}
type dqueue struct {

	head *node
	tail *node
	size int
}

func newDqueue () *dqueue {

	return &dqueue{}
}

func (d *dqueue) pushFront(v int ) {

	n := &node{
		val:v,
	}

	d.size++
	if d.head == nil {
		d.head = n
		d.tail = n
		return
	}

	n.next = d.head
	d.head.prev = n
	d.head = n

}

func (d *dqueue) pushEnd(v int ) {
	n := &node{
		val:v,
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

func (d *dqueue) popFront() (int, bool) {
	if d.size == 0 {
		return 0, false
	}

	v := d.head.val
	d.head = d.head.next

	if d.size == 1 {
		d.tail = nil
	}

	d.size--
	return v, true


}

func (d *dqueue) popEnd()(int, bool ) {

	if d.size == 0 {
		return 0, false
	}


	v := d.tail.val
	d.tail = d.tail.prev

	if d.size == 1 {
		d.head = nil
	}

	d.size--
	return v, true

}

func (d *dqueue) remove(v int ) bool {

	p := d.head

	for p != nil {
		if p.val == v {
			p.prev.next = p.next
			p.next.prev = p.prev
			d.size--
			return true
		}
		p = p.next
	}

	return false

}


func (d *dqueue) len() int {

	return d.size
}


func (d *dqueue) print() {

	p := d.head

	for i := 0; i < d.size; i++{
		fmt.Print(p.val, "\t")
		p = p.next
	}

	fmt.Println()
}

type qstack struct {

	dq *dqueue
}


func newQStack() *qstack {
	return &qstack{
		dq: newDqueue(),
	}
}

func (s *qstack) push(v int ) {
	s.dq.pushEnd(v)

}

func (s *qstack) pop() int  {
	v, _ := s.dq.popEnd()
	return v
}

type squeue struct {

	spush *stack
	spop *stack
}

func newSQueue() *squeue {

	return &squeue{
		spush: newStack(),
		spop: newStack(),
	}
}

func (s *squeue)enque(v int ) {
	s.spush.push(v)
}

var count int
func (s *squeue)deque() int {

	if s.size() <= 0 {
		return -1
	}
	count++
	if s.spop.empty() {
		for !s.spush.empty() {
			v := s.spush.pop()
			s.spop.push(v)
		}
	}
	val :=  s.spop.pop()
	return val
}

func ( s *squeue) size() int {
	return s.spush.size() + s.spop.size()
}

func( s *squeue) print() {
	size := s.size()
	for i := 0; i < size; i++ {
		fmt.Print(s.deque(), ", ")
	}
	fmt.Println()
}

func main() {
	//// stack
	//s := newStack()
	//
	//for i := 0;i < 5; i++ {
	//	s.push(i)
	//}
	//s.print()
	//
	//for i := 0; i < 5; i++ {
	//	fmt.Print( s.pop(), "\t")
	//}
	//fmt.Println()



	//// dequeue
	//dq := newDqueue()
	//for i := 1; i < 5; i++ {
	//	dq.pushEnd(i)
	//}
	//
	//for i := 5; i < 10; i++ {
	//	dq.pushFront(i)
	//}
	//dq.print()
	//
	//for i := 1; i < 5; i++ {
	//	v, _ := dq.popFront()
	//	fmt.Print( v, "\t")
	//}
	//fmt.Println()
	//
	//size := dq.size
	//for i := 0; i < size; i++ {
	//	v, _ := dq.popEnd()
	//	fmt.Print( v, "\t")
	//}
	//fmt.Println()
	//
	//dq.remove(5)
	//dq.print()

	//// qstack
	//dqs := newQStack()
	//
	//for i := 0; i < 5; i++ {
	//	dqs.push(i)
	//}
	//
	//for i := 0; i < 5; i++ {
	//	fmt.Print(dqs.pop(), "\t")
	//}
	//fmt.Println()

	// squeue
	sq := newSQueue()

	for i := 0; i < 5; i++ {
		sq.enque(i)
	}
	sq.print()

	//for i := 0; i < 3; i++ {
	//	fmt.Print(sq.deque(), "\t")
	//}
	//fmt.Println()
	////sq.print()
	//
	//for i := 5;i < 10; i ++ {
	//	sq.enque(i)
	//}
	////sq.print()
	//
	//
	//size := sq.size()
	//for i := 0; i < size; i++ {
	//	fmt.Print(sq.deque(), "\t")
	//}
	//fmt.Println()
	//
	//for i := 10;i < 23; i ++ {
	//	sq.enque(i)
	//}
	//
	//s := sq.size()
	//for i := 0; i < s; i++ {
	//	fmt.Print(sq.deque(), "\t")
	//}
	//fmt.Println()
	//
	////sq.print()


}

