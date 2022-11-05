package fizzbuzz

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"testing/quick"
	"time"
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

// TestFizzbuzzProp calls the default service and does property testing on it
func TestFizzbuzzProp(t *testing.T) {
    service := NewFizzbuzz()

    property := func (a, b int) bool {
        // skip these test cases since they are invalid
        if a < 1 || b < 1 {
            return true
        }
        a = a % 1000
        b = b % 1000

        req := FizzbuzzRequest{
            Int1: a,
            Int2: b,
            Limit: 1000,
            Str1: "a",
            Str2: "b",
        }
        resp, err := service.GetFizzbuzz(req)
        if err != nil {
            return false
        }

        fbs := strings.Split(resp, ",")
        if len(fbs) != req.Limit {
            t.Errorf("Returned length should be %v, got %v", req.Limit, len(fbs))
            return false
        }
        for i, fb := range fbs {
            idx := i + 1
            if idx % a == 0 && !strings.HasPrefix(fb, req.Str1) {
                t.Errorf("Expected prefix %q, got string %q", req.Str1, fb)
                return false
            }
            if idx % b == 0 && !strings.HasSuffix(fb, req.Str2) {
                t.Errorf("Expected suffix %q, got string %q", req.Str2, fb)
                return false
            }
        }
        return true
    }

    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    if err := quick.Check(property, &quick.Config{Rand: r}); err != nil {
        t.Errorf("Fizzbuzz tests failed in property testing: %v", err)
    }
}

func TestInvalidParams(t *testing.T) {
    testCases := []FizzbuzzRequest{
        {Int1: -1, Int2: 2, Limit: 10, Str1: "a", Str2: "b"},
        {Int1: 1, Int2: -2, Limit: 10, Str1: "a", Str2: "b"},
        {Int1: 1, Int2: 2, Limit: 10, Str1: "", Str2: "b"},
        {Int1: 1, Int2: 2, Limit: 10, Str1: "a", Str2: ""},
    }
    service := NewFizzbuzz()

    for _, tc := range testCases {
        if _, err := service.GetFizzbuzz(tc); err == nil {
            t.Errorf("Fizzbuzz didn't returned an error with following params: %q", formatRequest(tc))
        }
    }
}
