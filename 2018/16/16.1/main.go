package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

type opcode func(a, b, c int, registers []int) []int

type sample struct {
	input []int
	opcode, a, b, c int
	output []int
}

func countPossibleOpcodes(s sample) int {
	n := 0

	for _, op := range allOpCodes() {
		if reflect.DeepEqual(op(s.a, s.b, s.c, s.input), s.output) {
			n++
		}
	}

	return n
}

func allOpCodes() map[string]opcode {
	m := map[string]opcode{}
	m["addr"] = addr
	m["addi"] = addi
	m["mulr"] = mulr
	m["muli"] = muli
	m["banr"] = banr
	m["bani"] = bani
	m["borr"] = borr
	m["bori"] = bori
	m["setr"] = setr
	m["seti"] = seti
	m["gtir"] = gtir
	m["gtri"] = gtri
	m["gtrr"] = gtrr
	m["eqir"] = eqir
	m["eqri"] = eqri
	m["eqrr"] = eqrr

	return m
}

func addr(a, b, c int, registers []int) []int {
	output := append([]int{}, registers...)
	output[c] = output[a] + output[b]

	return output
}

func addi(a, b, c int, registers []int) []int {
	output := append([]int{}, registers...)
	output[c] = output[a] + b

	return output
}

func mulr(a, b, c int, registers []int) []int {
	output := append([]int{}, registers...)
	output[c] = output[a] * output[b]

	return output
}

func muli(a, b, c int, registers []int) []int {
	output := append([]int{}, registers...)
	output[c] = output[a] * b

	return output
}

func banr(a, b, c int, registers []int) []int {
	output := append([]int{}, registers...)
	output[c] = output[a] & output[b]
	return output
}

func bani(a, b, c int, registers []int) []int {
	output := append([]int{}, registers...)
	output[c] = output[a] & b

	return output
}

func borr(a, b, c int, registers []int) []int {
	output := append([]int{}, registers...)
	output[c] = output[a] | output[b]

	return output
}

func bori(a, b, c int, registers []int) []int {
	output := append([]int{}, registers...)
	output[c] = output[a] | b

	return output
}

func setr(a, b, c int, registers []int) []int {
	output := append([]int{}, registers...)
	output[c] = output[a]

	return output
}

func seti(a, b, c int, registers []int) []int {
	output := append([]int{}, registers...)
	output[c] = a

	return output
}

func gtir(a, b, c int, registers []int) []int {
	output := append([]int{}, registers...)
	output[c] = 0
	if a > registers[b] {
		output[c] = 1
	}

	return output
}

func gtri(a, b, c int, registers []int) []int {
	output := append([]int{}, registers...)
	output[c] = 0
	if registers[a] > b {
		output[c] = 1
	}

	return output
}

func gtrr(a, b, c int, registers []int) []int {
	output := append([]int{}, registers...)
	output[c] = 0
	if registers[a] > registers[b] {
		output[c] = 1
	}

	return output
}

func eqir(a, b, c int, registers []int) []int {
	output := append([]int{}, registers...)
	output[c] = 0
	if a == registers[b] {
		output[c] = 1
	}

	return output
}

func eqri(a, b, c int, registers []int) []int {
	output := append([]int{}, registers...)
	output[c] = 0
	if registers[a] == b {
		output[c] = 1
	}

	return output
}

func eqrr(a, b, c int, registers []int) []int {
	output := append([]int{}, registers...)
	output[c] = 0
	if registers[a] == registers[b] {
		output[c] = 1
	}

	return output
}

func readSamples() []sample {
	fh, _ := os.Open("input.txt")
	defer fh.Close()

	fs := bufio.NewScanner(fh)
	var initialState string
	fmt.Sscanf(fs.Text(), `initial state: %s`, &initialState)


	samples := []sample{}
	for fs.Scan() {
		var i0, i1, i2, i3 int
		fmt.Sscanf(fs.Text(), `Before: [%d, %d, %d, %d]`, &i0, &i1, &i2, &i3)

		fs.Scan()
		var opcode, a, b, c int
		fmt.Sscanf(fs.Text(), `%d %d %d %d`, &opcode, &a, &b, &c)

		fs.Scan()
		var o0, o1, o2, o3 int
		fmt.Sscanf(fs.Text(), `After:  [%d, %d, %d, %d]`, &o0, &o1, &o2, &o3)

		s := sample{
			input:  []int{i0, i1, i2, i3},
			opcode: opcode,
			a:      a,
			b:      b,
			c:      c,
			output: []int{o0, o1, o2, o3},
		}
		samples = append(samples, s)
		fs.Scan()
	}
	return samples
}

func main() {
	c := 0
	for _, s := range readSamples() {
		if countPossibleOpcodes(s) >= 3 {
			c++
		}
	}
	fmt.Println(c) // 590
}
