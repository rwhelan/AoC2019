package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

func readMapFile() []byte {
	d, err := ioutil.ReadFile("Map.data")
	panicOnError(err)

	return bytes.TrimRight(d, "\n")
}

func parseRow(row []byte) (string, string) {
	result := bytes.Split(row, []byte{')'})
	return string(result[0]), string(result[1])
}

func main() {
	nodes := make(map[string]*node)

	getNode := func(key string) *node {
		_, ok := nodes[key]
		if !ok {
			nodes[key] = newNode(key)
		}

		return nodes[key]
	}

	for _, row := range bytes.Split(readMapFile(), []byte{'\n'}) {
		left, right := parseRow(row)
		getNode(left).addMember(getNode(right))
	}

	populate(nodes["COM"], 0)

	for i, n := range reverseWalkToNearestCommonNode(nodes["YOU"], nodes["SAN"]) {
		fmt.Println(i, n)
	}

	fmt.Println()
	for i, n := range reverseWalkToNearestCommonNode(nodes["SAN"], nodes["YOU"]) {
		fmt.Println(i, n)
	}

}

func reverseWalkToNearestCommonNode(src, dst *node) []*node {
	result := make([]*node, 0)
	c := src.Parent

	for {
		result = append(result, c)
		if hasSubInChildrenPaths(c, dst, false) {
			return result
		}

		c = c.Parent
	}
}

func hasSubInChildrenPaths(from, to *node, found bool) bool {
	for _, sub := range from.Members {
		if sub.Name == to.Name {
			return true
		}

		if hasSubInChildrenPaths(sub, to, false) {
			return true
		}
	}

	return false
}

func populate(n *node, depth int) {
	n.Depth = depth
	for _, cn := range n.Members {
		cn.Parent = n
		populate(cn, depth+1)
	}
}
