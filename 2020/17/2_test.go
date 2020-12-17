package _2020

import "testing"

func TestCountActiveCubes4d(t *testing.T) {
	type args struct {
		rawInput string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given example",
			args: args{
				rawInput: `.#.
..#
###`,
			},
			want: 848,
		},
		{
			name: "given input",
			args: args{
				rawInput: `###...#.
.##.####
.####.##
###.###.
.##.####
#.##..#.
##.####.
.####.#.`,
			},
			want: 2324,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountActiveCubes4d(tt.args.rawInput); got != tt.want {
				t.Errorf("CountActiveCubes4d() = %v, want %v", got, tt.want)
			}
		})
	}
}
