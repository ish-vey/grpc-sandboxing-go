package main

// for imported items that are not installed, run 'go mod tidy'
import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
