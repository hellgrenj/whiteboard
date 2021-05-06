package bst

import "fmt"

type Node struct {
	Value int
	left  *Node
	right *Node
}
type Tree struct {
	root *Node
}

func (t *Tree) Insert(number int) {
	node := &Node{Value: number}
	if t.root == nil {
		t.root = node
	} else {
		t.root.insert(node)
	}
}
func (n *Node) insert(node *Node) {
	if node.Value > n.Value {
		if n.right == nil {
			n.right = node
		} else {
			n.right.insert(node)
		}
	} else if node.Value < n.Value {
		if n.left == nil {
			n.left = node
		} else {
			n.left.insert(node)
		}
	}
}
func (t *Tree) Invert() {
	if t.root != nil {
		invert(t.root)
	} else {
		fmt.Println("empty tree can not be inverted")
	}
}
func invert(n *Node) {
	if n == nil {
		return
	} else {
		invert(n.left)
		invert(n.right)
		n.left, n.right = n.right, n.left
	}
}
func (t *Tree) Print() {
	printTree(t.root, t.numberOfNodes(), true)
}
func printTree(n *Node, leftPadding int, isRoot bool) {
	if n == nil {
		return
	} else {
		for l := 0; l < leftPadding; l++ {
			fmt.Printf("  ")
		}
		fmt.Printf("%d\n", n.Value)
		if isRoot {
			printTree(n.left, leftPadding-3, false)
			printTree(n.right, leftPadding+3, false)
		} else {
			printTree(n.left, leftPadding-1, false)
			printTree(n.right, leftPadding+1, false)
		}
	}
}
func (t *Tree) numberOfNodes() int {
	return count(t.root)
}
func count(node *Node) int {
	c := 1
	if node == nil {
		return 0
	}
	c += count(node.left)
	c += count(node.right)
	return c
}
func (t *Tree) Search(number int) *Node {
	return search(t.root, number)
}
func search(node *Node, number int) *Node {
	if node == nil || node.Value == number {
		return node
	}
	if number > node.Value {
		return search(node.right, number)
	}
	return search(node.left, number)
}
