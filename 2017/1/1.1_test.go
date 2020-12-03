package _2017

import "testing"

func TestSolveFirstCaptcha(t *testing.T) {
	tests := []struct {
		name string
		captcha string
		want int
	}{
		{
			name:    "empty string",
			captcha: "",
			want:    0,
		},{
			name:    "string with one number",
			captcha: "1",
			want:    1,
		},{
			name:    "two consecutive numbers (1 and 2)",
			captcha: "1122",
			want:    3,
		},{
			name:    "a string with the same digit multiple times",
			captcha: "1111",
			want:    4,
		},{
			name:    "four consecutive numbers",
			captcha: "1234",
			want:    0,
		},{
			name:    "only last digit matching the first one",
			captcha: "91212129",
			want:    9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveFirstCaptcha(tt.captcha); got != tt.want {
				t.Errorf("SolveFirstCaptcha() = %v, want %v", got, tt.want)
			}
		})
	}
}