package main

import (
	"container/list"
	"fmt"
)

// LRUCache represents an LRU cache.
type LRUCache struct {
	capacity int                      // Maximum capacity of the cache
	table    map[string]*list.Element // Hash map for O(1) lookups
	list     *list.List               // Doubly linked list to maintain access order
}

// KeyValue stores a key-value pair in the list.
type KeyValue struct {
	key   string
	value string
}

// NewLRUCache initializes a new LRU cache.
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		table:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

// Get retrieves a value from the cache and marks it as recently used.
func (cache *LRUCache) Get(key string) (string, bool) {
	if element, exists := cache.table[key]; exists {
		// Move the accessed element to the front of the list
		cache.list.MoveToFront(element)
		return element.Value.(*KeyValue).value, true
	}
	return "", false
}

// Put adds a key-value pair to the cache or updates an existing key.
// If the cache exceeds its capacity, it evicts the least recently used item.
func (cache *LRUCache) Put(key, value string) {
	if element, exists := cache.table[key]; exists {
		// Update the value if the key already exists
		element.Value.(*KeyValue).value = value
		// Move the element to the front of the list
		cache.list.MoveToFront(element)
		return
	}

	// If the cache is at capacity, evict the least recently used item
	if cache.list.Len() >= cache.capacity {
		evict := cache.list.Back() // Least recently used item is at the back
		cache.list.Remove(evict)
		delete(cache.table, evict.Value.(*KeyValue).key)
	}

	// Add the new key-value pair
	newElement := cache.list.PushFront(&KeyValue{key: key, value: value})
	cache.table[key] = newElement
}

// Display prints the elements of the cache in access order.
func (cache *LRUCache) Display() {
	fmt.Print("Cache state: ")
	for e := cache.list.Front(); e != nil; e = e.Next() {
		kv := e.Value.(*KeyValue)
		fmt.Printf("(%s: %s) ", kv.key, kv.value)
	}
	fmt.Println()
}

func main() {
	// Create an LRU cache with capacity 3
	cache := NewLRUCache(3)

	cache.Put("a", "apple")
	cache.Put("b", "banana")
	cache.Put("c", "cherry")
	cache.Display()

	a, _ := cache.Get("a")
	// Access "a" to mark it as recently used
	fmt.Println("Get 'a':", a)
	cache.Display()

	// Add a new item, evicting the least recently used ("b")
	cache.Put("d", "date")
	cache.Display()

	// Add another item, evicting the least recently used ("c")
	cache.Put("e", "elderberry")
	cache.Display()

	d, _ := cache.Get("d")
	// Access "d" to mark it as recently used
	fmt.Println("Get 'd':", d)
	cache.Display()

	// Add another item, evicting the least recently used ("a")
	cache.Put("f", "fig")
	cache.Display()
}
