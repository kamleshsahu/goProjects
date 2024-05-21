package containerDesign

import "github.com/emirpasic/gods/maps/treemap"

type NumberContainers struct {
	valToKeys map[int]*treemap.Map
	keyToVal  map[int]int
}

func Constructor() NumberContainers {

	valToKeys := make(map[int]*treemap.Map)
	keyToVal := make(map[int]int)
	return NumberContainers{
		valToKeys,
		keyToVal,
	}
}

func (this *NumberContainers) Change(index int, number int) {
	if val, ok := this.keyToVal[index]; ok {
		this.valToKeys[val].Remove(index)
		delete(this.keyToVal, index)
	}
	this.keyToVal[index] = number
	_, exist := this.valToKeys[number]
	if !exist {
		this.valToKeys[number] = treemap.NewWithIntComparator()
	}

	this.valToKeys[number].Put(index, true)
}

func (this *NumberContainers) Find(number int) int {
	tree, exist := this.valToKeys[number]
	if !exist {
		return -1
	}
	key, _ := tree.Min()
	if key == nil {
		return -1
	}
	return key.(int)
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */
