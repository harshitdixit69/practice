package Factorial

import "testing"

// func TestRecusriveFactorial(t *testing.T) {
// 	t.Run("Factorial of 0", func(t *testing.T) {
// 		got := FactorialRecursive(0)
// 		want := 1

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})

// 	t.Run("Factorial of 5", func(t *testing.T) {
// 		got := FactorialRecursive(5)
// 		want := 120

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})

// 	t.Run("Factorial of 8", func(t *testing.T) {
// 		got := FactorialRecursive(8)
// 		want := 40320

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})

// 	t.Run("Factorial of 10", func(t *testing.T) {
// 		got := FactorialRecursive(10)
// 		want := 3628800

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})
// }

func TestFactorialRecursive(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sam",
			args: args{1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FactorialRecursive(tt.args.num); got != tt.want {
				t.Errorf("FactorialRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
