package main

import "testing"

func TestHello(t *testing.T) {
	got := hello()
	want := "Hello World"

	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func TestAdios(t *testing.T) {
	got := adios()
	want := "adios"

	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}
