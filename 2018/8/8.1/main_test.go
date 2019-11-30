package main

import "testing"

func TestItCanCountMetadataEntriesOfTreeWithOnlyOneLeafAndOneMetadataEntry(t *testing.T) {
	n := node{
		header{0, 1},
		[]node{},
		[]int{10},
	}

	if m := n.countMetadata(); m != 10 {
		t.Errorf("Expecting metadata to be %d, got %d instead\n", 10, m)
	}
}

func TestItCanCountMetadataEntriesOfTreeWithOnlyOneLeafWithMultipleMetadataEntries(t *testing.T) {
	n := node{
		header{0, 3},
		[]node{},
		[]int{10, 20, 31},
	}

	if m := n.countMetadata(); m != 61 {
		t.Errorf("Expecting metadata to be %d, got %d instead\n", 61, m)
	}
}

func TestItCanCountMetadataEntriesOfTreeWithMultipleNodesWithMultipleMetadataEntries(t *testing.T) {
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

	expectedMetadata := 10 + 20 + 30 + 100

	if m := root.countMetadata(); m != expectedMetadata {
		t.Errorf("Expecting metadata to be %d, got %d instead\n", expectedMetadata, m)
	}
}

func TestItCanCountMetadataEntriesOfTreeWithMultipleLevelsOfNodes(t *testing.T) {
	grandSon := node{
		header{0, 1},
		[]node{},
		[]int{30},
	}
	son := node{
		header{0, 2},
		[]node{grandSon},
		[]int{10, 20},
	}

	root := node{
		header{1, 1},
		[]node{son},
		[]int{100},
	}

	expectedMetadata := 10 + 20 + 30 + 100

	if m := root.countMetadata(); m != expectedMetadata {
		t.Errorf("Expecting metadata to be %d, got %d instead\n", expectedMetadata, m)
	}
}

func TestCreatingATreeWithOnlyOneLeafWithOneMetadataEntry(t *testing.T) {
	input := newNumbers([]int{0, 1, 10})

	root := newNode(input)

	assertChildrenCount(t, root, 0)
	assertMetadataEntries(t, root, 1, 10)
}

func TestCreatingATreeWithOnlyOneLeafWithMultipleMetadataEntries(t *testing.T) {
	input := newNumbers([]int{0, 3, 10, 20, 30})

	root := newNode(input)

	assertChildrenCount(t, root, 0)
	assertMetadataEntries(t, root, 3, 60)
}

func TestCreatingATreeWithARootWithOneChildrenWithOneMetadataEntryEach(t *testing.T) {
	input := newNumbers([]int{1, 1, 0, 1, 20, 10})

	root := newNode(input)

	assertChildrenCount(t, root, 1)
	assertMetadataEntries(t, root, 1, 30)

	assertChildrenCount(t, root.children[0], 0)
	assertMetadataEntries(t, root.children[0], 1, 20)
}

func TestCreatingATreeWithARootWithOneChildrenWithMultipleMetadataEntries(t *testing.T) {
	input := newNumbers([]int{1, 2, 0, 3, 20, 1, 2, 10, 30})

	root := newNode(input)

	assertChildrenCount(t, root, 1)
	assertMetadataEntries(t, root, 2, 63)

	assertChildrenCount(t, root.children[0], 0)
	assertMetadataEntries(t, root.children[0], 3, 23)
}

func TestCreatingATreeWithMultipleLevelsOfSingleNodes(t *testing.T) {
	input := newNumbers([]int{1, 1, 1, 1,  0, 1, 15, 20, 10})

	root := newNode(input)

	assertChildrenCount(t, root, 1)
	assertMetadataEntries(t, root, 1, 45)

	son := root.children[0]
	assertChildrenCount(t, son, 1)
	assertMetadataEntries(t, son, 1, 35)

	grandson := root.children[0].children[0]
	assertChildrenCount(t, grandson, 0)
	assertMetadataEntries(t, grandson, 1, 15)
}

func TestCreatingATreeWithARootWithTwoChildren(t *testing.T) {
	input := newNumbers([]int{2, 1, 0, 1, 20, 0, 1, 15, 10})

	root := newNode(input)

	assertChildrenCount(t, root, 2)
	assertMetadataEntries(t, root, 1, 45)

	assertChildrenCount(t, root.children[0], 0)
	assertMetadataEntries(t, root.children[0], 1, 20)

	assertChildrenCount(t, root.children[1], 0)
	assertMetadataEntries(t, root.children[1], 1, 15)
}

func assertChildrenCount(t *testing.T, node node, expectedChildren int) {
	c := node.header.childNodes
	if c != expectedChildren {
		t.Errorf("node should have %d children in the header, %d found\n", expectedChildren, c)
	}

	if childrenCount := len(node.children); childrenCount != c {
		t.Errorf("children count (%d) does not match actual number of children nodes (%d)\n", c, childrenCount)
	}
}

func assertMetadataEntries(t *testing.T, node node, expectedMetadataEntries, expectedMetadataCount int) {
	me := node.header.metadataEntries
	if me != expectedMetadataEntries {
		t.Errorf("node should have %d metadata entries in the header, %d found\n", expectedMetadataEntries, me)
	}

	if countMetadataEntries := len(node.metadata); countMetadataEntries != me {
		t.Errorf("node should have %d metadata entries according to header, found %d: %+v", me, countMetadataEntries, node.metadata)
	}

	if m := node.countMetadata(); m != expectedMetadataCount {
		t.Errorf("node should have metadata of %d, found %d\n", expectedMetadataCount, m)
	}
}
