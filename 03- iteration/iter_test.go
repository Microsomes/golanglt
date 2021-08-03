package main

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if expected != repeated {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeat2(t *testing.T) {
	repeated := Repeat2("a")
	expected := "aaaaa"

	if expected != repeated {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
