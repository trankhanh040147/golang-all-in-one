package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Chris"
	got := Hello(name)

	want := fmt.Sprintf("Hello, %s", name)

	assertCorrectMessage(t, got, want)
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
