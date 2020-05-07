package main

import (
	"fmt"

	"github.com/technosophos/moniker"
)

func main() {
	n := moniker.New()
	fmt.Printf("%s\n", n.NameSep("-"))
}
