package LevenshteinDistance

import (
	"testing"
)

func TestLevenshteinDistance(t *testing.T) {
	data := []struct {
		n    string
		k    string
		want int
	}{
		{"", "a", 1}, {"a", "", 1}, {"a", "a", 0}, {"islander", "slander", 1}, {"mart", "karma", 3}, {"football", "foot", 4},
	}

	for _, d := range data {
		if got := levenshteinDistance(d.n, d.k); got != d.want {
			t.Errorf("Invalid value for N: %v, got: %v, want: %d", d.n, got, d.want)
		}
	}
}
