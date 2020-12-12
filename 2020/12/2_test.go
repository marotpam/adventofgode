package _2020

import "testing"

func TestCalculateManhattanDistance2(t *testing.T) {
	type args struct {
		rawInput string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given example",
			args: args{
				rawInput: `F10
N3
F7
R90
F11`,
			},
			want: 286,
		},
		{
			name: "given input",
			args: args{
				rawInput: `F93
R90
F81
E3
F80
R90
W4
R90
F64
N1
R90
W4
F11
W1
L180
F28
R90
W4
L180
F95
N3
F71
L90
W1
F1
E2
L90
F85
W3
L90
F28
E3
N5
F47
S2
R90
N2
L90
W4
L180
W3
F51
F77
L90
W5
N5
W3
S5
W5
R180
N1
W3
S5
F36
W1
F34
N4
F40
E2
S3
R90
W5
S2
N4
R90
S2
L90
N4
L90
S2
E1
F2
N4
F65
N1
F46
R180
F60
N4
F45
R90
S3
L90
F70
W4
L270
N4
F49
E3
F52
R180
F5
E5
R90
F43
L90
S4
F54
N1
F7
S2
F91
S4
N3
E3
S4
L180
F15
S2
F90
W5
S2
F80
N4
L90
R90
N4
E2
R180
E5
F62
L90
N5
F77
N5
F75
E2
L90
S4
F55
E1
N5
F57
R90
F6
L90
N3
E5
L180
N2
L270
N5
F8
E2
F88
E3
N3
E4
R90
N1
W5
R90
N4
E3
S3
R90
N1
E2
F4
N4
R90
N3
W5
L180
E1
F2
N4
W5
F80
S4
F7
L180
F96
L90
F16
L90
E4
F78
W2
R90
S3
F29
W4
R90
E4
F39
N1
F48
E4
L180
E1
R180
E2
R90
N3
R180
W2
N5
W5
E3
F25
E1
L180
N5
F44
L90
N4
W4
L90
F72
L90
S1
E5
N5
R90
N1
W3
S5
R90
W1
F14
L90
W2
R90
W3
F76
S5
W5
F93
W4
R90
F57
E3
R90
S3
W1
R90
S3
F8
R90
N2
F46
W4
S1
L90
E4
W5
L90
W2
F69
N5
W5
F80
N3
E4
L90
E4
F25
S3
R180
F77
R90
W2
F19
E4
L180
W2
F37
S2
F68
L90
E2
F66
S1
R90
F66
E2
L180
F97
N3
W4
F43
S4
R180
N1
R270
E3
N2
N3
F65
L90
S2
L90
N3
L90
S3
F23
L270
W3
S5
E2
R90
S1
F85
N3
R90
W4
F58
E1
L90
N3
L90
E2
S1
F14
E2
N5
W1
N3
E1
L90
E3
F43
E3
N3
F21
E4
F53
E2
L180
E4
F20
E2
N2
E5
L90
N4
W3
N4
S2
L90
W3
F96
L90
S3
R90
N1
E3
S5
L270
F41
N5
W5
S1
W5
R90
F79
W3
L90
E3
F22
N1
L90
E2
L90
R90
F20
L90
W3
R90
W2
L180
W4
F57
R180
N3
L90
F36
L90
E2
R90
N2
E3
N1
W4
W4
N3
E5
F54
R180
F98
W1
R90
F21
S1
L90
S2
L90
S2
F90
E4
S2
R90
N5
F25
N4
W3
N2
F27
S3
E3
N3
F15
L180
S4
F62
W2
L180
E2
N2
L90
R90
F97
R90
S1
R90
E2
F16
W2
E1
F89
W1
L180
S3
W2
S3
E1
F92
F30
N1
E2
S1
F76
E1
S5
W5
F28
W4
L90
F44
E4
N5
F25
R90
F59
S4
F58
S2
F19
W5
S4
E5
N3
F37
E1
L90
F40
E5
F56
S2
W5
F73
N5
F2
L90
F18
E2
N5
L90
F56
R90
F18
W1
S5
E1
N1
L90
W2
E4
N4
E1
W4
N4
L90
N3
R90
W5
S5
F2
R180
F96
R90
W3
F26
L90
S4
E2
F43
S2
R90
F61
W5
F93
R90
F95
L90
E4
R90
N4
F47
R90
W5
L90
F42
L90
W5
F87
R90
N1
N2
E4
F64
S5
L270
F86
S4
W3
S1
L90
F72
R270
W1
F17
S1
E2
R90
W1
N1
F42
N5
L90
F87
F66
L90
N2
W2
L90
S5
F7
R90
N4
L270
F2
W1
N4
F94
W5
R180
S4
F15
E4
F76
N1
E2
F68
S3
F50
R270
E5
F77
R90
S3
E2
N3
S4
F39
S1
E5
S3
L90
E4
S2
W3
F54
R90
F44
L90
W3
F59
R90
N3
F37
E1
F75
R90
F31
W3
F70
S5
L90
E1
F67
W4
L90
S1
W5
S1
S4
W2
E2
S3
R90
S2
E5
L90
F43
R90
E1
S5
F42
L90
W1
N2
E1
L180
S4
W1
L90
F81
E1
R180
N5
R180
N5
L180
F65
R90
F64
W4
S4
R90
F70
E1
S1
F50
E2
S4
E4
N2
S1
R90
E4
R90
F70
R90
N4
F71
R180
F80
S3
L90
N5
L180
F11
N1
R90
W2
N1
R90
W4
R90
F67
W3
R90
W5
L90
E4
F90
L90
N4
L180
F48
R90
W2
F94
R90
N4
L90
W2
F2
L90
W1
E5
S2
W5
S5
E2
N3
W5
N1
F98
S3
W3
L90
S3
W4
R90
W1
F64
N5
R90
S5
W5
F84
S1
E3
L90
S3
E5
F6
N4
W1
R90
E4
F14
N1
R90
F31
L90
F24
F4
N4
F54
S3
R270
F98
E1
L180
F2
E4
F70
W1
R180
N5
F23`,
			},
			want: 29895,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateManhattanDistance2(tt.args.rawInput); got != tt.want {
				t.Errorf("CalculateManhattanDistance2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_position_rotateRight(t *testing.T) {
	t.Run("90 degrees once", func(t *testing.T) {
		pos := position{x: 10, y: -4}
		expected := position{x: 4, y: 10}

		pos.rotateRight(90)

		if expected != pos {
			t.Errorf("expected %+v, got %+v", expected, pos)
		}
	})
	t.Run("90 degrees twice", func(t *testing.T) {
		pos := position{x: 10, y: -4}
		expected := position{x: -10, y: 4}

		pos.rotateRight(90)
		pos.rotateRight(90)

		if expected != pos {
			t.Errorf("expected %+v, got %+v", expected, pos)
		}
	})
	t.Run("90 degrees thrice", func(t *testing.T) {
		pos := position{x: 10, y: -4}
		expected := position{x: -4, y: -10}

		pos.rotateRight(90)
		pos.rotateRight(90)
		pos.rotateRight(90)

		if expected != pos {
			t.Errorf("expected %+v, got %+v", expected, pos)
		}
	})
	t.Run("90 degrees four times", func(t *testing.T) {
		pos := position{x: 10, y: -4}
		expected := position{x: 10, y: -4}

		pos.rotateRight(90)
		pos.rotateRight(90)
		pos.rotateRight(90)
		pos.rotateRight(90)

		if expected != pos {
			t.Errorf("expected %+v, got %+v", expected, pos)
		}
	})
	t.Run("180 degrees once", func(t *testing.T) {
		pos := position{x: 10, y: -4}
		expected := position{x: -10, y: 4}

		pos.rotateRight(180)

		if expected != pos {
			t.Errorf("expected %+v, got %+v", expected, pos)
		}
	})
	t.Run("180 degrees twice", func(t *testing.T) {
		pos := position{x: 10, y: -4}
		expected := position{x: 10, y: -4}

		pos.rotateRight(180)
		pos.rotateRight(180)

		if expected != pos {
			t.Errorf("expected %+v, got %+v", expected, pos)
		}
	})
	t.Run("270 degrees once", func(t *testing.T) {
		pos := position{x: 10, y: -4}
		expected := position{x: -4, y: -10}

		pos.rotateRight(270)

		if expected != pos {
			t.Errorf("expected %+v, got %+v", expected, pos)
		}
	})
}

func Test_position_rotateLeft(t *testing.T) {
	t.Run("90 degrees once", func(t *testing.T) {
		pos := position{x: 4, y: 10}
		expected := position{x: 10, y: -4}

		pos.rotateLeft(90)

		if expected != pos {
			t.Errorf("expected %+v, got %+v", expected, pos)
		}
	})
	t.Run("90 degrees twice", func(t *testing.T) {
		pos := position{x: 4, y: 10}
		expected := position{x: -4, y: -10}

		pos.rotateLeft(90)
		pos.rotateLeft(90)

		if expected != pos {
			t.Errorf("expected %+v, got %+v", expected, pos)
		}
	})
	t.Run("90 degrees thrice", func(t *testing.T) {
		pos := position{x: 4, y: 10}
		expected := position{x: -10, y: 4}

		pos.rotateLeft(90)
		pos.rotateLeft(90)
		pos.rotateLeft(90)

		if expected != pos {
			t.Errorf("expected %+v, got %+v", expected, pos)
		}
	})
	t.Run("90 degrees four times", func(t *testing.T) {
		pos := position{x: 4, y: 10}
		expected := position{x: 4, y: 10}

		pos.rotateLeft(90)
		pos.rotateLeft(90)
		pos.rotateLeft(90)
		pos.rotateLeft(90)

		if expected != pos {
			t.Errorf("expected %+v, got %+v", expected, pos)
		}
	})
	t.Run("180 degrees once", func(t *testing.T) {
		pos := position{x: 4, y: 10}
		expected := position{x: -4, y: -10}

		pos.rotateLeft(180)

		if expected != pos {
			t.Errorf("expected %+v, got %+v", expected, pos)
		}
	})
	t.Run("270 degrees once", func(t *testing.T) {
		pos := position{x: 4, y: 10}
		expected := position{x: -10, y: 4}

		pos.rotateLeft(270)

		if expected != pos {
			t.Errorf("expected %+v, got %+v", expected, pos)
		}
	})
}