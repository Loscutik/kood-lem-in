package main

import (
	"testing"
)

func BenchmarkNogirutin(b *testing.B) {
	nogorutin()
}

func BenchmarkGorutin(b *testing.B) {
	gorutin()
}
func BenchmarkGorutinPoint(b *testing.B) {
	gorutinPoint()
}
