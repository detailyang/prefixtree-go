package prefixtree

import "sync"

// SafePrefixTree implements the radix tree with lock.
type SafePrefixTree struct {
	sync.RWMutex
	n *Node
}

// NewSafePrefixTree creates a SafePrefixTree.
func NewSafePrefixTree() *SafePrefixTree {
	return &SafePrefixTree{
		n: NewNode(),
	}
}

func (n *SafePrefixTree) Iterate(fn func(n *Node)) {
	n.RLock()
	Iterate(n.n, fn)
	n.RUnlock()
}

// Remove removes the given path from the tree.
func (t *SafePrefixTree) Remove(path string) (value interface{}) {
	t.Lock()
	value = t.n.Remove(path)
	t.Unlock()
	return value
}

// Lookup looks up the given path from the tree.
func (t *SafePrefixTree) Lookup(path string) (value interface{}) {
	t.RLock()
	value = t.n.Lookup(path)
	t.RUnlock()
	return value
}

// Add adds the given path to the tree.
func (t *SafePrefixTree) Add(path string, value interface{}) {
	t.Lock()
	t.n.Add(path, value)
	t.Unlock()
}

// String returns the string representation of the SafePrefixTree.
func (t *SafePrefixTree) String() string {
	t.RLock()
	s := t.n.String()
	t.Unlock()
	return s
}

// PrefixTree implements the radix tree.
type PrefixTree struct {
	n *Node
}

// NewPrefixTree creates a new PrefixTree.
func NewPrefixTree() *PrefixTree {
	return &PrefixTree{
		n: NewNode(),
	}
}

func (n *PrefixTree) Iterate(fn func(n *Node)) {
	n.n.Iterate(fn)
}

// Remove removes the given path from the tree.
func (t *PrefixTree) Remove(path string) (value interface{}) {
	return t.n.Remove(path)
}

// Lookup looks up the given path in the tree.
func (t *PrefixTree) Lookup(path string) (value interface{}) {
	return t.n.Lookup(path)
}

// Add adds the given path to the tree.
func (t *PrefixTree) Add(path string, value interface{}) {
	t.n.Add(path, value)
}

// String returns the view of tree.
func (t *PrefixTree) String() string {
	return t.n.String()
}
