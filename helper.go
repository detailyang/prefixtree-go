package prefixtree

import (
	"fmt"
	"io"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func longestCommonPrefix(a, b string) int {
	i := 0
	max := min(len(a), len(b))
	for i < max && a[i] == b[i] {
		i++
	}
	return i
}

func PPrint(w io.Writer, n *Node, prefix string) {
	fmt.Fprintf(w, " %03d %s%s[%d] %s %v \r\n", n.priority, prefix,
		n.path, len(n.children), n.typ, n.value)
	for l := len(n.path); l > 0; l-- {
		prefix += " "
	}
	for _, child := range n.children {
		PPrint(w, child, prefix)
	}
}

func Cleanup(p, n *Node) {
	for _, child := range n.children {
		Cleanup(n, child)
	}

	if n.value == nil && len(n.children) == 0 && n.typ == LeafNodeType { // orphan node
		if p != nil {
			for i := range p.children {
				if p.children[i] == n {
					p.children = append(p.children[:i], p.children[i+1:]...)
					break
				}
			}
		}
	}
}
