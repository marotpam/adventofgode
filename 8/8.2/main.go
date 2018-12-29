package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

type header struct {
	childNodes, metadataEntries int
}

type node struct {
	header   header
	children []node
	metadata []int
}

type stack struct {
	values []int
}

func newStack(ints []int) *stack {
	return &stack{ints}
}

func (s *stack) pop() int {
	last := s.values[0]
	s.values = s.values[1:]

	return last
}

func newNode(s *stack) node {
	rootHeader := header{
		s.pop(),
		s.pop(),
	}
	n := &node{
		header: rootHeader,
	}

	nodes := []node{}
	for i := 0; i < rootHeader.childNodes; i++ {
		nodes = append(nodes, newNode(s))
	}
	n.children = nodes

	metadata := []int{}
	for i := 0; i < rootHeader.metadataEntries; i++ {
		metadata = append(metadata, s.pop())
	}
	n.metadata = metadata

	return *n
}

func (n *node) countMetadata() int {
	r := 0

	childrenCount := len(n.children)
	if childrenCount == 0 {
		for _, m := range n.metadata {
			r += m
		}

		return r
	}

	for _, m := range n.metadata {
		if m <= childrenCount {
			r += n.children[m-1].countMetadata()
		}
	}

	return r
}

func main() {
	bs, _ := ioutil.ReadFile("input.txt")

	var ints []int
	for _, p := range bytes.Fields(bs) {
		n, _ := strconv.Atoi(string(p))

		ints = append(ints, n)
	}

	n := newNode(newStack(ints))

	fmt.Println(n.countMetadata()) // 22198
}
