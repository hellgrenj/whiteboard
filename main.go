package main

import (
	"fmt"

	"github.com/hellgrenj/whiteboard/bst"
	"github.com/hellgrenj/whiteboard/fibonacci"
)

func main() {
	// Fibonacci
	fmt.Printf("\nFibonacci\n")
	fibonacci.SlowFib(10)
	fibonacci.FasterFib(10)

	// BST
	fmt.Printf("\nBinary Search Tree\n")
	tree := &bst.Tree{}
	tree.Insert(8)
	tree.Insert(7)
	tree.Insert(13)
	tree.Insert(17)
	tree.Insert(12)
	tree.Insert(4)
	tree.Print()
	nonExisting := tree.Search(2000)
	if nonExisting != nil {
		panic("node with value 2000 does not exist so what gives?")
	}
	node12 := tree.Search(12)
	fmt.Printf("\nSearched for and found node with value %d\n", node12.Value)
	fmt.Println("... and now lets invert the BST")
	tree.Invert()
	tree.Print()

}
