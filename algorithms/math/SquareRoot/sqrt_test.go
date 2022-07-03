package SquareRoot

import (
	"testing"
)

// func TestSqrt(t *testing.T) {
// 	data := []struct {
// 		n    float64
// 		want float64
// 	}{
// 		{4.0, 2.0}, {9.0, 3.0}, {36.0, 6.0}, {81.0, 9.0}, {256.0, 16.0},
// 	}

// 	for _, d := range data {
// 		if got := squareRoot(d.n); got != d.want {
// 			t.Errorf("Invalid value for N: %f, got: %f, want: %f", d.n, got, d.want)
// 		}
// 	}
// }

func Test_squareRoot(t *testing.T) {
	type args struct {
		num float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "sam",
			args: args{4},
			want: 2,
		},
		{
			name: "sam",
			args: args{-1},
			want: -1,
		},
		{
			name: "sam",
			args: args{0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squareRoot(tt.args.num); got != tt.want {
				t.Errorf("squareRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
