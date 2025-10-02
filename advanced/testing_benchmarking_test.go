package main

import (
	"math/rand"
	"testing"
)

// func add(a, b int) int {
// 	return a + b
// }

func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(100)
	}
	return slice
}

func sumSlice(slice []int) int {
	sum := 0
	for i := range slice {
		sum += slice[i]
	}
	return sum
}

// ==============Profiling
//Command to run: go test -bench=. -memprofile mem.pprof testing_benchmarking_test.go | grep -v 'cpu:'
func TestGenerateRandomSlice(t *testing.T) {
	size := 100
	slice := generateRandomSlice(size)
	if len(slice) != size {
		t.Errorf("expected slice size %d, but received %d", size, len(slice))
	}

}

func BenchmarkGeneraterandomSlice(b *testing.B) {
	for range b.N {
		generateRandomSlice(1000)
	}
}

func BenchmarkSumSlice(b *testing.B) {
	slice := generateRandomSlice(1000)
	b.ResetTimer()
	for range b.N {
		sumSlice(slice)
	}
}

//====================================Benchmark
// func BenchmarkAddSmallInput(b *testing.B) {
// 	for range b.N {
// 		add(2, 3)
// 	}
// } //Command to run: go test -bench=. -benchmem testing_benchmarking_test.go | grep -v 'cpu:'

// func BenchmarkAddMediumInput(b *testing.B) {
// 	for range b.N {
// 		add(200, 300)
// 	}
// }

// func BenchmarkAddLargeInput(b *testing.B) {
// 	for range b.N {
// 		add(2000, 3000)
// 	}
// }

//================================Testing

// func TestAdd(t *testing.T) {
// 	res := add(2, 3)
// 	exp := 5
// 	if res != exp {
// 		t.Errorf("Add(2,3) is equal to %d, want %d", res, exp)
// 	}
// }

// //Table Drivet Testing

// func TestAddTableDriven(t *testing.T) {
// 	tests := []struct {
// 		a        int
// 		b        int
// 		expected int
// 	}{
// 		{2, 3, 5},
// 		{0, 0, 0},
// 		{-1, 1, 0},
// 		// {-1, -1, 0},
// 	}

// 	for _, test := range tests {
// 		res := add(test.a, test.b)
// 		if res != test.expected {
// 			t.Errorf("Add %d and %d should be %d but it's %d", test.a, test.b, test.expected, res)
// 		}
// 	}

// }

// func TestAddSubtests(t *testing.T) {
// 	tests := []struct{ a, b, expected int }{
// 		{2, 3, 5},
// 		{0, 0, 0},
// 		{-1, 1, 0},
// 		// {-1, -1, 0},
// 	}

// 	for _, test := range tests {
// 		t.Run(fmt.Sprintf("Add(%d,%d)",test.a, test.b), func(t *testing.T) {
// 			res := add(test.a, test.b)
// 			if res != test.expected {
// 				t.Errorf("Add %d and %d should be %d but it's %d", test.a, test.b, test.expected, res)
// 			}
// 		})

// 	}

// }
