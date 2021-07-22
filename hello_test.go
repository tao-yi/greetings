package greetings

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	name := "world"
	want := fmt.Sprintf("Hi, %v. Welcome!", name)

	got, err := Hello(name)
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
