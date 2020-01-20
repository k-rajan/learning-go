package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMsg := func(t *testing.T, got string, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	}

	t.Run("saying hello by name", func(t *testing.T) {
		got := Hello("rajan", "English")
		want := "Hello, rajan"
		assertCorrectMsg(t, got, want)
	})

	t.Run("saying hello world when empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMsg(t, got, want)
	})

	t.Run("saying hello in spanish", func(t *testing.T) {
		got := Hello("Amigo", "Spanish")
		want := "Hola, Amigo"
		assertCorrectMsg(t, got, want)
	})
	t.Run("saying hello in french", func(t *testing.T) {
		got := Hello("Cornor", "French")
		want := "Bonjour, Cornor"
		assertCorrectMsg(t, got, want)
	})
}
