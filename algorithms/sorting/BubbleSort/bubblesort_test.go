package BubbleSort

import (
	"reflect"
	"testing"
)

// func TestBubbleSort(t *testing.T) {
// 	random := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	array1 := make([]int, random.Intn(100-10)+10)

// 	for i := range array1 {
// 		array1[i] = random.Intn(100)
// 	}
// 	array2 := make(sort.IntSslice, len(array1))

// 	copy(array2, array1)
// 	arr := bubbleSort(array1)
// 	array2.Sort()

// 	for i := range arr {
// 		if arr[i] != array2[i] {
// 			t.Fail()
// 		}
// 	}
// }

func Test_bubbleSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "saam",
			args: args{[]int{5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bubbleSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
