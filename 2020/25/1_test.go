package _2020

import "testing"

func Test_transform(t *testing.T) {
	type args struct {
		subjectNumber int
		loopSize      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "card public key",
			args: args{
				subjectNumber: 7,
				loopSize:      8,
			},
			want: 5764801,
		},
		{
			name: "door public key",
			args: args{
				subjectNumber: 17807724,
				loopSize:      8,
			},
			want: 14897079,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transform(tt.args.subjectNumber, tt.args.loopSize); got != tt.want {
				t.Errorf("transform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateLoopSize(t *testing.T) {
	type args struct {
		subjectNumber int
		wantedPK      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "card public key",
			args: args{
				subjectNumber: 7,
				wantedPK:      5764801,
			},
			want: 8,
		},
		{
			name: "door public key",
			args: args{
				subjectNumber: 17807724,
				wantedPK:      14897079,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateLoopSize(tt.args.subjectNumber, tt.args.wantedPK); got != tt.want {
				t.Errorf("calculateLoopSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CalculateEncryptionKey(t *testing.T) {
	type args struct {
		cardPK int
		doorPK int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given example",
			args: args{
				cardPK: 5764801,
				doorPK: 17807724,
			},
			want: 14897079,
		},
		{
			name: "given input",
			args: args{
				cardPK: 14012298,
				doorPK: 74241,
			},
			want: 18608573,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateEncryptionKey(tt.args.cardPK, tt.args.doorPK); got != tt.want {
				t.Errorf("CalculateEncryptionKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
