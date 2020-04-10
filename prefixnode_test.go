package prefixtree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSafePathPrefixInsertDuplicated(t *testing.T) {
	pn := NewSafePrefixTree()
	for _, i := range []string{
		"/helloworld.Greeter",
		"/com.a.b.c.d.helloworld.Greeter",
		"/com.1.2.3.4.tripleservice.Greeter",
		"/helloworld.Greeter",
	} {
		pn.Add(i, i)
	}
	count := 0
	pn.Iterate(func(n *Node) {
		if n.GetType() == RootNodeType {
			count++
		}
	})
	require.Equal(t, 3, count)
}

func TestSafePathPrefixInsertLeaf(t *testing.T) {
	pn := NewSafePrefixTree()
	pn.Add("/helloworld.Greeter", 1)
	pn.Add("/astore.anthelloworld.AntGreeter", 2)
	pn.Add("/com.taobao.hsf.triple.helloworld.Greeter", 3)
	pn.Add("/", 4)

	pn.Iterate(func(n *Node) {
		if n.GetPath() == "/" {
			require.Equal(t, RootNodeType, n.GetType())
		}
	})
}

func TestSafePathPrefixInsertDuplicatedAdd(t *testing.T) {
	pn := NewSafePrefixTree()
	pn.Add("/com.a.b.c.helloworld.Greeter", 1)
	pn.Add("/com.fliggy.fcecore.tripleservice.Greeter", 2)
	pn.Add("/helloworld.Gretter", 3)

	pn.Iterate(func(n *Node) {
		if n.GetPath() == "com." {
			require.Equal(t, LeafNodeType, n.GetType())
		}
	})
}

func TestPathPrefixInsertDuplicatedAdd(t *testing.T) {
	pn := NewNode()
	pn.Add("/1", 1)
	pn.Add("/1", 2)

	count := 0
	pn.Iterate(func(n *Node) {
		count++
		require.Equal(t, 2, n.value)
	})
	require.Equal(t, 1, count)
}

func TestPathPrefixInsertDuplicatedInsert(t *testing.T) {
	pn := NewNode()
	pn.insert("/1", 1)
	pn.insert("/1", 2)

	count := 0
	pn.Iterate(func(n *Node) {
		count++
		require.Equal(t, 2, n.value)
	})
	require.Equal(t, 1, count)
}

func TestPathPrefix(t *testing.T) {
	pn := NewNode()
	pn.Add("/", 0)
	pn.Add("/1", 1)
	pn.Add("/2", 2)
	pn.Add("/3", 3)
	pn.Add("/1/0", 10)
	pn.Add("/1/1", 11)
	pn.Add("/2/0", 20)
	pn.Add("/2/1", 21)
	pn.Add("/3/0", 30)
	pn.Add("/3/1", 31)

	fmt.Println(pn.String())

	require.Equal(t, 1, pn.Lookup("/1"))
	require.Equal(t, 2, pn.Lookup("/2"))
	require.Equal(t, 3, pn.Lookup("/3"))

	require.Equal(t, 11, pn.Lookup("/1/1"))
	require.Equal(t, 11, pn.Lookup("/1/123"))
	require.Equal(t, 10, pn.Lookup("/1/0"))
	require.Equal(t, 10, pn.Lookup("/1/023"))

	require.Equal(t, 21, pn.Lookup("/2/1"))
	require.Equal(t, 21, pn.Lookup("/2/123"))
	require.Equal(t, 20, pn.Lookup("/2/0"))
	require.Equal(t, 20, pn.Lookup("/2/023"))

	require.Equal(t, 31, pn.Lookup("/3/1"))
	require.Equal(t, 31, pn.Lookup("/3/123"))
	require.Equal(t, 30, pn.Lookup("/3/0"))
	require.Equal(t, 30, pn.Lookup("/3/023"))

	require.Equal(t, 0, pn.Lookup("/a"))
	require.Equal(t, 0, pn.Lookup("/abcd"))
	require.Equal(t, 0, pn.Lookup("/abcdas/asdfasdf"))

	pn.Remove("/1")
	pn.Remove("/2")
	pn.Remove("/3")
	fmt.Println(pn.String())
	require.Equal(t, 0, pn.Lookup("/1"))
	require.Equal(t, 0, pn.Lookup("/2"))
	require.Equal(t, 0, pn.Lookup("/3"))

	pn.Remove("/1/0")
	pn.Remove("/1/1")
	pn.Cleanup()
	fmt.Println(pn.String())

	pn.Iterate(func(n *Node) {
		fmt.Println("-", n.GetPath(), n.GetValue())
	})
}
