package main

import (
    "testing"
)

func TestEquality(t *testing.T) {
    a := 5
    b := 5
    if a != b {
        t.Errorf("Expected %d to equal %d", a, b)
    }
}

func TestEqualityFail (t *testing.T) {
    a := 5
    b := 6
    if a != b {
        t.Errorf("Expected %d to equal %d", a, b)
    }
}

