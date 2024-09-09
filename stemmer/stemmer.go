package stemmer

import (
	"strings"
	"fmt"
	"context"
	"net/http"
	"time"
)

var (
	suffixes = []string{"ed", "ing", "s"}
)

// Stem returns the stemmed version of word.
func Stem(word string) string {
	for _, suffix := range suffixes {
		if strings.HasSuffix(word, suffix) {
			return word[:len(word)-len(suffix)]
		}
	}
	return word
}

func BiggestNumber(numbers []int)(int){
	fmt.Println("only for testing")
	max := numbers[0]
	for _, num := range(numbers){
		if max<num {
			max=num
		}
	}
	return max
}

func CheckURL(url string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return false
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false
	}

	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}