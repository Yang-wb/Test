package main

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 0},
		{30000, 40000, 50000},
	}

	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d,%d); got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	a1 := 30000
	a2 := 40000
	a3 := 50000

	for i := 0; i < b.N; i++ {
		if actual := calcTriangle(a1, a2); actual != a3 {
			b.Errorf("calcTriangle(%d,%d); got %d; expected %d", a1, a2, actual, a3)
		}
	}
}
