package main

import "testing"

func Example_main() {
	main()
	// Output:
	// Hello World!
}

func Test_hello(t *testing.T) {
	want := "Hello World!"
	if got := hello(); got != want {
		t.Errorf("hello() = %q, want %q", got, want)
	}
}
