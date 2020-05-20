package main

import (
	"testing"
)

var valid = []struct {
	in  string
	out string
}{
	{"++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.+++++++++++++++++++++++++++++.+++++++..+++.-------------------------------------------------------------------------------.+++++++++++++++++++++++++++++++++++++++++++++++++++++++.++++++++++++++++++++++++.+++.------.--------.-------------------------------------------------------------------.-----------------------.", "Hello World!\n"},
	{"++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.", "Hello World!\n"},
	{"+[-[<<[+[--->]-[<<<]]]>>>-]>-.---.>..>.<<<<-.<+.>>>>>.>.<<.<-.", "hello world"},
}

func TestValidInput(t *testing.T) {
	for _, tt := range valid {
		t.Run(tt.in, func(t *testing.T) {
			s, err := interpret(tt.in)
			if err != nil {
				t.Error(err)
			}
			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}

var invalid = []struct {
	in  string
	out error
}{
	{"1+++.<>,[]", nil},
}

func TestInvalidInput(t *testing.T) {
	for _, tt := range invalid {
		t.Run(tt.in, func(t *testing.T) {
			_, err := interpret(tt.in)
			if err == tt.out {
				t.Error("got 'nil', want 'error'")
			}
		})
	}
}
