package fizzbuzz

import (
	"fmt"
	"testing"
)

// formatRequest to show in errors
func formatRequest(req FizzbuzzRequest) string {
    return fmt.Sprintf("int1=%v,int2=%v,str1=%q,str2=%q", req.Int1, req.Int2, req.Str1, req.Str2)
}

// TestFizzbuzz calls the default service with classic fizzbuzz parameters
func TestFizzbuzz(t *testing.T) {
	req := FizzbuzzRequest{
		Int1:  3,
		Int2:  5,
		Limit: 20,
		Str1:  "fizz",
		Str2:  "buzz",
	}

	service := NewFizzbuzz()

	resp, err := service.GetFizzbuzz(req)
	expected := "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,17,fizz,19,buzz"
    if err != nil {
        t.Fatalf(`Fizzbuzz returned an unexpected error: %v`, err)
    }
	if resp != expected {
        t.Fatalf(`Fizzbuzz with params %q, returned %q, but we expected %q`, formatRequest(req), resp, expected)
	}
}
