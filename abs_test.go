package main

import (
	"fmt"
	"testing"
)

func TestAbs(t *testing.T) {
	type args struct {
		strs []string
		str  string
	}
	testcases := []struct {
		name string
		args int
		want int
	}{
		{
			name: "positive",
			args: 1,
			want: 1,
		},
		{
			name: "negative",
			args: -5,
			want: 5,
		},
		{
			name: "zero",
			args: 0,
			want: 0,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			if res := Abs(tt.args); res != tt.want {
				t.Errorf("Abs() = %v, want %v", res, tt.want)
			}
		})
	}
}
func BenchmarkAbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Abs(6)
	}
}

func ExampleAbs() {
	fmt.Println(Abs(6))
	fmt.Println(Abs(0))
	fmt.Println(Abs(-2))
	// Output:
	// 6
	// 0
	// 2
}
