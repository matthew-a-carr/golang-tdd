package hello

import "testing"

func TestHello(t *testing.T) {

	assertCorrectErrorMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Matt", "")
		want := "Hello, Matt"
		assertCorrectErrorMessage(t, got, want)
	})

	t.Run("say 'Hello, World!' when an empty string is provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"
		assertCorrectErrorMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Sergio", "Spanish")
		want := "Hola, Sergio"
		assertCorrectErrorMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Thiery", "French")
		want := "Bonjour, Thiery"
		assertCorrectErrorMessage(t, got, want)
	})
}
