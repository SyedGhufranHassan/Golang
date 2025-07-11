package hello

import "testing"

func TestHello(t *testing.T) {
    got := Hello()
    want := "Hello, world." // or hardcode expected result instead of quote.Hello()

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}