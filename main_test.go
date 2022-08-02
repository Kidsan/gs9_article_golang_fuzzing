package main

import (
	"strings"
	"testing"
)

func FuzzRepeat(f *testing.F) { // fuzz tests use a `Fuzz` prefix instead of `Test`
	// Seed corpus
	f.Add("a", 50)
	f.Add("Geek Space 9", 6)
	f.Add("I Love Go!", 8)

	f.Fuzz(func(t *testing.T, input string, times int) { // takes a target function and passes it fuzzed parameters
		result := Repeat(input, times)

		if result != "" && strings.Count(result, input) != times {
			t.Errorf("Invalid result, input  %s got %s", input, result)
		}
	})
}
