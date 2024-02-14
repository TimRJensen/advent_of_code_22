package main

import (
	"common"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type node struct {
	name   string
	parent *node
	size   int
	nodes  map[string]*node
}

func (node *node) Compare(a *node, b *node) int {
	switch {
	case a.size < b.size:
		return -1
	case a.size > b.size:
		return 1
	default:
		return 0
	}
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
func getSize(tree *tree) (result int) {
	m := make(map[string]int)
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
func getLowestSize(tree *tree) (result int) {
	free := maxSpace - tree.root.size
	list := make([]*node, 0, 256)
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

	sort.Slice(list, func(i int, j int) bool {
		return list[i].size < list[j].size
	})

	return list[0].size
}

/**
 * driver
 */
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

func getInput(buffer []byte) (result *tree) {
	result = new(tree)
	parent := new(node)

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
					nodes:  make(map[string]*node)}
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

	if arg := os.Args[1]; arg == "part_1" {
		fmt.Println("result:", getSize(getInput(buffer)))
	} else {
		fmt.Println("result:", getLowestSize(getInput(buffer)))
	}
}
