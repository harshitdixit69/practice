package Radian

import (
	"testing"
)

// func TestDegreeToRadian(t *testing.T) {
// 	data := []struct {
// 		n    float64
// 		want float64
// 	}{
// 		{45, math.Pi / 4}, {90, math.Pi / 2}, {180, math.Pi}, {270, 3 * math.Pi / 2},
// 	}

// 	for _, d := range data {
// 		if got := degreeToRadian(d.n); got != d.want {
// 			t.Errorf("Invalid value for N: %f, got: %f, want: %f", d.n, got, d.want)
// 		}
// 	}
// }

// func TestRadianToDegree(t *testing.T) {
// 	data := []struct {
// 		n    float64
// 		want float64
// 	}{
// 		{math.Pi / 4, 45}, {math.Pi / 2, 90}, {math.Pi, 180}, {3 * math.Pi / 2, 270},
// 	}

// 	for _, d := range data {
// 		if got := radianToDegree(d.n); got != d.want {
// 			t.Errorf("Invalid value for N: %f, got: %f, want: %f", d.n, got, d.want)
// 		}
// 	}
// }

func Test_degreeToRadian(t *testing.T) {
	type args struct {
		degree float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "sam",
			args: args{2},
			want: 0.03490658503988659,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := degreeToRadian(tt.args.degree); got != tt.want {
				t.Errorf("degreeToRadian() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_radianToDegree(t *testing.T) {
	type args struct {
		radian float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "sam",
			args: args{2},
			want: 114.59155902616465,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := radianToDegree(tt.args.radian); got != tt.want {
				t.Errorf("radianToDegree() = %v, want %v", got, tt.want)
			}
		})
	}
}
