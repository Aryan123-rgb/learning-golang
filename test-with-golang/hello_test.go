package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people in different languages", func(t *testing.T) {
		for lang, greet := range Greetings {
			t.Run("language: " + lang, func(t *testing.T) {
				got := Hello("Aryan", lang)
				want := greet + " Aryan" 
				assertCorrectMessage(t, got, want)
			})
		}
	})

	t.Run("saying hello to the entire world if empty string is provided", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("passing a language not present in map", func(t *testing.T) {
		got := Hello("Aryan", "hindi")
		want := "Invalid Language"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}