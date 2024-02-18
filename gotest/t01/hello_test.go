package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		name := "Chris"
		got := Hello(name)

		want := fmt.Sprintf("Hello, %s", name)

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to the World", func(t *testing.T) {
		got := Hello("")

		want := fmt.Sprintf("Hello, World")

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
