package main

import "testing"

func TestCountFirstPasswords(t *testing.T) {
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
			want: 1650,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountFirstPasswords(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("CountFirstPasswords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchesFirstPassword(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "all digits are equal",
			n:    111111,
			want: true,
		}, {
			name: "with a decreasing pair",
			n:    223450,
			want: false,
		}, {
			name: "without a repeated digit",
			n:    123789,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchesFirstPassword(tt.n); got != tt.want {
				t.Errorf("matchesFirstPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
