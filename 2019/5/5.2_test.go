package main

import (
	"reflect"
	"testing"
)

func TestCalculateSecondOpcode(t *testing.T) {
	type args struct {
		ints []int
		in   *fakeInput
		out  *fakeOutput
	}
	tests := []struct {
		name           string
		args           args
		want           []int
		expectedOutput []int
	}{
		{
			name: "given input",
			args: args{
				ints: []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1001, 191, 50, 224, 101, -64, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 5, 224, 224, 1, 224, 223, 223, 2, 150, 218, 224, 1001, 224, -1537, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 2, 224, 1, 223, 224, 223, 1002, 154, 5, 224, 101, -35, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 5, 224, 1, 224, 223, 223, 1102, 76, 17, 225, 1102, 21, 44, 224, 1001, 224, -924, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 4, 224, 1, 224, 223, 223, 101, 37, 161, 224, 101, -70, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 6, 224, 224, 1, 223, 224, 223, 102, 46, 157, 224, 1001, 224, -1978, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 5, 224, 1, 224, 223, 223, 1102, 5, 29, 225, 1101, 10, 7, 225, 1101, 43, 38, 225, 1102, 33, 46, 225, 1, 80, 188, 224, 1001, 224, -73, 224, 4, 224, 102, 8, 223, 223, 101, 4, 224, 224, 1, 224, 223, 223, 1101, 52, 56, 225, 1101, 14, 22, 225, 1101, 66, 49, 224, 1001, 224, -115, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 7, 224, 1, 224, 223, 223, 1101, 25, 53, 225, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 108, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 329, 101, 1, 223, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 344, 1001, 223, 1, 223, 8, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 359, 101, 1, 223, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 374, 101, 1, 223, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 389, 101, 1, 223, 223, 7, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 404, 1001, 223, 1, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 419, 1001, 223, 1, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 434, 101, 1, 223, 223, 1008, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 449, 1001, 223, 1, 223, 1007, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 464, 1001, 223, 1, 223, 1008, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 479, 101, 1, 223, 223, 1007, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 494, 1001, 223, 1, 223, 108, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 509, 101, 1, 223, 223, 8, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 524, 1001, 223, 1, 223, 107, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 539, 101, 1, 223, 223, 107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 554, 101, 1, 223, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 569, 1001, 223, 1, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 584, 1001, 223, 1, 223, 1008, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 599, 1001, 223, 1, 223, 1107, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 614, 101, 1, 223, 223, 7, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 629, 1001, 223, 1, 223, 1108, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 644, 1001, 223, 1, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 659, 101, 1, 223, 223, 1108, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226},
				in:   &fakeInput{inputs: []int{5}},
				out:  &fakeOutput{outputs: []int{}},
			},
			expectedOutput: []int{12410607},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CalculateSecondOpcode(tt.args.ints, tt.args.in, tt.args.out)
			if len(tt.expectedOutput) > 0 && !reflect.DeepEqual(tt.expectedOutput, tt.args.out.outputs) {
				t.Error("output does not match expectation", tt.expectedOutput, tt.args.out.outputs)
			}
		})
	}
}

func TestJumping(t *testing.T) {
	t.Run("first example with input 8", func(t *testing.T) {
		in := fakeInput{inputs: []int{8}}
		out := fakeOutput{outputs: []int{}}
		example := []int{3,9,8,9,10,9,4,9,99,-1,8}

		CalculateSecondOpcode(example, &in, &out)

		if len(out.outputs) == 0 {
			t.Errorf("nothing was written to the output")
		}

		if out.outputs[0] != 1 {
			t.Errorf("expecting to read value in input (1), got %d", out.outputs[0])
		}
	})
	t.Run("first example with input different than 8", func(t *testing.T) {
		in := fakeInput{inputs: []int{42}}
		out := fakeOutput{outputs: []int{}}
		example := []int{3,9,8,9,10,9,4,9,99,-1,8}

		CalculateSecondOpcode(example, &in, &out)

		if len(out.outputs) == 0 {
			t.Errorf("nothing was written to the output")
		}

		if out.outputs[0] != 0 {
			t.Errorf("expecting to read value in input (1), got %d", out.outputs[0])
		}
	})
	t.Run("second example with input lower than 8", func(t *testing.T) {
		in := fakeInput{inputs: []int{3}}
		out := fakeOutput{outputs: []int{}}
		example := []int{3,9,7,9,10,9,4,9,99,-1,8}

		CalculateSecondOpcode(example, &in, &out)

		if len(out.outputs) == 0 {
			t.Errorf("nothing was written to the output")
		}

		if out.outputs[0] != 1 {
			t.Errorf("expecting to read value in input (1), got %d", out.outputs[0])
		}
	})
	t.Run("second example with input greater than 8", func(t *testing.T) {
		in := fakeInput{inputs: []int{42}}
		out := fakeOutput{outputs: []int{}}
		example := []int{3,9,7,9,10,9,4,9,99,-1,8}

		CalculateSecondOpcode(example, &in, &out)

		if len(out.outputs) == 0 {
			t.Errorf("nothing was written to the output")
		}

		if out.outputs[0] != 0 {
			t.Errorf("expecting to read value in input (0), got %d", out.outputs[0])
		}
	})
	t.Run("third example with input 8", func(t *testing.T) {
		in := fakeInput{inputs: []int{8}}
		out := fakeOutput{outputs: []int{}}
		example := []int{3,3,1108,-1,8,3,4,3,99}

		CalculateSecondOpcode(example, &in, &out)

		if len(out.outputs) == 0 {
			t.Errorf("nothing was written to the output")
		}

		if out.outputs[0] != 1 {
			t.Errorf("expecting to read value in input (1), got %d", out.outputs[0])
		}
	})
	t.Run("third example with input different than 8", func(t *testing.T) {
		in := fakeInput{inputs: []int{42}}
		out := fakeOutput{outputs: []int{}}
		example := []int{3,3,1108,-1,8,3,4,3,99}

		CalculateSecondOpcode(example, &in, &out)

		if len(out.outputs) == 0 {
			t.Errorf("nothing was written to the output")
		}

		if out.outputs[0] != 0 {
			t.Errorf("expecting to read value in input (0), got %d", out.outputs[0])
		}
	})
	t.Run("fourth example with input lower than 8", func(t *testing.T) {
		in := fakeInput{inputs: []int{3}}
		out := fakeOutput{outputs: []int{}}
		example := []int{3,3,1107,-1,8,3,4,3,99}

		CalculateSecondOpcode(example, &in, &out)

		if len(out.outputs) == 0 {
			t.Errorf("nothing was written to the output")
		}

		if out.outputs[0] != 1 {
			t.Errorf("expecting to read value in input (1), got %d", out.outputs[0])
		}
	})
	t.Run("fifth example with input below 8", func(t *testing.T) {
		in := fakeInput{inputs: []int{1}}
		out := fakeOutput{outputs: []int{}}
		example := []int{3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,
			1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,
			999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99}

		CalculateSecondOpcode(example, &in, &out)

		if len(out.outputs) == 0 {
			t.Errorf("nothing was written to the output")
		}

		if out.outputs[0] != 999 {
			t.Errorf("expecting to read value in input (999), got %d", out.outputs[0])
		}
	})
	t.Run("fifth example with input 8", func(t *testing.T) {
		in := fakeInput{inputs: []int{8}}
		out := fakeOutput{outputs: []int{}}
		example := []int{3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,
			1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,
			999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99}

		CalculateSecondOpcode(example, &in, &out)

		if len(out.outputs) == 0 {
			t.Errorf("nothing was written to the output")
		}

		if out.outputs[0] != 1000 {
			t.Errorf("expecting to read value in input (1000), got %d", out.outputs[0])
		}
	})
	t.Run("fifth example with input greater than 8", func(t *testing.T) {
		in := fakeInput{inputs: []int{100}}
		out := fakeOutput{outputs: []int{}}
		example := []int{3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,
			1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,
			999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99}

		CalculateSecondOpcode(example, &in, &out)

		if len(out.outputs) == 0 {
			t.Errorf("nothing was written to the output")
		}

		if out.outputs[0] != 1001 {
			t.Errorf("expecting to read value in input (1001), got %d", out.outputs[0])
		}
	})
}