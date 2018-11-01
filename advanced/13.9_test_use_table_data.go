package main

import "testing"

var tests = []struct{
	in string
	out string
}{
	{"in1", "exp1"},
	{"in2", "exp2"},
	{"in3", "exp3"},
}

func TestFunction(t *testing.T) {
	for i, tt := range tests {
		s := FuncToBeTested(tt.in)
		if s != tt.out {
			t.Errorf("%d. %q => %q, wanted %q", i, tt.in, s, tt.out)
		}
	}
}