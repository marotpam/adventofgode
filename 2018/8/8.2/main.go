package _2018

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

type number struct {
	value int
	next *number
}

type numbers struct {
	first, last *number
}

func newNumbers(ints []int) *numbers {
	ns := &numbers{}
	if len(ints) > 0 {
		n := &number{value: ints[0]}
		previous := n
		ns.first = n
		for i:= 1; i < len(ints); i++ {
			n = &number{value: ints[i]}
			previous.next = n
		}
		ns.last = n
	}

	return ns
}

func (s *numbers) add(i int) {
	n := &number{value: i}
	if s.first == nil {
		s.first = n
	}
	if s.last != nil {
		s.last.next = n
	}
	s.last = n
}

func (s *numbers) pop() int {
	first := s.first.value
	s.first = s.first.next

	return first
}

func newNode(s *numbers) node {
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

	numbers := newNumbers([]int{})
	for _, p := range bytes.Fields(bs) {
		n, _ := strconv.Atoi(string(p))

		numbers.add(n)
	}

	n := newNode(numbers)

	fmt.Println(n.countMetadata()) // 22198
}
