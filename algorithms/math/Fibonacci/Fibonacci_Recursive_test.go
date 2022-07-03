package Fibonacci

import "testing"

func TestFibonacciRecursive(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sam",
			args: args{
				2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibonacciRecursive(tt.args.n); got != tt.want {
				t.Errorf("FibonacciRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
