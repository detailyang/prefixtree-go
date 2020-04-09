package prefixtree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPathPrefixInsertDuplicated(t *testing.T) {
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
