package main

import (
	"fmt"
	"strings"
)

func printTree(bstNode Node) {
	lines := []string{}
	formatTreeString(&bstNode, &lines, 0)
	fmt.Println(strings.Join(lines, "\n"))
}

func formatTreeString(bstNode *Node, lines *[]string, level int) {
	if bstNode != nil {
		formatTreeString(bstNode.right, lines, level+1)
		line := fmt.Sprintf("%s> %v", strings.Repeat(" ", 4*level), bstNode.value)
		*lines = append(*lines, line)
		formatTreeString(bstNode.left, lines, level+1)
	}
}
