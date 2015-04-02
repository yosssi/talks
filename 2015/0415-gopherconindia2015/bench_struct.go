package main

import "testing"

// START OMIT
func BenchmarkStructAccess(b *testing.B) {
	m := struct{ a, b int }{0, 1} // HL
	for i := 0; i < b.N; i++ {
		_ = m.a + m.b
	}
}

// END OMIT
