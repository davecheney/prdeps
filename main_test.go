package main

import "testing"

var tests = []struct {
	wd   string
	want string
}{
	{"/Users/kevin/src/github.com/davecheney/prdeps", "github.com/davecheney/prdeps"},
	{"/Users/kevin/src", ""},
	{"/Users/kevin/src/1", "1"},
	{"/Users/kevin", ""},
	{"/Users/kevin/var", ""},
	{"/Users/kevin/var/tmp/foo", ""},
	{"/Users/blah/src/github.com/davecheney/prdeps", ""},
	{"/", ""},
	{"", ""},
	{"/Users/blah", ""},
}

func TestGetGoSubpath(t *testing.T) {
	for _, tt := range tests {
		got, err := getGoSubpath("/Users/kevin/src", tt.wd)
		if err == nil && got == "" {
			t.Errorf("getGoSubpath(%q, %q): want an error, got none", tt.wd, got)
		}
		if err != nil && got != "" {
			t.Errorf("getGoSubpath(%q, %q): want a result, got an error", tt.wd, got)
		}
		if got != tt.want {
			t.Errorf("getGoSubpath(%q, %q): got %t, want %t", "/Users/kevin/src", tt.wd, got, tt.want)
		}
	}
}
