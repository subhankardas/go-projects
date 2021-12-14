package basics_tests

import (
	"reflect"
	"testing"
)

type WebsiteChecker func(string) bool

func CheckWebsites(check WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		results[url] = check(url)
	}
	return results
}

func mockWebsiteChecker(url string) bool {
	if url == "www.facebook.com" || url == "www.fb.com" {
		return false
	}
	return true
}

func TestCheckWebsites(test *testing.T) {
	websites := []string{"www.google.com", "www.facebook.com", "www.microsoft.com"}

	test.Run("check websites", func(t *testing.T) {
		expected := map[string]bool{
			websites[0]: true,
			websites[1]: false,
			websites[2]: true,
		}

		actual := CheckWebsites(mockWebsiteChecker, websites)

		if !reflect.DeepEqual(expected, actual) {
			test.Fatalf("expected %v but got %v", expected, actual)
		}
	})
}
