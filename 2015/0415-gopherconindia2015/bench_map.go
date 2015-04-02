package main

import "testing"

// START OMIT
func BenchmarkMapAccess(b *testing.B) {
	m := map[int]int{0: 0, 1: 1} // HL
	for i := 0; i < b.N; i++ {
		_ = m[0] + m[1]
	}
}

// END OMIT
