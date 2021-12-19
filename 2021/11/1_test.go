package _2021

import "testing"

func TestCountFlashes(t *testing.T) {
	type args struct {
		rawInput string
		steps    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given example with 100 steps",
			args: args{
				rawInput: `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`,
				steps: 100,
			},
			want: 1656,
		},
		{
			name: "given input with 100 steps",
			args: args{
				rawInput: `6318185732
1122687135
5173237676
8754362612
5718474666
8443654137
1247634346
1446514585
6717288267
1727871228`,
				steps: 100,
			},
			want: 1634,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountFlashes(tt.args.rawInput, tt.args.steps); got != tt.want {
				t.Errorf("CountFlashes() = %v, want %v", got, tt.want)
			}
		})
	}
}
