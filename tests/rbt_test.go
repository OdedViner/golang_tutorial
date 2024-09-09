package rbt

import (
	"fmt"
	"testing"
	
	"simple_app/stemmer"
)


func TestSecondBest(t *testing.T) {
	prices := []int{1, 0, 2, 3}
	p := stemmer.BiggestNumber(prices)
	fmt.Printf("res1=%d\n\n",p)
	expected := 3
	if p != expected {
		t.Fatalf("expected %d, got %d", p, expected)
	}
}

// go test -v tests/rbt_test.go 