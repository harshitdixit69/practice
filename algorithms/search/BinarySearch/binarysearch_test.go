package BinarySearch

import "testing"

// func SortArray(array []int) {
// 	for itemIndex, itemValue := range array {
// 		for itemIndex != 0 && array[itemIndex-1] > itemValue {
// 			array[itemIndex] = array[itemIndex-1]
// 			itemIndex -= 1
// 		}
// 		array[itemIndex] = itemValue
// 	}
// }

// func TestBinarySearch(t *testing.T) {
// 	random := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	array := make([]int, random.Intn(100-10)+10)
// 	fmt.Println("array", array)
// 	for i := range array {
// 		array[i] = random.Intn(100)
// 	}

// 	SortArray(array)

// 	for _, value := range array {
// 		result := binarySearch(array, value)
// 		if result == -1 {
// 			t.Fail()
// 		}
// 	}
// }

func Test_binarySearch(t *testing.T) {
	type args struct {
		arr   []int
		query int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sam",
			args: args{[]int{1, 2, 3, 4, 5, 6}, 5},
			want: 4,
		},
		{
			name: "sam",
			args: args{[]int{1, 2, 3, 4, 5, 6}, 1},
			want: 0,
		},
		{
			name: "sam",
			args: args{[]int{1, 2, 3, 4, 5, 6}, 7},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.arr, tt.args.query); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
