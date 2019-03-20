package ascii

import (
	"testing"
	"fmt"
	"unicode"
	)

var s = GreetingASCII()

func TestGreetingASCII(t *testing.T) {
	for _, c := range s {
		if c > unicode.MaxASCII {
			t.Error("The string contains characters not in ASCII.")
		}
	}
	fmt.Print("The string contains only ASCII characters.")
}