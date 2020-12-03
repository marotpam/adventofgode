package _2019

import "testing"

func TestCountSecondPasswords(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given example",
			args: args{
				a: 178416,
				b: 676461,
			},
			want: 1129,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSecondPasswords(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("CountSecondPasswords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchesSecondPassword(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "all digits are equal",
			n:    111111,
			want: false,
		}, {
			name: "with a decreasing pair",
			n:    223450,
			want: false,
		}, {
			name: "without a repeated digit",
			n:    123789,
			want: false,
		}, {
			name: "only two are repeated",
			n:    112233,
			want: true,
		}, {
			name: "only a group of 3 repeated numbers",
			n:    123444,
			want: false,
		}, {
			name: "a group with 2 and a group with more than 2",
			n:    111122,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchesSecondPassword(tt.n); got != tt.want {
				t.Errorf("matchesSecondPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
