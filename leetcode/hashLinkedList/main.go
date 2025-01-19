package main

import (
	"container/list"
	"fmt"
)

// HashLinkedList represents a hash-linked list data structure.
type HashLinkedList struct {
	table map[string]*list.Element // Hash table for fast lookups
	list  *list.List               // Doubly linked list for maintaining order
}

// KeyValue stores a key-value pair in the list.
type KeyValue struct {
	key   string
	value string
}

// NewHashLinkedList initializes a new hash-linked list.
func NewHashLinkedList() *HashLinkedList {
	return &HashLinkedList{
		table: make(map[string]*list.Element),
		list:  list.New(),
	}
}

// Add adds a new key-value pair or updates the value if the key exists.
func (hll *HashLinkedList) Add(key, value string) {
	if element, exists := hll.table[key]; exists {
		// Update the value if the key already exists
		element.Value.(*KeyValue).value = value
		return
	}

	// Add to the linked list
	newElement := hll.list.PushBack(&KeyValue{key: key, value: value})

	// Add to the hash table
	hll.table[key] = newElement
}

// Get retrieves a value by its key.
func (hll *HashLinkedList) Get(key string) (string, bool) {
	element, exists := hll.table[key]
	if !exists {
		return "", false
	}
	return element.Value.(*KeyValue).value, true
}

// Delete removes a key-value pair by key.
func (hll *HashLinkedList) Delete(key string) bool {
	element, exists := hll.table[key]
	if !exists {
		return false
	}

	// Remove from the linked list
	hll.list.Remove(element)

	// Remove from the hash table
	delete(hll.table, key)
	return true
}

// Display prints the elements of the list in order.
func (hll *HashLinkedList) Display() {
	for e := hll.list.Front(); e != nil; e = e.Next() {
		kv := e.Value.(*KeyValue)
		fmt.Printf("(%s: %s) -> ", kv.key, kv.value)
	}
	fmt.Println("nil")
}

func main() {
	hll := NewHashLinkedList()

	hll.Add("a", "apple")
	hll.Add("b", "banana")
	hll.Add("c", "cherry")

	fmt.Println("Initial list:")
	hll.Display()

	fmt.Println("Value for 'b':", hll.Get("b"))

	fmt.Println("Deleting 'b'")
	hll.Delete("b")
	hll.Display()

	fmt.Println("Adding 'b' back with new value 'blueberry'")
	hll.Add("b", "blueberry")
	hll.Display()
}
