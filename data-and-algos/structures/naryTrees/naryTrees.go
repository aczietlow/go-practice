package main

import "fmt"

type Node struct {
	value    string
	children []*Node
}

func main() {
	t := Node{
		value: "Richard",
		children: []*Node{
			{
				value: "Rick",
				children: []*Node{
					{
						value: "Richard",
					},
				},
			},
			{
				value: "Marty",
				children: []*Node{
					{
						value: "Chirs",
						children: []*Node{
							{
								value: "Levi",
							},
						},
					},
					{
						value: "Olivia",
					},
				},
			},
			{
				value: "Judy",
			},
			{
				value: "Dana",
				children: []*Node{
					{
						value: "Kim",
					},
				},
			},
		},
	}

	t.show()
}

func (n *Node) show() {
	if len(n.children) > 0 {
		for _, child := range n.children {
			fmt.Printf("name  %v\n", child.value)
		}
	}
	// fmt.Printf("|─\n|\n|──")
	fmt.Printf("name  %v\n", n.value)
}

func printTree(n *Node) string {
	var output string
	if len(n.children) > 0 {
		for _, child := range n.children {
			output += fmt.Sprintf("──%v\n", child.value)
		}
	}
	return output
}
