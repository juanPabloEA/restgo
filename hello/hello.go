package hello


import (
	"rsc.io/quote";
	"fmt"
)

func Hello() string {
	fmt.Println(quote.Hello())
    return quote.Hello()
}