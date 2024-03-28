package main

import (
	"common"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

const (
	minSpace    = 100000
	maxSpace    = 70000000
	updateSpace = 30000000
	ls          = "$ ls"
	cd          = "$ cd"
	up          = ".."
	dir         = "dir"
	seperator   = "/"
)

type node struct {
	name   string
	parent *node
	size   int
	nodes  map[string]*node
}

type tree struct {
	root *node
}

func (tree *tree) insert(dest *node, node *node) bool {
	if tree.root == nil {
		tree.root = node
		return true
	}

	if _, ok := dest.nodes[node.name]; !ok {
		dest.nodes[node.name] = node
		return true
	}

	return false
}

/**
 * part_1
 */
func lowestSizes(tree *tree) (result int) {
	m := map[string]int{}
	stack := common.NewStack[*node]()
	stack.Push(tree.root)

	for !stack.IsEmpty() {
		node, ok := stack.Pop()

		if !ok {
			continue
		}

		if node.size < minSpace {
			m[node.name] = node.size
		}

		for _, n := range node.nodes {
			stack.Push(n)
		}
	}

	return common.MapReduce(m, func(size int, acc int) int {
		acc += size
		return acc
	}, result)
}

/**
 * part_2
 */
func lowestFreeingSize(tree *tree) (result int) {
	free := maxSpace - tree.root.size
	list := []*node{}
	stack := common.NewStack[*node]()
	stack.Push(tree.root)

	for !stack.IsEmpty() {
		node, ok := stack.Pop()

		if !ok {
			continue
		}

		if node.size+free >= updateSpace {
			list = append(list, node)
		}

		for _, n := range node.nodes {
			stack.Push(n)
		}
	}

	slices.SortFunc(list, func(a *node, b *node) int {
		return a.size - b.size
	})

	return list[0].size
}

/**
 * driver
 */
func parseInput(buffer []byte) (result *tree) {
	result = &tree{}
	parent := &node{}

	for _, line := range strings.Split(string(buffer), "\n") {
		switch {
		case strings.Contains(line, ls) || strings.Contains(line, dir):
			continue
		case strings.Contains(line, cd):
			name := ""

			if i, _ := fmt.Sscanf(line, cd+"%s", &name); i != 1 {
				continue
			}

			if name == up {
				parent = parent.parent
			} else {
				n := &node{
					name:   parent.name + name + seperator,
					parent: parent,
					nodes:  map[string]*node{},
				}
				result.insert(parent, n)
				parent = n
			}
		default:
			size := 0

			if i, _ := fmt.Sscanf(line, "%d", &size); i != 1 {
				continue
			}

			p := parent

			for p != nil {
				p.size += size
				p = p.parent
			}
		}
	}

	return result
}

func main() {
	buffer, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 3 || os.Args[1] != "part" || !strings.Contains("12", os.Args[2]) {
		log.Fatal("usage: part <1|2>")
	}

	if arg := os.Args[2]; arg == "1" {
		fmt.Println("result:", lowestSizes(parseInput(buffer)))
	} else {
		fmt.Println("result:", lowestFreeingSize(parseInput(buffer)))
	}
}
