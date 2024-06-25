package main

import (
	"fmt"
	"strings"

	link "github.com/sawilkhan/gophercises-html-parser"
)

var exampleHtml = `
<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
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