package basics_tests

import (
	"reflect"
	"testing"
	"time"
)

type WebsiteChecker func(string) bool

type result struct {
	url     string
	checked bool
}

func CheckWebsites(check WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		results[url] = check(url)
	}
	return results
}

func CheckWebsitesWithConcurrency(check WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultschannel := make(chan result)

	for _, url := range urls {
		go func(url string) {
			resultschannel <- result{url, check(url)}
		}(url)
	}

	for idx := 0; idx < len(urls); idx++ {
		result := <-resultschannel
		results[result.url] = result.checked
	}
	return results
}

func mockWebsiteChecker(url string) bool {
	if url == "www.facebook.com" || url == "www.fb.com" {
		return false
	}
	return true
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsites(test *testing.T) {
	websites := []string{"www.google.com", "www.facebook.com", "www.microsoft.com"}
	expected := map[string]bool{
		websites[0]: true,
		websites[1]: false,
		websites[2]: true,
	}

	test.Run("check websites without concurrency", func(test *testing.T) {
		actual := CheckWebsites(mockWebsiteChecker, websites)
		if !reflect.DeepEqual(expected, actual) {
			test.Fatalf("expected %v but got %v", expected, actual)
		}
	})

	test.Run("check websites with concurrency", func(test *testing.T) {
		actual := CheckWebsitesWithConcurrency(mockWebsiteChecker, websites)
		if !reflect.DeepEqual(expected, actual) {
			test.Fatalf("expected %v but got %v", expected, actual)
		}
	})
}

func BenchmarkCheckWebsitesWithSlowChecker(bench *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < bench.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func BenchmarkCheckWebsitesWithSlowCheckerAndConcurrency(bench *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < bench.N; i++ {
		CheckWebsitesWithConcurrency(slowStubWebsiteChecker, urls)
	}
}
