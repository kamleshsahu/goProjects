package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//Vals := []int{1, 2, 3}

	s := []int{2, 2, 5, 3, 6}

	node := toTree(&s)
	fmt.Println(node)
	data := make([]int, 0)
	node.toArray(&data)
	fmt.Println(data)
	node2 := toTree(&data)
	fmt.Println(node2)
	data2 := make([]int, 0)
	node2.toArray(&data2)
	fmt.Println(data2)
	//arr := make([]*int, 0)
	//err := json.Unmarshal([]byte(s), &arr)
	//if err != nil {
	//	return
	//}
	//
	//k, _ := json.Marshal(arr)
	//fmt.Println(k)
	//fmt.Println(arr)
	//v1 := make([]int, 2)
	//err = json.Unmarshal(k, &v1)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(v1)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	arr := make([]int, 0)
	root.toArray(&arr)
	str, _ := json.Marshal(arr)
	return string(str)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	arr := make([]int, 0)
	err := json.Unmarshal([]byte(data), &arr)
	if err != nil {
		return nil
	}
	return toTree(&arr)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) toArray(arr *[]int) {
	if node == nil {
		//*arr = append(*arr, -1)
		return
	}
	*arr = append(*arr, node.Val)
	node.Left.toArray(arr)
	node.Right.toArray(arr)
}

func toTree(arrp *[]int) *TreeNode {
	arr := *arrp
	if len(arr) == 0 || (arr)[0] == -1 {
		return nil
	}
	i := 0
	for i = 1; i < len(arr); i++ {
		if (arr)[i] > (arr)[0] {
			break
		}
	}
	node := &TreeNode{Val: arr[0]}
	left := arr[1:i]
	node.Left = toTree(&left)
	right := arr[i:]
	node.Right = toTree(&right)
	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
