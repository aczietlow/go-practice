package main

import (
	"fmt"
)

type Node struct {
	value User // Should user be a ref or value
	left  *Node
	right *Node
}

func main() {
	users := getUsers(18)
	fmt.Println(users)
	tree := Node{}
	for _, u := range users {
		tree.insert(u)
	}
	printTree(tree)

	lastUser := tree.getMax()
	fmt.Printf("Max user: %v\n", lastUser)

	minUser := tree.getMin()
	fmt.Printf("Min user: %v\n", minUser)
}

func (n *Node) insert(u User) {
	if n.value == (User{}) {
		n.value = u
		return
	}

	// BST values must be unique
	if u.equal(n.value) {
		return
	}

	if u.lt(n.value) {
		if n.left == nil {
			n.left = &Node{
				value: u,
			}
			return
		} else {
			n.left.insert(u)
			return
		}
	}
	// assume gt at this point
	if n.right == nil {
		n.right = &Node{
			value: u,
		}
		return
	} else {
		n.right.insert(u)
		return
	}
}

func (n *Node) getMax() User {
	current := n
	for current.right != nil {
		current = current.right
	}
	return current.value
}

func (n *Node) getMin() User {
	current := n
	for current.left != nil {
		current = current.left
	}
	return current.value
}
