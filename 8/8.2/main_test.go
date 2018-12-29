package main

import "testing"

func TestCountingMetadataForANodeWithoutChildren(t *testing.T) {
	n := node{
		header{0, 3},
		[]node{},
		[]int{10, 20, 31},
	}

	if m := n.countMetadata(); m != 61 {
		t.Errorf("Expecting metadata to be %d, got %d instead\n", 61, m)
	}
}

func TestCountingMetadataForANodeWithChildrenThatAreNotReferencedByItsMetadataEntries(t *testing.T) {
	firstChild := node{
		header{0, 2},
		[]node{},
		[]int{10, 20},
	}
	secondChild := node{
		header{0, 1},
		[]node{},
		[]int{30},
	}

	root := node{
		header{2, 1},
		[]node{firstChild, secondChild},
		[]int{100},
	}

	expectedMetadata := 0

	if m := root.countMetadata(); m != expectedMetadata {
		t.Errorf("Expecting metadata to be %d, got %d instead\n", expectedMetadata, m)
	}
}

func TestCountingMetadataForANodeWithChildrenThatAreReferencedByItsMetadataEntries(t *testing.T) {
	son := node{
		header{0, 2},
		[]node{},
		[]int{10, 20},
	}

	root := node{
		header{1, 1},
		[]node{son},
		[]int{1},
	}

	expectedMetadata := 10 + 20

	if m := root.countMetadata(); m != expectedMetadata {
		t.Errorf("Expecting metadata to be %d, got %d instead\n", expectedMetadata, m)
	}
}

func TestCountingMetadataForGivenExample(t *testing.T) {
	input := newStack([]int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2})

	root := newNode(input)

	expectedMetadata := 66

	if m := root.countMetadata(); m != expectedMetadata {
		t.Errorf("Expecting metadata to be %d, got %d instead\n", expectedMetadata, m)
	}
}
