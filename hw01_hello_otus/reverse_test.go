package main

import "testing"

func TestReverseString(t *testing.T)  {
	cases := []struct {
		in, want string
	}{
		{"Hello, OTUS!", "!SUTO ,olleH"},
	}
	for _, c := range cases {
		got := reverseStr(c.in)
		if got != c.want {
			t.Errorf("reverseStr(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}