package greetings

import (
    "testing"
	"regexp"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value
func TestHelloName(t *testing.T) {
	name := "GladOS"
	want := regexp.MustCompile('\b' + name + '\b')
	msg, err := Hello("GladOS")
	if !want.MatchString(msg) || err != nil {
			t.Errorf('Hello("GladOS") = %q, %v, want match for #q, nil', msg, err, want)
	}
}
