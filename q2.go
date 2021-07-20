package main

import "fmt"

type Node struct {
	key string
	left *Node
	right *Node
}

func PreOrder(root *Node){

	fmt.Println(root.key)

	if root.left != nil {
		PreOrder(root.left)
	}

	if root.right != nil {
		PreOrder(root.right)
	}
}

func PostOrder(root *Node){

	if root.left != nil {
		PostOrder(root.left)
	}

	if root.right != nil {
		PostOrder(root.right)
	}
	fmt.Println(root.key)
}

func main(){
	Tree := [5]Node{}
	Tree[0].key = "+"
	Tree[1].key = "a"
	Tree[2].key = "-"
	Tree[3].key = "b"
	Tree[4].key = "c"

	Tree[0].left = &Tree[1]
	Tree[0].right = &Tree[2]
	Tree[2].left = &Tree[3]
	Tree[2].right = &Tree[4]

	PreOrder(&Tree[0])
	PostOrder(&Tree[0])
}