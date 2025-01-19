package main

import (
	"container/list"
	"fmt"
)

type LFUCache struct {
	Capacity   int
	keyValue   map[int]*list.Element
	freqToKeys map[int]*list.List
	leastFreq  int
}

type Pair struct {
	Key   int
	Value int
	Freq  int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		Capacity:   capacity,
		keyValue:   make(map[int]*list.Element),
		freqToKeys: make(map[int]*list.List),
		leastFreq:  1,
	}
}

func (this *LFUCache) Get(key int) int {
	element := this.keyValue[key]
	if element == nil {
		return -1
	}
	this.UpdateFreq(element)
	return element.Value.(*Pair).Value
}

func (this *LFUCache) Put(key int, value int) {
	element := this.keyValue[key]
	if element == nil {
		this.AddNewElement(key, value)
	} else {
		pair := element.Value.(*Pair)
		pair.Value = value
		this.UpdateFreq(element)
	}
}

func (this *LFUCache) AddNewElement(key, value int) *list.Element {
	this.CheckOverflow()
	l := this.GetOrDeafultList(1)
	element := l.PushBack(&Pair{key, value, 1})
	this.keyValue[key] = element
	this.leastFreq = 1
	return element
}

func (this *LFUCache) CheckOverflow() {
	if len(this.keyValue) >= this.Capacity {
		this.RemoveLFU()
	}
}

func (this *LFUCache) UpdateFreq(element *list.Element) {
	this.RemoveFromList(element)
	pair := element.Value.(*Pair)
	pair.Freq++
	this.AddToList(element)
}

func (this *LFUCache) AddToList(element *list.Element) {
	pair := element.Value.(*Pair)
	l2 := this.GetOrDeafultList(pair.Freq)
	ne := l2.PushBack(pair)
	this.keyValue[pair.Key] = ne
}

func (this *LFUCache) GetOrDeafultList(freq int) *list.List {
	if _, ok := this.freqToKeys[freq]; !ok {
		this.freqToKeys[freq] = list.New()
	}
	return this.freqToKeys[freq]
}

func (this *LFUCache) RemoveFromList(element *list.Element) {
	if element == nil {
		return
	}
	pair := element.Value.(*Pair)
	l := this.freqToKeys[pair.Freq]
	l.Remove(element)
	if l.Len() == 0 {
		delete(this.freqToKeys, pair.Freq)
		if this.leastFreq == pair.Freq {
			this.leastFreq++
		}
	}
}

func (this *LFUCache) RemoveLFU() {
	l := this.freqToKeys[this.leastFreq]
	element := l.Front()
	this.RemoveFromList(element)
	delete(this.keyValue, element.Value.(*Pair).Key)
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {

	lfu := Constructor(2)

	lfu.Put(1, 1)
	lfu.Put(2, 2)
	lfu.Put(2, 2)
	fmt.Println(lfu.Get(2))
	fmt.Println(lfu.Get(1))
	lfu.Put(2, 2)
	fmt.Println(lfu.Get(2))
	lfu.Put(3, 3)
	fmt.Println(lfu.Get(3))
	fmt.Println(lfu.Get(2))
	fmt.Println(lfu.Get(1))

}
