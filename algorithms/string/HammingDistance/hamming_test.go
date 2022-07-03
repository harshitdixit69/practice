package HammingDistance

import "testing"

// func TestHammingDistance(t *testing.T) {
// 	data := []struct {
// 		n    string
// 		k    string
// 		want int
// 	}{
// 		{"a", "ab", 0}, {"a", "a", 0}, {"abc", "azf", 2}, {"1011101", "1001001", 2},
// 	}

// 	for _, d := range data {
// 		if got := hammingDistance(d.n, d.k); got != d.want {
// 			t.Errorf("Invalid value for N: %v, got: %v, want: %d", d.n, got, d.want)
// 		}
// 	}
// }

func Test_hammingDistance(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "sam",
			args: args{"a", "b"},
			want: 1,
		},
		{
			name: "sam",
			args: args{"a", "ab"},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingDistance(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("hammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
