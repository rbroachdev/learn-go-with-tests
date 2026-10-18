package hello_world

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola, Chris"
		asserCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string and no language is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		asserCorrectMessage(t, got, want)
	})

	t.Run("say 'Bonjour, World' when an empty string is supplied and French language is supplied", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, World"
		asserCorrectMessage(t, got, want)
	})
}

func asserCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
