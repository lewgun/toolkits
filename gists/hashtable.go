package main

import "fmt"

var (
	minLoadFactor = 0.25
	maxLoadFactor = 0.75
	defaultTableSize = 3
)

type Record struct {
	key int
	value int
	next *Record
}

type Hash struct {
	records []*Record
}

type HashTable struct {
	table *Hash
	nRecords *int
}

// createHashTable: Called by checkLoadFactorAndUpdate when creating a new hash, for internal use only.
func createHashTable(tableSize int) HashTable {
	num := 0
	hash := Hash{make([]*Record, tableSize)}
	return HashTable{table: &hash, nRecords:&num}
}

// CreateHashTable: Called by the user to create a hashtable.
func CreateHashTable() HashTable {
	num := 0
	hash := Hash{make([]*Record, defaultTableSize)}
	return HashTable{table: &hash, nRecords:&num}
}

// hashFunction: Used to calculate the index of record within the slice
func hashFunction(key int, size int) (int) {
	return key%size
}

// Display: Print the hashtable in a legible format (publicly callable)
func (h *HashTable) Display() {
	fmt.Printf("----------%d elements-------\n", *h.nRecords)
	for i, node := range h.table.records {
		fmt.Printf("%d :", i)
		for node != nil {
			fmt.Printf("[%d, %d]->", node.key, node.value)
			node = node.next
		}
		fmt.Println("nil")
	}
}

// put: inserts a key into the hash table, for internal use only
func (h *HashTable) put(key int, value int) (bool){
	index := hashFunction(key, len(h.table.records))
	iterator := h.table.records[index]
	node := Record{key, value, nil}
	if iterator == nil {
		h.table.records[index] = &node
	} else {
		prev := &Record{0, 0, nil}
		for iterator != nil {
			if iterator.key == key { // Key already exists
				iterator.value = value
				return false
			}
			prev = iterator
			iterator = iterator.next
		}
		prev.next = &node
	}
	*h.nRecords += 1
	return true
}

// Put: inserts a key into the hash table (publicly callable)
func (h *HashTable) Put(key int, value int) {
	sizeChanged := h.put(key, value)
	if sizeChanged == true {
		h.checkLoadFactorAndUpdate()
	}
}

// Get: Retrieve a value for a key from the hash table (publicly callable)
func (h *HashTable) Get(key int) (bool, int) {
	index := hashFunction(key, len(h.table.records))
	iterator := h.table.records[index]
	for iterator != nil {
		if iterator.key == key {	// Key already exists
			return true, iterator.value
		}
		iterator = iterator.next
	}
	return false, 0
}

// del: remove a key-value record from the hash table, for internal use only
func (h *HashTable) del(key int) (bool) {
	index := hashFunction(key, len(h.table.records))
	iterator := h.table.records[index]
	if iterator == nil {
		return false
	}
	if iterator.key == key {
		h.table.records[index] = iterator.next
		*h.nRecords--
		return true
	} else {
		prev := iterator
		iterator = iterator.next
		for iterator != nil {
			if iterator.key == key {
				prev.next = iterator.next
				*h.nRecords--
				return true
			}
			prev = iterator
			iterator = iterator.next
		}
		return false
	}
}

// Del: remove a key-value record from the hash table (publicly available)
func (h *HashTable) Del(key int) (bool) {
	sizeChanged := h.del(key)
	if sizeChanged == true {
		h.checkLoadFactorAndUpdate()
	}
	return sizeChanged
}

// getLoadFactor: calculate the loadfactor for the hashtable
// Calculated as: number of records stored / length of underlying slice used
func (h *HashTable) getLoadFactor() (float64) {
	return float64(*h.nRecords)/float64(len(h.table.records))
}

// checkLoadFactorAndUpdate: if 0.25 > loadfactor or 0.75 < loadfactor,
// update the underlying slice to have have loadfactor close to 0.5
func (h *HashTable) checkLoadFactorAndUpdate() {
	if *h.nRecords == 0 {
		return
	} else {
		loadFactor := h.getLoadFactor()
		if loadFactor < minLoadFactor {
			fmt.Println("** Loadfactor below limit, reducing hashtable size **")
			hash := createHashTable(len(h.table.records)/2)
			for _, record := range h.table.records {
				for record != nil {
					hash.put(record.key, record.value)
					record = record.next
				}
			}
			h.table = hash.table
		} else if loadFactor > maxLoadFactor {
			fmt.Println("** Loadfactor above limit, increasing hashtable size **")
			hash := createHashTable(*h.nRecords*2)
			for _, record := range h.table.records {
				for record != nil {
					hash.put(record.key, record.value)
					record = record.next
				}
			}
			h.table = hash.table
		}
	}
}
