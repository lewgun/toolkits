package main

import "fmt"

const (
	maxBuckets   = 30
	minBuckets   = 10
	expandFactor = 0.75
	shrinkFactor = 0.25
)

type node struct {
	k    int
	v    int
	next *node
}

type hashTable struct {
	buckets []node

	bucketNum int

	count int
}

func newHashTableHelper(buckets int) *hashTable {
	return &hashTable{
		bucketNum: buckets,
		buckets:   make([]node, buckets),
	}
}
func newHashTable() *hashTable {
	return newHashTableHelper(minBuckets)
}

func (h *hashTable) hash(k int) int {
	return k % h.bucketNum
}

func (h *hashTable) put(k, v int) {
	ok := h.putHelper(k, v)
	if !ok {
		return
	}

	 h.adjust(true)
}
func (h *hashTable) putHelper(k, v int) bool {

	index := h.hash(k)

	n := &node{
		k: k,
		v: v,
	}
	if h.buckets[index].next == nil {
		h.buckets[index].next = n
		h.count++
		return true

	}

	p := h.buckets[index].next

	//update
	for p != nil {
		if p.k == k {
			p.v = v
			return false
		}
	}
	n.next = h.buckets[index].next
	h.buckets[index].next = n
	h.count++

	return true

}

func (h *hashTable) adjust(isExpand bool)  {
	factor := float32(h.count) / float32(h.bucketNum)

	num := h.bucketNum
	if isExpand {
		// do nothing
		if factor <= expandFactor || h.bucketNum == maxBuckets {
			return
		}

		num *= 2
		if num > maxBuckets {
			num = maxBuckets
		}
		fmt.Println("++++expand the  hash table", num)

	} else {
		// do nothing
		if factor >= shrinkFactor || h.bucketNum == minBuckets {
			return
		}

		num /= 2
		if num < minBuckets {
			num = minBuckets
		}

		fmt.Println("----shrink the hash table", num)
	}


	ht := newHashTableHelper(num)

	for _, bucket := range h.buckets {
		for p := bucket.next; p != nil; p = p.next {
			ht.put(p.k, p.v)
		}
	}

	h.bucketNum = ht.bucketNum
	h.buckets = ht.buckets
	h.count = ht.count

}

func (h *hashTable) get(k int) (int, bool) {
	index := h.hash(k)
	if h.buckets[index].next == nil {
		return 0, false
	}

	p := h.buckets[index].next

	//update
	for p != nil {
		if p.k == k {
			return p.v, true
		}
	}

	return 0, false

}

func (h *hashTable) print() {
	for _, bucket := range h.buckets {
		for p := bucket.next; p != nil; p = p.next {
			fmt.Println(p.k, " => ", p.v)
		}
	}

	fmt.Println("count", h.count, "size:", h.bucketNum)
}

func (h *hashTable)delete(k int) {

	ok := h.deleteHelper(k)
	if !ok {
		return
	}
	h.adjust(false )
}
func (h *hashTable) deleteHelper(k int) bool  {
	index := h.hash(k)
	if h.buckets[index].next == nil {
		return  false
	}

	prev := &h.buckets[index]
	for p := prev.next; p != nil; p = p.next {
		if p.k == k {
				prev.next = p.next
				h.count--
				return true
		}
		prev = p
	}

	return false
}

func (h *hashTable) size() int  {
	return h.count
}

func main() {
	h := newHashTable()

	for i := 0; i < 10; i++ {
		h.put(i, i*10)
	}
	//h.print()

	fmt.Println(h.get(5))

	for i := 0; i < 10; i++ {
		h.delete(i)
	}
	h.print()

	fmt.Println(h.get(5))
}
