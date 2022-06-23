package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"asddsadkaksk", 4},
		{"asdjhadjkjdj", 5},
		{"djsakjdklas", 6},
	}
	for _, tt := range tests {
		if asdjk := lengthOfNunRepeatingSubStr(tt.s); asdjk != tt.ans {
			t.Errorf("lengthOfNunRepeatingSubStr(%s)"+"got %d;expected %d", tt.s, asdjk, tt.ans)
		}
	}
}
