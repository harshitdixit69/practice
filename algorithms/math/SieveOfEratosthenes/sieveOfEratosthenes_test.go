package SieveOfEratosthenes

import (
	"reflect"
	"testing"
)

// func TestSieveOfEratosthenes(t *testing.T) {
// 	t.Run("Prime Numbers until 10", func(t *testing.T) {
// 		got := sieveOfEratosthenes(10)
// 		want := []int{
// 			2, 3, 5, 7,
// 		}

// 		if len(got) != len(want) {
// 			t.Errorf("Different array size then expected")
// 		}

// 		if !Equal(got, want) {
// 			t.Errorf("Error")
// 		}
// 	})

// 	t.Run("Prime Numbers until 50", func(t *testing.T) {
// 		got := sieveOfEratosthenes(50)
// 		want := []int{
// 			2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47,
// 		}

// 		if len(got) != len(want) {
// 			t.Errorf("Different array size then expected")
// 		}

// 		if !Equal(got, want) {
// 			t.Errorf("Error")
// 		}
// 	})

// 	t.Run("Prime Numbers until 100", func(t *testing.T) {
// 		got := sieveOfEratosthenes(100)
// 		want := []int{
// 			2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
// 		}

// 		if len(got) != len(want) {
// 			t.Errorf("Different array size then expected")
// 		}

// 		if !Equal(got, want) {
// 			t.Errorf("Error")
// 		}
// 	})
// }

// func Equal(a, b []int) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}

// 	for i, v := range a {
// 		if v != b[i] {
// 			return false
// 		}
// 	}

// 	return true
// }

func Test_sieveOfEratosthenes(t *testing.T) {
	type args struct {
		maxNumber int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "sam",
			args: args{5},
			want: []int{2,3,5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sieveOfEratosthenes(tt.args.maxNumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sieveOfEratosthenes() = %v, want %v", got, tt.want)
			}
		})
	}
}
