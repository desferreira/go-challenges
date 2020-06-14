package main

import (
	"fmt"
	"strings"

	"github.com/desferreira/go-challenges/htmlparser"
)

var exampleHTML = `
<div>
<a href="/dog">
  <span>Something in a span</span>
  Text not in a span
  <b>Bold text!</b>
</a>
<a href="/cat">
<p> oi </p>
</a>
</div>
`

func main() {
	r := strings.NewReader(exampleHTML)

	links, err := htmlparser.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", links)

}
