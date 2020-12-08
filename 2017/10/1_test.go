package _2017

import (
	"testing"
)

func TestMultiplyFirstTwoInList(t *testing.T) {
	type args struct {
		rawInput   string
		listLength int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given input",
			args: args{
				rawInput: "106,16,254,226,55,2,1,166,177,247,93,0,255,228,60,36",
			},
			want: 11413,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplyFirstTwoInList(tt.args.rawInput); got != tt.want {
				t.Errorf("MultiplyFirstTwoInList() = %v, want %v", got, tt.want)
			}
		})
	}
}
