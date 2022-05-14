package main

import (
	"testing"
	"unicode/utf8"
)

// writing unit test
func TestReverse(t *testing.T) {
	testcases := []struct {
		input           string
		expected_output string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}


	for _,tc:=range testcases{
	  rev:=Reverse(tc.input)

	  if rev!=tc.expected_output{
	    t.Errorf("Reverse: %q, want %q",rev,tc.expected_output)
	  }
	}
}
// Add a fuzz test
// The unit test has limitations, namely that each input must be added to the test by the developer. One benefit of fuzzing is that it comes up with inputs for your code, and may identify edge cases that the test cases you came up with didn’t reach.
//
// In this section you will convert the unit test to a fuzz test so that you can generate more inputs with less work!
//
// Note that you can keep unit tests, benchmarks, and fuzz tests in the same *_test.go file, but for this example you will convert the unit test to a fuzz test.
//

func FuzzReverse(f *testing.F){
	testcases:= []string{"Hello, world"," ","!12345"}

	for _,tc:=range testcases{
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string){
		rev:=Reverse(orig)
		doubleRev:= Reverse(rev)

		if orig!=doubleRev{
			t.Errorf("Before: %q, after: %q",orig,doubleRev)
		}

		if utf8.ValidString(orig) && utf8.ValidString(rev){
			t.Errorf("Reverse produced invalid UTF-8 string %q",rev)
		}
	})
}
