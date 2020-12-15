package _2020

import "testing"

func TestGet2020th(t *testing.T) {
	type args struct {
		rawInput string
		wanted int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given example part 1",
			args: args{
				rawInput: "0,3,6",
				wanted: 2020,
			},
			want: 436,
		},
		{
			name: "given input part 1",
			args: args{
				"0,20,7,16,1,18,15",
				2020,
			},
			want: 1025,
		},
		{
			name: "given example part 2",
			args: args{
				rawInput: "0,3,6",
				wanted: 30000000,
			},
			want: 175594,
		},
		{
			name: "given input part 2",
			args: args{
				"0,20,7,16,1,18,15",
				30000000,
			},
			want: 129262,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNth(tt.args.rawInput, tt.args.wanted); got != tt.want {
				t.Errorf("GetNth() = %v, want %v", got, tt.want)
			}
		})
	}
}
