package IsPowerOfTwo

import "testing"

func Test_isPowerOfTwoBitwise(t *testing.T) {
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
				3,
			},
			want: false,
		},
		{
			name: "sam",
			args: args{
				-1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwoBitwise(tt.args.num); got != tt.want {
				t.Errorf("isPowerOfTwoBitwise() = %v, want %v", got, tt.want)
			}
		})
	}
}
