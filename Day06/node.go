package main

import "fmt"

type node struct {
	Parent  *node
	Name    string
	Members []*node
	Depth   int
}

func newNode(name string) *node {
	return &node{
		Name:    name,
		Parent:  nil,
		Members: make([]*node, 0),
		Depth:   0,
	}
}

func (n *node) addMember(m *node) {
	n.Members = append(n.Members, m)
}

func (n *node) String() string {
	var m string
	for _, i := range n.Members {
		m += i.Name
	}

	return fmt.Sprintf("Name: %s, Parent: %s, Members: %s, Depth: %d",
		n.Name, n.Parent.Name, m, n.Depth)
}
