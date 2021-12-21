package _2021

import "testing"

func TestSubtractCommonQuantities(t *testing.T) {
	type args struct {
		rawInput  string
		stepCount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given example with 10 steps",
			args: args{
				rawInput: `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`,
				stepCount: 10,
			},
			want: 1588,
		},
		{
			name: "given input with 10 steps",
			args: args{
				rawInput: `BCHCKFFHSKPBSNVVKVSK

OV -> V
CO -> V
CS -> O
NP -> H
HH -> P
KO -> F
VO -> B
SP -> O
CB -> N
SB -> F
CF -> S
KS -> P
OH -> H
NN -> O
SF -> K
FH -> F
VV -> B
VH -> O
BV -> V
KF -> K
CC -> F
NF -> H
VS -> O
SK -> K
HV -> O
CK -> K
VP -> F
HP -> S
CN -> K
OB -> H
NS -> F
PS -> S
KB -> S
VF -> S
FP -> H
BB -> N
HF -> V
CH -> N
BH -> F
KK -> B
OO -> N
NO -> K
BP -> K
KH -> P
KN -> P
OF -> B
VC -> F
NK -> F
ON -> O
OC -> P
VK -> O
SH -> C
NH -> C
FB -> B
FC -> K
OP -> O
PV -> V
BN -> V
PC -> K
PK -> S
FF -> C
SV -> S
HK -> H
NB -> C
OK -> C
PH -> B
SO -> O
PP -> F
KV -> V
FO -> B
FN -> H
HN -> C
VB -> K
CV -> O
BC -> C
CP -> S
FS -> S
KP -> V
BS -> V
BK -> B
PN -> C
PF -> S
HO -> V
NC -> N
SS -> N
BO -> P
BF -> N
NV -> P
PB -> K
HB -> H
VN -> H
FV -> B
FK -> K
PO -> S
SC -> S
HS -> S
KC -> F
HC -> S
OS -> K
SN -> N`,
				stepCount: 10,
			},
			want: 2797,
		},
		{
			name: "given input with 40 steps",
			args: args{
				rawInput: `BCHCKFFHSKPBSNVVKVSK

OV -> V
CO -> V
CS -> O
NP -> H
HH -> P
KO -> F
VO -> B
SP -> O
CB -> N
SB -> F
CF -> S
KS -> P
OH -> H
NN -> O
SF -> K
FH -> F
VV -> B
VH -> O
BV -> V
KF -> K
CC -> F
NF -> H
VS -> O
SK -> K
HV -> O
CK -> K
VP -> F
HP -> S
CN -> K
OB -> H
NS -> F
PS -> S
KB -> S
VF -> S
FP -> H
BB -> N
HF -> V
CH -> N
BH -> F
KK -> B
OO -> N
NO -> K
BP -> K
KH -> P
KN -> P
OF -> B
VC -> F
NK -> F
ON -> O
OC -> P
VK -> O
SH -> C
NH -> C
FB -> B
FC -> K
OP -> O
PV -> V
BN -> V
PC -> K
PK -> S
FF -> C
SV -> S
HK -> H
NB -> C
OK -> C
PH -> B
SO -> O
PP -> F
KV -> V
FO -> B
FN -> H
HN -> C
VB -> K
CV -> O
BC -> C
CP -> S
FS -> S
KP -> V
BS -> V
BK -> B
PN -> C
PF -> S
HO -> V
NC -> N
SS -> N
BO -> P
BF -> N
NV -> P
PB -> K
HB -> H
VN -> H
FV -> B
FK -> K
PO -> S
SC -> S
HS -> S
KC -> F
HC -> S
OS -> K
SN -> N`,
				stepCount: 40,
			},
			want: 2797,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubtractCommonQuantities(tt.args.rawInput, tt.args.stepCount); got != tt.want {
				t.Errorf("SubtractCommonQuantities() = %v, want %v", got, tt.want)
			}
		})
	}
}
