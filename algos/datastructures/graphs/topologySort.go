package main

import (
	"fmt"

	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

func main() {
	dag := simple.NewDirectedGraph()
	fmt.Println("hello")

	pages := []int{97, 13, 75, 29, 47}

	for i, page := range pages {
		foo := string(i)
		fmt.Printf("i: %v, page: %v, string: %v\n", i, page, foo)
	}

	node0 := dag.NewNode()
	dag.AddNode(node0)
	node1 := dag.NewNode()
	dag.AddNode(node1)
	node2 := dag.NewNode()
	dag.AddNode(node2)
	node3 := dag.NewNode()
	dag.AddNode(node3)
	node4 := dag.NewNode()
	dag.AddNode(node4)

	fmt.Printf("node1: %v", node1.ID())
	fmt.Printf("node0: %v", node0.ID())

	// Add edges
	// 97 | 13
	dag.SetEdge(dag.NewEdge(node0, node1))
	// 97 | 47
	dag.SetEdge(dag.NewEdge(node0, node4))
	// 75 | 29
	dag.SetEdge(dag.NewEdge(node2, node3))
	// 29 | 13
	dag.SetEdge(dag.NewEdge(node3, node1))
	// 97 | 29
	dag.SetEdge(dag.NewEdge(node0, node3))
	// 47 | 13
	dag.SetEdge(dag.NewEdge(node4, node1))
	// 75 | 47
	dag.SetEdge(dag.NewEdge(node2, node4))
	// 97 | 75
	dag.SetEdge(dag.NewEdge(node0, node2))
	// 47 | 29
	dag.SetEdge(dag.NewEdge(node4, node3))
	// 75 | 13
	dag.SetEdge(dag.NewEdge(node2, node1))

	fmt.Println("Sorted")
	// topo sort
	sorted, err := topo.Sort(dag)

	if err != nil {
		fmt.Printf("Err: %v", err)
	}

	for _, node := range sorted {
		fmt.Println(pages[node.ID()])
	}
}
