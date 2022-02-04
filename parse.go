package link

import "io"
import (
	"golang.org/x/net/html"
	"fmt"
)
type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error){
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	dfs(doc, "")
	return nil, nil
}

func dfs(n *html.Node, padding string){
	fmt.Println(padding, n.Data)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding + " ")
	}
}