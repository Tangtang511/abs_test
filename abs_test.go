package main

import "testing"

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
