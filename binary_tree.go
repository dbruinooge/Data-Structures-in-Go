package main

import "fmt"

// Node
type Node struct {
	Key int
	Left *Node
	Right *Node
}

// Insert
func (parent *Node) Insert(k int) {
	child := &Node{ Key: k }

	switch {
	case child.Key < parent.Key:
		if parent.Left == nil {
			parent.Left = child
		} else {
			parent.Left.Insert(k)
		}
	case child.Key > parent.Key:
		if parent.Right == nil {
			parent.Right = child
		} else {
			parent.Right.Insert(k)
		}
	default:
		return
	}
}

// Search
func (n *Node) Search(k int) bool {
	switch {
	case k < n.Key && n.Left != nil:
		return n.Left.Search(k)
	case k > n.Key && n.Right != nil:
		return n.Right.Search(k)
	case k == n.Key:
		return true
	}
	return false
}

func main() {
	tree := &Node{ Key: 100 }
	tree.Insert(150)
	tree.Insert(125)
	tree.Insert(175)
	tree.Insert(200)
	tree.Insert(200)

	fmt.Println(tree) // 100
	fmt.Println(tree.Right) // 150
	fmt.Println(tree.Right.Left) // 125
	fmt.Println(tree.Right.Right) // 175
	fmt.Println(tree.Right.Right.Right) // 200

	search100 := tree.Search(100) // true
	search125 := tree.Search(125) // true
	search150 := tree.Search(150) // true
	search175 := tree.Search(175) // true
	search200 := tree.Search(200) // true
	fmt.Println(search100, search125, search150, search175, search200)
	search0 := tree.Search(0) // false
	search300 := tree.Search(300) // false
	searchNeg100 := tree.Search(-100) // false
	search110 := tree.Search(110) // false
	fmt.Println(search0, search300, searchNeg100, search110)




}