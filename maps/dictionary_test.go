package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("test2")
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definiton := "this is just a test"
		err := dictionary.Add(word, definiton)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definiton)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definiton := "this is just a test"
		dictionary := Dictionary{word: definiton}
		err := dictionary.Add(word, "newDefinition")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definiton)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definiton := "this is just a test"
		dictionary := Dictionary{word: definiton}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}
	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("expected %q to be deleted", word)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definiton string) {
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != definiton {
		t.Errorf("got %v, want %v", got, definiton)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
