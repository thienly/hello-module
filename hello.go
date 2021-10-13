package hello

import (
	"fmt"

	quoteV3 "rsc.io/quote/v3"
)

func Hello() string {
	return quoteV3.HelloV3()
}
func Proverb() string {
	fmt.Printf("Hello world")
	return quoteV3.Concurrency()
}
