package main

import (
	"fmt"

	"github.com/tobiashort/utils-go/ternary"
)

func main() {
	age := 18
	access := ternary.IfThenElse(age < 18, "Access denied", "Access granted")
	fmt.Println(access)
}
