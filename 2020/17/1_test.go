package _2020

import "testing"

func TestCountActiveCubes(t *testing.T) {
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
			want: 112,
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
			want: 319,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountActiveCubes(tt.args.rawInput); got != tt.want {
				t.Errorf("CountActiveCubes() = %v, want %v", got, tt.want)
			}
		})
	}
}
