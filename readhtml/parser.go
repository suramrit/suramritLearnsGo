package readhtml 

import ("encoding/json";
		"os";
		"io";
		"fmt"
		"golang.org/x/net/html";
		"strings")


type linkElem struct {
	Link string `json:"link"`
	Href string `json:"href"`
	Text string `json:"text"`
}

var r io.Reader

func Convert(r string) []byte{
	res := map[string]interface{}{
		"href": "/wikipedia",
	}
	jso, _:= json.MarshalIndent(res,"", "	")
	os.Stdout.Write(jso)
	return jso
}

func Parse(r io.Reader) ([]linkElem, error){
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	dfsHtml(doc, "")
	links := searchLinks(doc)
	var ret []linkElem
	for _, node := range links {
		ret = append(ret,buildLink(node))
	}
	return ret, nil 
}

func buildLink(n *html.Node) linkElem {
	var ret linkElem
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			ret.Text = nodeText(n)
		}
	}
	return ret
}

func nodeText(n *html.Node) string {
	
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}

	var ret string

	for c := n.FirstChild; c != nil; c = c.NextSibling {
			ret += nodeText(c)
	}
	return strings.Join(strings.Fields(ret), " ")

}

func searchLinks(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, searchLinks(c)...)
	}
	return ret
}

func dfsHtml(n *html.Node, padding string){
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<"+msg+">"
	}
	fmt.Println(padding,msg)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfsHtml(c, padding + "  ")
	}
}