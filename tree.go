package main

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

// func main() {
// 	var root *node = nil
// 	inputArray := []int{5, 9, 2, 3, 7, 6, 10, 1}
// 	for i := 0; i < len(inputArray); i++ {
// 		root = InsertIntoBTree(root, inputArray[i])
// 	}
// 	Preorder(root)
// 	fmt.Print("\n")

// 	Postorder(root)
// 	fmt.Print("\n")

// 	Inorder(root)
// }

func InsertIntoBTree(root *node, inputInt int) *node {
	if root != nil {
		if root.data >= inputInt { // In Go we can directly use pointer to the struct
			root.left = InsertIntoBTree(root.left, inputInt)
		} else {
			root.right = InsertIntoBTree(root.right, inputInt)
		}
	} else {
		root = &node{
			data:  inputInt,
			left:  nil,
			right: nil,
		}
	}
	return root
}

func Preorder(root *node) {
	if root == nil {
		return
	}
	fmt.Printf("%v --> ", root.data)
	Preorder(root.left)
	Preorder(root.right)
}

func Postorder(root *node) {
	if root == nil {
		return
	}
	Postorder(root.left)
	Postorder(root.right)
	fmt.Printf("%v --> ", root.data)
}

func Inorder(root *node) {
	if root == nil {
		return
	}
	Inorder(root.left)
	fmt.Printf("%v --> ", root.data)
	Inorder(root.right)
}
