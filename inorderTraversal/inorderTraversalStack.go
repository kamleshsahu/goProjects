package main

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {

	root := &Node{Value: 5}

	root.Left = &Node{Value: 3}
	root.Left.Left = &Node{Value: 2}
	root.Left.Right = &Node{Value: 4}
	root.Right = &Node{Value: 6}

	stack := []*Node{}

	for true {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			if len(stack) == 0 {
				break
			}
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			print(root.Value, " ->")
			root = root.Right
		}
	}
}
