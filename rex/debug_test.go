package rex

import "testing"

func TestStringify(t *testing.T) {
	tokens, err := parse("(a|b)?c*d+\\?")
	if err != nil {
		t.Fail()
	}
	want := "(((<a>|<b>))){0,1}(c){0,-1}(d){1,-1}?"
	if have := tokens.stringify(); have != want {
		t.Fatalf("want %s, have %s", want, have)
	}
}
