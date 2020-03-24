// Package prefix tree implements the radix tree.
package prefixtree

import (
	"bytes"
)

// NodeType represents the type of node.
type NodeType uint8

const (
	LeafNodeType NodeType = iota // default
	RootNodeType
)

// String returns the node type representation.
func (n NodeType) String() string {
	switch n {
	case LeafNodeType:
		return "leaf"
	case RootNodeType:
		return "root"
	default:
		return "unknown node type"
	}
}

// NodeValue represents the value with original path.
type NodeValue struct {
	path  string
	value interface{}
}

// NewNodeValue creates a new NodeValue.
func NewNodeValue(path string, value interface{}) *NodeValue {
	return &NodeValue{
		path:  path,
		value: value,
	}
}

// Node represents a leaf node.
type Node struct {
	typ      NodeType
	path     string
	indices  []byte
	value    interface{}
	priority uint32
	children []*Node
}

// NewNode creates a new leaf node.
func NewNode() *Node {
	return &Node{}
}

func (n *Node) Iterate(fn func(n *Node)) {
	Iterate(n, fn)
}

func (n *Node) GetValue() interface{} { return n.value }
func (n *Node) GetPath() string       { return n.path }

// Remove removes the given path from the tree and return the value.
func (n *Node) Remove(path string) (value interface{}) {
walk:
	for {
		prefix := n.path
		if len(path) > len(prefix) {
			path = path[len(prefix):]
			idxc := path[0]

			if i := bytes.IndexByte(n.indices, idxc); i != -1 {
				n = n.children[i]
				continue walk
			}

		} else if path == prefix {
			value = n.value
			n.value = nil
			n.typ = LeafNodeType
		}

		return nil
	}
}

// Cleanup cleans up the orphan nodes.
func (n *Node) Cleanup() {
	Cleanup(nil, n)
}

// Lookup looks up the longest prefix path.
func (n *Node) Lookup(path string) (value interface{}) {
walk:
	for {
		prefix := n.path
		if len(path) > len(prefix) {
			if path[:len(prefix)] == prefix {
				path = path[len(prefix):]
				idxc := path[0]

				if n.typ == LeafNodeType {
					if i := bytes.IndexByte(n.indices, idxc); i != -1 {
						n = n.children[i]
						continue walk
					}

					value = nil
					return
				}

				// try longest match else use current node value
				if i := bytes.IndexByte(n.indices, idxc); i != -1 {
					value = n.children[i].Lookup(path)
					if value != nil {
						return value
					}
				}

				value = n.value
				return
			}

		} else if path == prefix {
			// We should have reached the node containing the handle.
			// Check if this node has a handle registered.
			if value = n.value; value != nil {
				return
			}

			// No handle found. Check if a handle for this path + a
			// trailing slash exists for trailing slash recommendation
			return
		}

		return
	}
}

// Add adds the path to the tree with value.
//
// Value cannot be nil.
func (n *Node) Add(path string, value interface{}) {
	if value == nil {
		panic("value cannot be nil")
	}

	n.priority++

	if n.path == "" && len(n.indices) == 0 {
		n.insert(path, value)
		return
	}

walk:
	for {
		i := longestCommonPrefix(path, n.path)

		// split node
		if i < len(n.path) {
			child := &Node{
				path:     n.path[i:],
				typ:      RootNodeType,
				indices:  n.indices,
				children: n.children,
				value:    n.value,
				priority: n.priority - 1,
			}
			n.typ = LeafNodeType
			n.children = []*Node{child}
			n.indices = []byte{n.path[i]}
			n.path = n.path[:i]
			n.value = nil
		}

		if i < len(path) {
			path = path[i:]
			idxc := path[0]

			// Check if a child with the next path byte exists
			if i := bytes.IndexByte(n.indices, idxc); i >= 0 {
				i = n.incrementChildPrio(i)
				n = n.children[i]
				continue walk
			}

			// []byte for proper unicode char conversion, see #65
			n.indices = append(n.indices, idxc)
			child := &Node{}
			n.children = append(n.children, child)
			n.incrementChildPrio(len(n.indices) - 1)
			n = child
			n.insert(path, value)
			return
		}

		n.value = value
		return
	}
}

// Increments priority of the given child and reorders if necessary
func (n *Node) incrementChildPrio(pos int) int {
	cs := n.children
	cs[pos].priority++
	prio := cs[pos].priority

	// Adjust position (move to front)
	newPos := pos
	for ; newPos > 0 && cs[newPos-1].priority < prio; newPos-- {
		// Swap node positions
		cs[newPos-1], cs[newPos] = cs[newPos], cs[newPos-1]
	}

	// Build new index char string
	if newPos != pos {
		n.indices = append(n.indices[:newPos],
			n.indices[pos:pos+1]...)
		n.indices = append(n.indices, n.indices[newPos:pos]...)
		n.indices = append(n.indices, n.indices[pos+1:]...)
	}

	return newPos
}

func (n *Node) insert(path string, value interface{}) {
	n.path = path
	n.value = value
	n.path = path
	n.typ = RootNodeType
}

// String represents the tree view.
func (n *Node) String() string {
	var w bytes.Buffer
	PPrint(&w, n, "")
	return w.String()
}
