package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "world!")

	got := buffer.String()
	want := "Hello, world!"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
