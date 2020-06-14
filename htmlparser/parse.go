package htmlparser

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

/*
Struct containing content of html element
*/
type Link struct {
	Href string
	Text string
}

/*
Method to parse the html file, looking for links
*/
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

/*
Method for filter html elements
*/
func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}

/*
Method responsible for going down the tree looking for html elements
*/
func dfs(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}

}

/*
Build a link from a html element
*/
func buildLink(n *html.Node) Link {
	var link Link

	for _, att := range n.Attr {
		if att.Key == "href" {
			link.Href = att.Val
			break
		}
	}
	link.Text = text(n)
	return link
}

/*
Return the text inside a html element
*/
func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}
