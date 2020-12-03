package _2019

import "testing"

func TestCalculateBiodiversityRating(t *testing.T) {
	w := newWorld(`
.....
.....
.....
#....
.#...`)

	if got := w.calculateBiodiversityRating(); got != 2129920 {
		t.Errorf("calculateBiodiversityRating() = %v, want %v", got, 2129920)
	}
}

func TestGetFirstRepeatedRating(t *testing.T) {
	tests := []struct {
		name      string
		rawLayout string
		want      int
	}{
		{
			name: "input for first part",
			rawLayout: `
#..##
##...
.#.#.
#####
####.`,
			want: 23967691,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFirstRepeatedRating(tt.rawLayout); got != tt.want {
				t.Errorf("GetFirstRepeatedRating() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTick(t *testing.T) {
	tests := []struct {
		name     string
		layout   map[position]bool
		wantBug  bool
		position position
	}{
		{
			name: "a bug will die when alone",
			layout: map[position]bool{
				position{0, 0}: true,
			},
			wantBug:  false,
			position: position{0, 0},
		}, {
			name: "a bug will die when surrounded by more than one",
			layout: map[position]bool{
				position{0, 0}: true,
				position{1, 0}: true,
				position{0, 1}: true,
			},
			wantBug:  false,
			position: position{0, 0},
		}, {
			name: "a bug will die when surrounded by four",
			layout: map[position]bool{
				position{1, 1}: true,
				position{1, 0}: true,
				position{1, 2}: true,
				position{0, 1}: true,
				position{2, 1}: true,
			},
			wantBug:  false,
			position: position{1, 1},
		}, {
			name: "a bug will survive when surrounded by exactly one",
			layout: map[position]bool{
				position{0, 0}: true,
				position{1, 0}: true,
			},
			wantBug:  true,
			position: position{0, 0},
		}, {
			name: "an empty space will be empty when not surrounded by any bug",
			layout: map[position]bool{
				position{0, 0}: false,
			},
			wantBug:  false,
			position: position{0, 0},
		}, {
			name: "an empty space will be empty when surrounded by three bugs",
			layout: map[position]bool{
				position{1, 1}: false,
				position{0, 1}: true,
				position{2, 1}: true,
				position{1, 2}: true,
			},
			wantBug:  false,
			position: position{1, 1},
		}, {
			name: "an empty space will be empty when surrounded by four bugs",
			layout: map[position]bool{
				position{1, 1}: false,
				position{0, 1}: true,
				position{2, 1}: true,
				position{1, 2}: true,
				position{1, 0}: true,
			},
			wantBug:  false,
			position: position{1, 1},
		}, {
			name: "an empty space will become infected when surrounded by a bug",
			layout: map[position]bool{
				position{0, 0}: false,
				position{0, 1}: true,
			},
			wantBug:  true,
			position: position{0, 0},
		}, {
			name: "an empty space will become infected when surrounded by two bugs",
			layout: map[position]bool{
				position{0, 0}: false,
				position{0, 1}: true,
				position{1, 1}: true,
			},
			wantBug:  true,
			position: position{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &world{
				layout: tt.layout,
			}
			w.tick()
			if got := w.hasBugIn(tt.position); got != tt.wantBug {
				t.Errorf("tick() = %v, want %v", got, tt.wantBug)
			}
		})
	}
}
