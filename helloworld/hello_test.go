package helloworld

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Habib", "English")
		want := "Hello, Habib"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want:= "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Turkish", func(t *testing.T) {
		got := Hello("Ahmet", "Turkish")
		want := "Selam, Ahmet"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Chris", "French")
		want := "Bonjour, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Adam", "Spanish")
		want := "Hola, Adam"
		assertCorrectMessage(t, got, want)
	})
}
