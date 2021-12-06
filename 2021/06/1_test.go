package _2021

import "testing"

func TestCountLanternfishAfterNDays(t *testing.T) {
	type args struct {
		lanternfish []int
		nDays       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given example with 18 days",
			args: args{
				lanternfish: []int{3, 4, 3, 1, 2},
				nDays:       18,
			},
			want: 26,
		},
		{
			name: "given example with 18 days",
			args: args{
				lanternfish: []int{3, 4, 3, 1, 2},
				nDays:       80,
			},
			want: 5934,
		},
		{
			name: "given input with 80 days",
			args: args{
				lanternfish: []int{3, 5, 1, 5, 3, 2, 1, 3, 4, 2, 5, 1, 3, 3, 2, 5, 1, 3, 1, 5, 5, 1, 1, 1, 2, 4, 1, 4, 5, 2, 1, 2, 4, 3, 1, 2, 3, 4, 3, 4, 4, 5, 1, 1, 1, 1, 5, 5, 3, 4, 4, 4, 5, 3, 4, 1, 4, 3, 3, 2, 1, 1, 3, 3, 3, 2, 1, 3, 5, 2, 3, 4, 2, 5, 4, 5, 4, 4, 2, 2, 3, 3, 3, 3, 5, 4, 2, 3, 1, 2, 1, 1, 2, 2, 5, 1, 1, 4, 1, 5, 3, 2, 1, 4, 1, 5, 1, 4, 5, 2, 1, 1, 1, 4, 5, 4, 2, 4, 5, 4, 2, 4, 4, 1, 1, 2, 2, 1, 1, 2, 3, 3, 2, 5, 2, 1, 1, 2, 1, 1, 1, 3, 2, 3, 1, 5, 4, 5, 3, 3, 2, 1, 1, 1, 3, 5, 1, 1, 4, 4, 5, 4, 3, 3, 3, 3, 2, 4, 5, 2, 1, 1, 1, 4, 2, 4, 2, 2, 5, 5, 5, 4, 1, 1, 5, 1, 5, 2, 1, 3, 3, 2, 5, 2, 1, 2, 4, 3, 3, 1, 5, 4, 1, 1, 1, 4, 2, 5, 5, 4, 4, 3, 4, 3, 1, 5, 5, 2, 5, 4, 2, 3, 4, 1, 1, 4, 4, 3, 4, 1, 3, 4, 1, 1, 4, 3, 2, 2, 5, 3, 1, 4, 4, 4, 1, 3, 4, 3, 1, 5, 3, 3, 5, 5, 4, 4, 1, 2, 4, 2, 2, 3, 1, 1, 4, 5, 3, 1, 1, 1, 1, 3, 5, 4, 1, 1, 2, 1, 1, 2, 1, 2, 3, 1, 1, 3, 2, 2, 5, 5, 1, 5, 5, 1, 4, 4, 3, 5, 4, 4},
				nDays:       80,
			},
			want: 351188,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountLanternfishAfterNDays(tt.args.lanternfish, tt.args.nDays); got != tt.want {
				t.Errorf("CountLanternfishAfterNDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
