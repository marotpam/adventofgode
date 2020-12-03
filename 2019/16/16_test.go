package _2019

import (
	"reflect"
	"testing"
)

func TestGetPattern(t *testing.T) {
	type args struct {
		basePattern []int
		signal      string
		position    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first example in position 1",
			args: args{
				basePattern: []int{0, 1, 0, -1},
				signal:      "12345678",
				position:    1,
			},
			want: []int{1, 0, -1, 0, 1, 0, -1, 0},
		},
		{
			name: "first example in position 2",
			args: args{
				basePattern: []int{0, 1, 0, -1},
				signal:      "12345678",
				position:    2,
			},
			want: []int{0, 1, 1, 0, 0, -1, -1, 0},
		},
		{
			name: "first example in position 3",
			args: args{
				basePattern: []int{0, 1, 0, -1},
				signal:      "12345678",
				position:    3,
			},
			want: []int{0, 0, 1, 1, 1, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getPattern(tt.args.basePattern, tt.args.signal, tt.args.position)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculatePhase(t *testing.T) {
	type args struct {
		signal  string
		pattern []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first example until first phase",
			args: args{
				signal:  "12345678",
				pattern: []int{0, 1, 0, -1},
			},
			want: "48226158",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := applyPhase(tt.args.signal, tt.args.pattern); got != tt.want {
				t.Errorf("applyPhase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	type args struct {
		signal      string
		basePattern []int
		phases      int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first example until first phase",
			args: args{
				signal:      "12345678",
				basePattern: []int{0, 1, 0, -1},
				phases:      1,
			},
			want: "48226158",
		},
		{
			name: "first example until fourth phase",
			args: args{
				signal:      "12345678",
				basePattern: []int{0, 1, 0, -1},
				phases:      4,
			},
			want: "01029498",
		},
		{
			name: "given input",
			args: args{
				signal:      "59790677903322930697358770979456996712973859451709720515074487141246507419590039598329735611909754526681279087091321241889537569965210074382210124927546962637736867742660227796566466871680580005288100192670887174084077574258206307557682549836795598410624042549261801689113559881008629752048213862796556156681802163843211546443228186862314896620419832148583664829023116082772951046466358463667825025457939806789469683866009241229487708732435909544650428069263180522263909211986231581228330456441451927777125388590197170653962842083186914721611560451459928418815254443773460832555717155899456905676980728095392900218760297612453568324542692109397431554",
				basePattern: []int{0, 1, 0, -1},
				phases:      100,
			},
			want: "19239468144644194406346900233733234456077910448797259794288736379391422894251775535255171152993816670094740505691376991919293948380925358101946161570008338970912608868847806898306732417648346271575720245326923161028866513314234507059558112922104888876459802037687157715628133357571287799333173360649646508737254335051239167531111832172719673991665991968228860976767735052777947065046370536428634238365575484588363557754248005936479858668756797222649143735879531400639315233838306593646868207422252040658176423162191499292508208304508862282784331808634471847817865439222256585320864731847829353205316748467607534392722373352418068374047692104842931554",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateFlawedFrequencyTransmission(tt.args.signal, tt.args.basePattern, tt.args.phases); got != tt.want {
				t.Errorf("CalculateFlawedFrequencyTransmission() = %v, want %v", got, tt.want)
			}
		})
	}
}
