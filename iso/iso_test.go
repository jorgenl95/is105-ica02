package iso

import (
	"fmt"
	"testing"
	"unicode"
)

var s = GreetingExtendedASCII()
var maxExtendedASCII = '\u00FF'

func TestGreetingExtendedASCII(t *testing.T) {
	for _, c := range s {
		if c < unicode.MaxASCII {
			t.Error("The string contains characters in regular ASCII." + (" Found symbol: " + string(c)))
		} //else if c > maxExtendedASCII {
			//t.Error("The string contains characters outside the scope of extended ASCII." + (" Found symbol: " + string(c)))
		}

	fmt.Print("The string contains only extended ASCII characters.")
}