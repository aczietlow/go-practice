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

	// Creates BST
	tree := Node{}
	for _, u := range users {
		tree.insert(u)
	}
	printTree(tree)

	lastUser := tree.getMax()
	fmt.Printf("Max user: %v\n", lastUser)

	minUser := tree.getMin()
	fmt.Printf("Min user: %v\n", minUser)

	fmt.Printf("Deleting user: %v\n", users[13])
	bst := tree.delete(users[13])
	// starting with a leaf node for simplification
	printTree(bst)
}

func (n *Node) delete(u User) Node {
	// If the node is empty return, there's nothing to do
	if n.value == (User{}) {
		// TODO: should we throw an error at this point? to let the user know they attempted to delete a node that doesnt exist
		return Node{}
	}

	// parse down the left tree
	if u.lt(n.value) {
		node := n.left.delete(u)
		n.left = &node
	} else if u.equal(n.value) { // delete the node?
		// TODO: this doesn't delete the node, instead inserts an empty node
		return Node{}
	}

	return *n
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
