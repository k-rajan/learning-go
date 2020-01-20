package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("rajan")
	want := "Hello, rajan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
