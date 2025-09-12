package main

import (
	"fmt"

	"github.com/tobiashort/utils-go/xstrings"
)

func main() {
	fmt.Println("-------------")
	s := `Lorem ipsum dolor sit amet,
		 |consectetur adipiscing elit.
		 |Curabitur justo tellus, facilisis nec efficitur dictum,
		 |fermentum vitae ligula. Sed eu convallis sapien.`
	fmt.Println(s)
	fmt.Println("-------------")
	fmt.Println(xstrings.Dedent(s))
	fmt.Println("-------------")
}
