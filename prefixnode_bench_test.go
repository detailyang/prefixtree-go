package prefixtree

import (
	"testing"
)

func BenchmarkPathPrefix(b *testing.B) {
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

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v := pn.Lookup("/3/023")
		if v != 30 {
			b.Fatal("failed to route")
		}
	}
}
