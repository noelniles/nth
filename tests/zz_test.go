package tests

import (
    "fmt"
    "testing"
    "github.com/noelniles/nth/numbers"
)

func TestNewZZ (t *testing.T) {
    fmt.Println("Testing NewZZ.")
    z := numbers.NewZZ("1110000000000000000000000000000000000000000")
    fmt.Println(z)
}
