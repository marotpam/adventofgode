package _2020

import "testing"

func TestGetResult(t *testing.T) {
	type args struct {
		rawInput string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first example",
			args: args{
				rawInput: `16
10
15
5
1
11
7
19
6
12
4`,
			},
			want: 35,
		}, {
			name: "second example",
			args: args{
				rawInput: `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`,
			},
			want: 220,
		}, {
			name: "given input",
			args: args{
				rawInput: `107
13
116
132
24
44
56
69
28
135
152
109
42
112
10
43
122
87
49
155
175
71
39
173
50
156
120
145
176
45
149
148
15
1
68
9
168
131
150
59
83
167
3
169
6
123
174
81
138
72
157
144
65
75
33
19
140
160
16
57
93
90
8
58
98
130
141
114
84
29
22
94
113
129
108
36
14
115
102
151
78
139
170
82
2
70
126
101
25
62
95
104
23
163
32
103
121
119
48
166
7
53`,
			},
			want: 2470,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetResult(tt.args.rawInput); got != tt.want {
				t.Errorf("GetResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
