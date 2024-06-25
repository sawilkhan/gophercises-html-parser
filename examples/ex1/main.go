package main

import (
	"fmt"
	"strings"

	link "github.com/sawilkhan/gophercises-html-parser"
)

var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to first page</a>
  <a href="/other-page">A link to second page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := link.Parse(r)
	if err != nil{
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}