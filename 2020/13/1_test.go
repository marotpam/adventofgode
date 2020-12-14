package _2020

import "testing"

func TestSolvePart1(t *testing.T) {
	type args struct {
		earliestTimestamp int
		rawBusIDs         string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given example",
			args: args{
				earliestTimestamp: 939,
				rawBusIDs:         "7,13,x,x,59,x,31,19",
			},
			want: 295,
		},
		{
			name: "given input",
			args: args{
				earliestTimestamp: 1006697,
				rawBusIDs:         "13,x,x,41,x,x,x,x,x,x,x,x,x,641,x,x,x,x,x,x,x,x,x,x,x,19,x,x,x,x,17,x,x,x,x,x,x,x,x,x,x,x,29,x,661,x,x,x,x,x,37,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,23",
			},
			want: 3966,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolvePart1(tt.args.earliestTimestamp, tt.args.rawBusIDs); got != tt.want {
				t.Errorf("SolvePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}
