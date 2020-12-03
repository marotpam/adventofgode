package _2017

import (
	"reflect"
	"testing"
)

func TestCalculateDistance(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "first example",
			n:    12,
			want: 3,
		},
		{
			name: "second example",
			n:    23,
			want: 2,
		},
		{
			name: "third example",
			n:    1024,
			want: 31,
		},
		{
			name: "given input",
			n:    361527,
			want: 326,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateDistance(tt.n); got != tt.want {
				t.Errorf("CalculateDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMatrix(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want *matrix
	}{
		{
			name: "zero",
			n:    0,
			want: &matrix{
				layers: 0,
				size:   0,
			},
		}, {
			name: "one layer",
			n:    1,
			want: &matrix{
				layers: 1,
				size:   1,
			},
		}, {
			name: "two layers",
			n:    2,
			want: &matrix{
				layers: 2,
				size:   3,
			},
		},
		{
			name: "three layers",
			n:    23,
			want: &matrix{
				layers: 3,
				size:   5,
			},
		},
		{
			name: "last in three layers",
			n:    25,
			want: &matrix{
				layers: 3,
				size:   5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMatrixFitting(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMatrixFitting() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestFindPointOne(t *testing.T) {
	tests := []struct {
		name string
		m    *matrix
		want point
	}{
		{
			name: "only one",
			m:    NewMatrixFitting(1),
			want: point{
				x: 0,
				y: 0,
			},
		},
		{
			name: "one layer",
			m:    NewMatrixFitting(4),
			want: point{
				x: 1,
				y: 1,
			},
		},
		{
			name: "two layers",
			m:    NewMatrixFitting(13),
			want: point{
				x: 2,
				y: 2,
			},
		},
		{
			name: "three layers",
			m:    NewMatrixFitting(26),
			want: point{
				x: 3,
				y: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.findPointOne(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPointOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindPointInOuterLayer(t *testing.T) {
	tests := []struct {
		name string
		m    *matrix
		n    int
		want point
	}{
		{
			name: "in bottom row of two layers",
			m:    NewMatrixFitting(4),
			n:    8,
			want: point{
				x: 2,
				y: 1,
			},
		},
		{
			name: "in bottom row of three layers",
			m:    NewMatrixFitting(23),
			n:    25,
			want: point{
				x: 4,
				y: 4,
			},
		},
		{
			name: "in left column of three layers",
			m:    NewMatrixFitting(18),
			n:    20,
			want: point{
				x: 3,
				y: 0,
			},
		},
		{
			name: "in top row of three layers",
			m:    NewMatrixFitting(15),
			n:    14,
			want: point{
				x: 0,
				y: 3,
			},
		},
		{
			name: "in right column of three layers",
			m:    NewMatrixFitting(11),
			n:    10,
			want: point{
				x: 4,
				y: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.findPointInOuterLayer(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPointInOuterLayer() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
