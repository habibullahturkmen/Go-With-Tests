package main

import "testing"

func TestHello(t *testing.T) {
	type args struct {
		name     string
		language string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"saying hello to people", args{"Habib", "English"}, "Hello, Habib"},
		{"say 'Hello, World' when an empty string is supplied", args{"", "English"}, "Hello, world"},
		{"in Turkish", args{"Ahmet", "Turkish"}, "Selam, Ahmet"},
		{"in French", args{"Chris", "French"}, "Bonjour, Chris"},
		{"in Spanish", args{"Adam", "Spanish"}, "Hola, Adam"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Hello(test.args.name, test.args.language); got != test.want {
				t.Errorf("got %q want %q", got, test.want)
			}
		})
	}
}
