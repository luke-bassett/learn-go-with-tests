package main

import "bytes"

func TestGreet(t *Testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Luke"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
