package maporsmth

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("search test", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add test", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		got, err := dictionary.Search("test")
		want := "this is just a test"

		if err != nil	{
			t.Fatal("got an error, should find the added word")
		}

		assertStrings(t, got, want)
	})
}

func assertStrings(t testing.TB, got string, want string) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
