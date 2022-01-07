package hello

import "testing"

func TestHello(t *testing.T) {
	h := Hello()

	if h != "Hello" {
		t.Errorf("Hello returned wrongly, got: %s, want: %s", h, "Hello")
	}
}
