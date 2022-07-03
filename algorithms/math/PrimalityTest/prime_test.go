package PrimalityTest

import (
	"testing"
)

// func TestIsPrimeNumber(t *testing.T) {
// 	data := []struct {
// 		n    int
// 		want bool
// 	}{
// 		{3, true}, {10, false}, {5, true}, {1024, false}, {29, true}, {7, true}, {11, true},
// 	}

// 	for _, d := range data {
// 		if got := isPrimeNumber(d.n); got != d.want {
// 			t.Errorf("Invalid value for N: %d, got: %t, want: %t", d.n, got, d.want)
// 		}
// 	}
// }

// func TestIsPrimeNumberUsingSqrt(t *testing.T) {
// 	data := []struct {
// 		n    int
// 		want bool
// 	}{
// 		{3, true}, {10, false}, {5, true}, {1024, false}, {29, true}, {7, true}, {11, true},
// 	}

// 	for _, d := range data {
// 		if got := isPrime(d.n); got != d.want {
// 			t.Errorf("Invalid value for N: %d, got: %t, want: %t", d.n, got, d.want)
// 		}
// 	}
// }

func Test_isPrimeNumber(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "sam",
			args: args{
				5,
			},
			want: true,
		},
		{
			name: "sam1",
			args: args{
				4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrimeNumber(tt.args.num); got != tt.want {
				t.Errorf("isPrimeNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mod(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sam",
			args: args{2, 2},
			want: 0,
		},
		{
			name: "sam1",
			args: args{-1, -1},
			want: 1,
		},
		{
			name: "sam3",
			args: args{-2, 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mod(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("mod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPrime(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "sam",
			args: args{5},
			want: true,
		},
		{
			name: "sam1",
			args: args{16},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrime(tt.args.num); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
