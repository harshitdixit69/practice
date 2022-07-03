package main

import (
	"reflect"
	"testing"
)

func Test_newFunction(t *testing.T) {
	type args struct {
		points [][]int
		vector []int
		arr    []int
		c      int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 map[int][]int
	}{
		// TODO: Add test cases.
		{
			name:  "sam",
			args:  args{[][]int{{1, 3}, {2, 9}, {3, 7}}, []int{2, 2}, []int{}, 2},
			want:  []int{2, 4, 7},
			want1: map[int][]int{2: {1, 3}, 4: {3, 7}, 7: {2, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := newFunction(tt.args.points, tt.args.vector, tt.args.arr, tt.args.c)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newFunction() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("newFunction() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "sam",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
