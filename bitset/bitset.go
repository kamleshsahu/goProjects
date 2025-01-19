package main

import (
	"bytes"
	"fmt"
	"math/bits"
	"sort"
)

func main() {
	ans := maxTotalReward([]int{28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	fmt.Println(ans)
}

func removeDubs(nums []int) []int {
	ans := make([]int, 0)

	for i, val := range nums {
		if i > 0 && val == nums[i-1] {
			continue
		}

		ans = append(ans, val)
	}
	return ans
}

func maxTotalReward(nums []int) int {
	sort.Ints(nums)
	nums = removeDubs(nums)
	buffer := make([]uint, 16)

	lastIdx := len(nums) - 1
	bdp := New(uint(nums[lastIdx]))
	bdp.Set(0)
	// mask := New(uint(2 * nums[len(nums)-1]))
	maxReward := nums[lastIdx]
	if lastIdx == 0 {
		return maxReward
	}
	nums = nums[:lastIdx]
	ans := 0
	for _, reward := range nums {

		// for j := lastReward; j < reward; j++ {
		// 	mask.Set(uint(j))
		// }

		// pv := mask.Intersection(bdp)

		//fmt.Println("pv", pv.DumpAsBits())
		// npv := New(uint(2 * nums[len(nums)-1]))
		//count1, is1 := bdp.NextSetMany(0, buffer)
		//fmt.Println(count1, is1)
		for count, is := bdp.NextSetMany(0, buffer); len(is) > 0 && is[len(is)-1] < uint(reward); count, is = bdp.NextSetMany(is[len(is)-1]+1, buffer) {
			for _, i := range is {
				bdp.Set(i + uint(reward))
				// ans = max(ans, int(i)+reward)
			}
			if count == 0 {
				break
			}
		}
		//fmt.Println("bdp", bdp.DumpAsBits())

		// bdp = bdp.Union(npv)
		//fmt.Println("dp", bdp.DumpAsBits())
		// lastReward = reward
	}

	for count, is := bdp.NextSetMany(0, buffer); len(is) > 0 && is[len(is)-1] < uint(maxReward); count, is = bdp.NextSetMany(is[len(is)-1]+1, buffer) {
		ans = max(ans, int(is[len(is)-1]))
		if count == 0 {
			break
		}
	}
	return ans + maxReward
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// New creates a new BitSet with a hint that length bits will be required
func New(length uint) (bset *BitSet) {
	defer func() {
		if r := recover(); r != nil {
			bset = &BitSet{
				0,
				make([]uint64, 0),
			}
		}
	}()

	bset = &BitSet{
		length,
		make([]uint64, wordsNeeded(length)),
	}

	return bset
}

// A BitSet is a set of bits. The zero value of a BitSet is an empty set of length 0.
type BitSet struct {
	length uint
	set    []uint64
}

// wordsNeeded calculates the number of words needed for i bits
func wordsNeeded(i uint) int {
	if i > (Cap() - wordSize + 1) {
		return int(Cap() >> log2WordSize)
	}
	return int((i + (wordSize - 1)) >> log2WordSize)
}

// the wordSize of a bit set
const wordSize = uint(64)

// the wordSize of a bit set in bytes
const wordBytes = wordSize / 8

// log2WordSize is lg(wordSize)
const log2WordSize = uint(6)

// allBits has every bit set
const allBits uint64 = 0xffffffffffffffff

// Cap returns the total possible capacity, or number of bits
func Cap() uint {
	return ^uint(0)
}

// Test whether bit i is set.
func (b *BitSet) Test(i uint) bool {
	if i >= b.length {
		return false
	}
	return b.set[i>>log2WordSize]&(1<<wordsIndex(i)) != 0
}

// wordsIndex calculates the index of words in a `uint64`
func wordsIndex(i uint) uint {
	return i & (wordSize - 1)
}

// Set bit i to 1, the capacity of the bitset is automatically
// increased accordingly.
// If i>= Cap(), this function will panic.
// Warning: using a very large value for 'i'
// may lead to a memory shortage and a panic: the caller is responsible
// for providing sensible parameters in line with their memory capacity.
func (b *BitSet) Set(i uint) *BitSet {
	if i >= b.length { // if we need more bits, make 'em
		b.extendSet(i)
	}
	b.set[i>>log2WordSize] |= 1 << wordsIndex(i)
	return b
}

// extendSet adds additional words to incorporate new bits if needed
func (b *BitSet) extendSet(i uint) {
	if i >= Cap() {
		panic("You are exceeding the capacity")
	}
	nsize := wordsNeeded(i + 1)
	if b.set == nil {
		b.set = make([]uint64, nsize)
	} else if cap(b.set) >= nsize {
		b.set = b.set[:nsize] // fast resize
	} else if len(b.set) < nsize {
		newset := make([]uint64, nsize, 2*nsize) // increase capacity 2x
		copy(newset, b.set)
		b.set = newset
	}
	b.length = i + 1
}

func (b *BitSet) Intersection(compare *BitSet) (result *BitSet) {
	panicIfNull(b)
	panicIfNull(compare)
	b, compare = sortByLength(b, compare)
	result = New(b.length)
	for i, word := range b.set {
		result.set[i] = word & compare.set[i]
	}
	return
}

// Union of base set and other set
// This is the BitSet equivalent of | (or)
func (b *BitSet) Union(compare *BitSet) (result *BitSet) {
	panicIfNull(b)
	panicIfNull(compare)
	b, compare = sortByLength(b, compare)
	result = compare.Clone()
	for i, word := range b.set {
		result.set[i] = word | compare.set[i]
	}
	return
}

func panicIfNull(b *BitSet) {
	if b == nil {
		panic(Error("BitSet must not be null"))
	}
}

// Clone this BitSet
func (b *BitSet) Clone() *BitSet {
	c := New(b.length)
	if b.set != nil { // Clone should not modify current object
		copy(c.set, b.set)
	}
	return c
}

// Convenience function: return two bitsets ordered by
// increasing length. Note: neither can be nil
func sortByLength(a *BitSet, b *BitSet) (ap *BitSet, bp *BitSet) {
	if a.length <= b.length {
		ap, bp = a, b
	} else {
		ap, bp = b, a
	}
	return
}

// Error is used to distinguish errors (panics) generated in this package.
type Error string

// NextSet returns the next bit set from the specified index,
// including possibly the current index
// along with an error code (true = valid, false = no set bit found)
// for i,e := v.NextSet(0); e; i,e = v.NextSet(i + 1) {...}
//
// Users concerned with performance may want to use NextSetMany to
// retrieve several values at once.
func (b *BitSet) NextSet(i uint) (uint, bool) {
	x := int(i >> log2WordSize)
	if x >= len(b.set) {
		return 0, false
	}
	w := b.set[x]
	w = w >> wordsIndex(i)
	if w != 0 {
		return i + trailingZeroes64(w), true
	}
	x++
	// bounds check elimination in the loop
	if x < 0 {
		return 0, false
	}
	for x < len(b.set) {
		if b.set[x] != 0 {
			return uint(x)*wordSize + trailingZeroes64(b.set[x]), true
		}
		x++

	}
	return 0, false
}

func trailingZeroes64(v uint64) uint {
	return uint(bits.TrailingZeros64(v))
}

// DumpAsBits dumps a bit set as a string of bits. Following the usual convention in Go,
// the least significant bits are printed last (index 0 is at the end of the string).
func (b *BitSet) DumpAsBits() string {
	if b.set == nil {
		return "."
	}
	buffer := bytes.NewBufferString("")
	i := len(b.set) - 1
	for ; i >= 0; i-- {
		fmt.Fprintf(buffer, "%064b.", b.set[i])
	}
	return buffer.String()
}
func (b *BitSet) NextSetMany(i uint, buffer []uint) (uint, []uint) {
	myanswer := buffer
	capacity := cap(buffer)
	x := int(i >> log2WordSize)
	if x >= len(b.set) || capacity == 0 {
		return 0, myanswer[:0]
	}
	skip := wordsIndex(i)
	word := b.set[x] >> skip
	myanswer = myanswer[:capacity]
	size := int(0)
	for word != 0 {
		r := trailingZeroes64(word)
		t := word & ((^word) + 1)
		myanswer[size] = r + i
		size++
		if size == capacity {
			goto End
		}
		word = word ^ t
	}
	x++
	for idx, word := range b.set[x:] {
		for word != 0 {
			r := trailingZeroes64(word)
			t := word & ((^word) + 1)
			myanswer[size] = r + (uint(x+idx) << 6)
			size++
			if size == capacity {
				goto End
			}
			word = word ^ t
		}
	}
End:
	if size > 0 {
		return myanswer[size-1], myanswer[:size]
	}
	return 0, myanswer[:0]
}
