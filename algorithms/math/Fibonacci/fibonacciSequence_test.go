package Fibonacci

import (
	"reflect"
	"testing"
)

func Test_fibonacciSequence(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sam",
			args: args{
				2,
			},
			want: []int{0, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacciSequence(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fibonacciSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
