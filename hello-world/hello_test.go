package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("to a Person", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Halo, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Halo, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Indo", func(t *testing.T) {
		got := Hello("Chris", indo)
		want := "Halo, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Chris", spanish)
		want := "Hola, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Chris", french)
		want := "Bonjour, Chris"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTDDHello(t *testing.T) {
	testCases := []struct {
		testName string
		userName string
		language string
		want     string
	}{
		{testName: "to a Person", userName: "Chris", language: "", want: "Halo, Chris"},
		{testName: "empty string", userName: "", language: "", want: "Halo, World"},
		{testName: "in Indo", userName: "Chris", language: indo, want: "Halo, Chris"},
		{testName: "in Spanish", userName: "Chris", language: spanish, want: "Hola, Chris"},
		{testName: "in French", userName: "Chris", language: french, want: "Bonjour, Chris"},
	}

	for _, test := range testCases {
		got := Hello(test.userName, test.language)
		if got != test.want {
			assertCorrectMessage(t, got, test.want)
		}
	}
}
