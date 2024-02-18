package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Chris"
	got := Hello(name)

	want := fmt.Sprintf("Hello, %s!", name)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
