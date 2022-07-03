package IsPowerOfTwo

import (
	"testing"
)

// func TestIsPowerOfTwo(t *testing.T) {
// 	data := []struct {
// 		n    int
// 		want bool
// 	}{
// 		{2, true}, {10, false}, {64, true}, {1024, true},
// 	}

// 	for _, d := range data {
// 		if got := isPowerOfTwo(d.n); got != d.want {
// 			t.Errorf("Invalid value for N: %d, got: %t, want: %t", d.n, got, d.want)
// 		}
// 	}
// }

// func TestIsPowerOfTwoBitwise(t *testing.T) {
// 	data := []struct {
// 		n    int
// 		want bool
// 	}{
// 		{2, true}, {10, false}, {64, true}, {1024, true},
// 	}

// 	for _, d := range data {
// 		if got := isPowerOfTwoBitwise(d.n); got != d.want {
// 			t.Errorf("Invalid value for N: %d, got: %t, want: %t", d.n, got, d.want)
// 		}
// 	}
// }

// func Test_isPowerOfTwo(t *testing.T) {
// 	type args struct {
// 		num int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := isPowerOfTwo(tt.args.num); got != tt.want {
// 				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_isPowerOfTwo(t *testing.T) {
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
				1,
			},
			want: false,
		},
		{
			name: "sam",
			args: args{
				2,
			},
			want: true,
		},
		{
			name: "sam",
			args: args{
				3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.args.num); got != tt.want {
				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
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
			args: args{
				-1, -1,
			},
			want: 1,
		},
		{
			name: "sam1",
			args: args{
				-1, 1,
			},
			want: 1,
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
