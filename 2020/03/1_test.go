package _2020

import "testing"

func TestCountTrees(t *testing.T) {
	type args struct {
		rawInput string
		slope    slope
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given example",
			args: args{
				rawInput: `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`,
				slope: slope{
					incRow: 1,
					incCol: 3,
				},
			},
			want: 7,
		}, {
			name: "",
			args: args{
				rawInput: `.........#.#.#.........#.#.....
...#......#...#.....#.....#....
#.....#.....#.....#.#........#.
......#..#......#.......#......
.#..........#.............#...#
............#..##.......##...##
....#.....#..#....#............
.#..#.........#....#.#....#....
#.#...#...##..##.#..##..#....#.
.#.......#.#...#..........#....
...#...#........##...#..#.....#
..................#..........#.
.....#.##..............#.......
........#....##..##....#.......
...#.....#.##..........#...##..
.......#.#....#............#...
..............#......#......#..
#.......#...........#........##
.......#.......##......#.......
................#....##...#.#.#
#.......#....................#.
.##.#..##..#..#.#.....#.....#..
#...#............#......##....#
.#....##.#......#.#......#..#..
..........#........#.#.#.......
...#...#..........#..#....#....
..#.#...#...#...##...##......#.
......#...#........#.......###.
....#...............#.###...#.#
..................#.....#..#.#.
.#...#..#..........#........#..
#..........##................##
...#.....#...#......#.#......#.
......#..........#.#......#..#.
..#......#.....................
............#.........##.......
......#.......#........#.......
#.#...#...........#.......#....
.#.#........#.#.#....#........#
#.....##........#.#.....#.#....
.#...#..........##...#.....#..#
.........#....###............#.
..#...#..............#.#.#.....
.....#.#.#..#.#.#.###......#.#.
.#.##...#.......###..#.........
.....##....#.##..#.##..#.......
..#...........##......#..#.....
................##...#.........
##.....................#..#.###
...#..#....#...........#.......
.#.............##.#.....#.#....
.......#.#.#....##..#....#...#.
...##..#..........#....#.......
....##......#.........#........
.##....#...........#.#.......#.
...#...#.##.......#.#..........
..#.........#.##...........#...
....#.##........#.......#...##.
...................#..#..#...##
#...#......###..##.#..###......
#.............#.#........#...#.
...#...#..#..#..............#..
#.....#..#...#...#......#.#..#.
...##.............#........##.#
......#.#.........#.#....#...#.
........##............#...#....
............#.#...#......#.....
...#...........#...#...........
.........#.#......#............
....#.............#..#.....#..#
#.....#...........#.....#.#.#.#
.............#.....##......#...
...................###.#......#
#.##.....#...#..#..#...#....#..
............#.....#....#.#.....
#....#..#..........#....#..#...
..........##..................#
....#.......###..#......###....
.......#...#.##.##..#....##....
...##...............#.....#...#
#...........#...#......#..#..#.
..##....#.......#...#.....#....
.......##..............#.##..#.
.#......#..........#.......#...
....##...................#.#..#
......#....###................#
.#........#...........#........
.#.....##....#..##...........#.
##..............##...#.......#.
......#..........#..........##.
..#.....#.#.........####....#..
.............#......#......#...
..#.............#...........##.
#.#...#........#..........##...
...#....#.....#.....#.....##...
......#........................
......#..#...#......#.....#....
......#....##.....#....#.......
#.#......#.##..#...............
..........#.#.##..##..#........
......#.#..#....###.#..........
........................#....#.
#.#.............#....#.....##.#
.#.#..........#.......#..#....#
..#...#......#..#..#...#.#...#.
...#.##...###..................
........#.#...........#...#....
........#..........#....#......
#....#........#.......##.....#.
......###...##...#......#......
............#.......#.....##..#
....#..#.......##......#.......
#............##................
...............#..#......#...#.
.#....##...#......#............
#...#................#.........
..#....#..#........##......#...
....#....###......##.#.......#.
......#.#..#.#..#....#..#......
....#..........#..#..#.........
.#..#.......#......#........#..
.......#..#..#............#....
.............#...#....#..#.....
..............#.#.#.........#..
#.....##.......#....#..........
...#.....#......#..............
...##.#..#.#........#..##....#.
.......#.#.....##..#...........
.....#.#....#..................
.#......#.###..............##..
..#....#...#..#...##...##....#.
..........##..#...#..#.........
..#............#........#.#...#
.........................#.#.#.
......#........#.#..#......##.#
#.#......#..#.........#........
.....#........#......#...#.#...
........##....##....#.#........
....#....#.#...#...##....#..#..
#.............#.....#..........
#............##..#............#
..#.#......#........#..........
.#......#......#.#.##.##.......
..#.....#..........#......##...
...#......#...#.##....#.....##.
......#......#...........#.#...
....#........#..#..#........#.#
....#.........#.....#...#.#.#..
....#.....###........#.........
.............##........#.#.....
...#............#........#.#.#.
......#....#.......#.#.........
.....#................#........
.#....#..#.#.............#...#.
#..##...#............#......#..
...#..#........................
.#.#...........#.......#.......
#....###............##.........
...##....#.#............##.....
.........####......#........#..
.....#.......#.#...............
.......#...#.###..#....#...#..#
...#.....##..#....#..#.#...###.
.............#........#.#.#..#.
................#..........##..
.......####..##..##........##.#
..#......#..#..#.#.##..........
#....#........#....#...###.#.#.
........##........##.....#.....
...........#.#...........#....#
#.............#...........#....
...#.........#...#....#.....#..
..##......#...#...............#
.............#.........#....#..
..#...........#..#........#.##.
.#.#......#.............##...#.
.#......#.....##.#..#..#.......
....##......#..................
.#.#..##............#....#....#
........#...##.............#..#
........#....##.....#......###.
.........#....#.#..............
#.....#........................
.#..#....#.....#......#.###..#.
..........#...#....##....#..#..
...#.#.#...##..#..#......#..#.#
#............#.....#....#......
#.###...#.#......###..#....#..#
...#.###........#......#....#..
..#.##...#.........#.#.........
............##.................
....#..........#.#.#.#.#....#..
...##.#...#.......#.......##..#
....##.#........#....#...#.....
.............#.#....#...#.....#
...#......................##...
..#...#.....#.....#........#..#
..#..#.......#....#..##.....#..
..#..#.#.......................
.......##..#....#....#..#......
....#......##....#............#
.#...#..#..#.##...#.#...#......
.....#......#....#.........#...
.##......##.........#....#.....
#...........#...##.....#......#
.....#.#.......#.........#.....
.........#..........#..####.##.
............#..#......#.#......
.#.............#........#.#....
......#......#...#..#....#####.
.........##.#..##...###..#....#
....#.#....#.#..#.........#....
..#.............#...##...##....
........#..........#.##..#....#
.....#...#..##........#.#..#...
##..#.#.....#............#.....
.............#........##...##..
#......####.....##.............
..##.....##....###..#.#....#...
......##.##.#...#..#.#..##.....
......#.................#......
#.....#.#...#......#.#....#....
....#.#........#..............#
##........#.......##.#...##...#
..#..................#.#....#..
...........#..........#.#.....#
........##.#.....#......#..#..#
.....#....#..#.....#.........##
#.#..#..#...#......#..........#
#...##.....#..#.#.......#.##...
..#....##...............#......
#..........#.#.........#.#....#
..............#......#....#....
.....#...........#...#...#...#.
...#......#....#....#..........
.#..........#.#....##..##....#.
..............#.........#.#....
.......#.....#.....#...##....#.
##.#.........#....#.....#.#....
....#..#......#................
......##.....#.......##........
.....##...#........#...#...#...
..#...#...#..#..#.#......#..#..
....#...#.......#..............
....#..#.........###........#..
....#.............##..#........
..........##.#.......##..##....
#.##..................#.....#..
#........#........#.....#......
.#...#......#..................
#....##.##......#...#.........#
......#.##..##................#
............#.........##.......
..........####.#........#.....#
.##...#...#....#..#............
.#.##...#..#...#......#......##
.....#.#....#..###......#.#.#..
...#.......................##..
......................#.......#
..#....#.........#..#.#.....#..
.#....#..#....#...#............
..........#...##.....#.#..#....
........#..#..#....#...#...#...
.....#......#.#................
.....#...........#...#.........
.....#...##..#.#....#..#.....#.
#.......#.............##.......
................#....#.#..#....
.#..##...#.#........#......#.#.
.#.##..........#...............
....##......#....#........#....
....#..#....#.##.#.............
.......#..#......##.#.....#....
.......#.....#.............#...
.....#....#.......#............
........#.#...##..##..##.......
#.........##....##...##........
........#..#.#..........###.#..
..........................#.#..
#.....#.......#..#........#....
...##.....#.......#......#.....
.#.#..#...........#...........#
.....##..#........#...####.....
.#.#...##.#.#..#..#.#..#.......
..#.##.#...#.#.#...#..#........
............#..........#..#....
...............#..##.#.........
.............#.....#....#......
...##..##......##..........#...
..#.......#....#..........#...#
.##................#.#.#.......
.....##.....#..#.....#.........
......#.#.......#......#..#....
.....#.....#........#.......##.
......#.......##......#...#...#
....#...........###.........#..
...#.....#.........##........#.
..#.....#..............#.......
....#.......#...#....#....#..##
......#...........#...........#
.##......#......#.#.....#.##...
....#..##......#...#..#.#.###..
.......#.#....#......#..#......
..........#........#...........
#.##.........#.#.#...#...#.#...
.#......###.....#....#.#....#..
...................##..#.......
....#..#..............#.#.....#
#..................#.....#.....
...........##.##.......#..#.#..
........#.#......#...........#.
#..#.......#...#...........#.#.
......##...........#...........
.........#.#........#........#.
#......#....#.#.....#..#.......
............#..#.....##...#....
.#......#..#......#.........#..
.......#...#.........#.##.....#
........................#..#...
.###..............#.#..#.......
.....#.........#.......#......#
..##..##....#.....#.......#.#..
...###.#..#.##............#....`,
				slope: slope{
					incRow: 1,
					incCol: 3,
				},
			},
			want: 282,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountTreesInSlope(tt.args.rawInput, tt.args.slope); got != tt.want {
				t.Errorf("CountTreesInSlope() = %v, want %v", got, tt.want)
			}
		})
	}
}
