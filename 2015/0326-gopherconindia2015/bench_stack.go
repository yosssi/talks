package main

import "testing"

// START OMIT
func BenchmarkStack(b *testing.B) {
	f := func() [128]int64 { return [128]int64{} }
	for i := 0; i < b.N; i++ {
		f()
	}
}

// END OMIT
