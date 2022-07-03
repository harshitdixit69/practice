package FastPowering

import "testing"

func Test_fastPowering(t *testing.T) {
	type args struct {
		base  float64
		power int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "cam",
			args: args{
				base:  2,
				power: 1,
			},
			want: 2,
		},
		{
			name: "cam1",
			args: args{
				base:  2,
				power: 2,
			},
			want: 4,
		},
		{
			name: "cam1",
			args: args{
				base:  2,
				power: 0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fastPowering(tt.args.base, tt.args.power); got != tt.want {
				t.Errorf("fastPowering() = %v, want %v", got, tt.want)
			}
		})
	}
}
