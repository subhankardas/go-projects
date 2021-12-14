package basics_tests

import (
	"testing"
)

func TestDictionarySearch(test *testing.T) {
	dictionary := Dictionary{"word 1": "sample definition 1"}

	test.Run("search known word", func(test *testing.T) {
		expected := "sample definition 1"        // Given
		actual, _ := dictionary.Search("word 1") // When
		assertString(test, expected, actual)     // Then
	})

	test.Run("search unknown word", func(test *testing.T) {
		expected := NotFoundError             // Given
		_, err := dictionary.Search("word 2") // When
		assertError(test, expected, err)      // Then
	})
}

func TestDictionaryAdd(test *testing.T) {
	dictionary := Dictionary{}

	test.Run("add new word", func(test *testing.T) {
		err := dictionary.Add("word 1", "sample 1") // When
		assertError(test, nil, err)                 // Then
	})

	test.Run("does not add existing word", func(test *testing.T) {
		expected := WordExistsError                 // Given
		err := dictionary.Add("word 1", "sample 2") // When
		assertError(test, expected, err)            // Then
	})

	test.Run("search added word", func(test *testing.T) {
		expected := "sample 2" // Given

		dictionary.Add("word 2", "sample 2") // When
		actual, err := dictionary.Search("word 2")

		assertString(test, expected, actual) // Then
		assertError(test, nil, err)
	})
}

func TestDictionaryUpdate(test *testing.T) {
	dictionary := Dictionary{}

	test.Run("does not update word that does not exists", func(test *testing.T) {
		expected := WordDoesNotExistsError             // Given
		err := dictionary.Update("word 1", "sample 1") // When

		assertError(test, expected, err) // Then
	})

	test.Run("update existing word", func(test *testing.T) {
		expected := "sample 2" // Given
		dictionary.Add("word 2", "sample 1")
		err := dictionary.Update("word 2", "sample 2") // WHen
		actual, _ := dictionary.Search("word 2")

		assertString(test, expected, actual) // Then
		assertError(test, nil, err)
	})
}

func assertString(test *testing.T, expected, actual string) {
	test.Helper()
	if expected != actual {
		test.Errorf("expected %q and got %q", expected, actual)
	}
}

func assertError(test *testing.T, expected, actual error) {
	test.Helper()
	if expected != actual { // Check for error message
		test.Errorf("expected %q and got %q", expected, actual)
	}
}

/* Custom error implementations */
type DictionaryError string

const (
	NotFoundError          = DictionaryError("word not found")
	WordExistsError        = DictionaryError("word already exists")
	WordDoesNotExistsError = DictionaryError("word does not exists")
)

func (err DictionaryError) Error() string {
	return string(err)
}

/* Dictionary implementation */
type Dictionary map[string]string

func (dictionary Dictionary) Search(word string) (string, error) {
	definition, ok := dictionary[word]
	if !ok {
		return "", NotFoundError
	}
	return definition, nil
}

func (dictionary Dictionary) Add(word, definition string) error {
	_, err := dictionary.Search(word)

	switch err { // Handle errors
	case nil:
		return WordExistsError
	case NotFoundError:
		dictionary[word] = definition
	default:
		return err
	}

	return nil
}

func (dictionary Dictionary) Update(word, definition string) error {
	_, err := dictionary.Search(word)

	switch err {
	case NotFoundError:
		return WordDoesNotExistsError
	case nil:
		dictionary[word] = definition
	default:
		return err
	}

	return nil
}
