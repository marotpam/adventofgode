package main

import "fmt"

type tree struct {
	root node
}

func newTree(ints []int) *tree {
	return &tree{*newNode(ints)}
}

func (t *tree) countMetadata() int {
	return t.root.countMetadata()
}

type header struct {
	childNodes, metadataEntries int
}

type node struct {
	header   header
	children []*node
	metadata []int
}

func newNode(ints []int) *node {
	fmt.Printf("size: %d, %+v\n", len(ints), ints)
	if len(ints) < 3 {
		return &node{}
	}

	rootHeader := header{
		ints[0],
		ints[1],
	}

	nodes := []*node{}
	for i := 0; i < rootHeader.childNodes; i++ {
		header := header{
			ints[i],
			ints[i+1],
		}

		n := &node{
			header,
			nodes,
			ints[len(ints)-rootHeader.metadataEntries:],
		}

		nodes = append(n)
	}

	return &node{
		rootHeader,
		nodes,
		ints[len(ints)-rootHeader.metadataEntries:],
	}
}

func (n *node) countMetadata() int {
	if n == nil {
		return 0
	}

	r := 0

	for _, m := range n.metadata {
		r += m
	}

	for _, c := range n.children {
		r += c.countMetadata()
	}

	return r
}
