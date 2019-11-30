package main

import "testing"

func TestSolveSecondCaptcha(t *testing.T) {
	tests := []struct {
		name string
		captcha string
		want int
	}{
		{
			name:    "four digits all matching the ones two digits ahead",
			captcha: "1212",
			want:    6,
		},{
			name:    "no digit matching",
			captcha: "1221",
			want:    0,
		},{
			name:    "only one digit matching",
			captcha: "123425",
			want:    4,
		},{
			name:    "123123",
			captcha: "123123",
			want:    12,
		},{
			name:    "12131415",
			captcha: "12131415",
			want:    4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveSecondCaptcha(tt.captcha); got != tt.want {
				t.Errorf("SolveSecondCaptcha() = %v, want %v", got, tt.want)
			}
		})
	}
}