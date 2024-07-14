package main

import "testing"

func TestEvenOrOdd(t *testing.T){
    cases := []struct {
        in int
        want string
    }{
        {10, "even"},
        {7, "odd"},
        {82, "even"},
        {99, "odd"},
    }

    for _, c := range cases {
        got := EvenOrOdd(c.in)
        if got != c.want {
            t.Errorf("expected: %q, actual %q", c.want, got)
        }
    }
}
