package EuclideanAlgorithm

import (
	"testing"
)

func TestEuclideanAlgorithm(t *testing.T) {
	data := []struct {
		n    int
		k    int
		want int
	}{
		{0, 0, 0}, {2, 0, 2}, {0, 2, 2},
	}

	for _, d := range data {
		if got := GCD(d.n, d.k); got != d.want {
			t.Errorf("Invalid value for N: %d, got: %d, want: %d", d.n, got, d.want)
		}
	}
}
