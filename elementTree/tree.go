package main
//Working on https://gophercises.com/exercises/ ---- check out for small projects on golang! 

//Disclaimer: The code was developed while following execises at https://gophercises.com/exercises/ and were solely developed from an educational perspective. 


import ("fmt";
		"flag";
		"io";
		"net/http";
		"net/url";
		"../readhtml";
		"strings";
		"encoding/xml";
		"os")

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {

	Xmlns string `xml:"xmlns,attr"`
	Urls []loc `xml:"url"`
}

func main(){
	urlFlag := flag.String("url","https://google.com","url to build sitemap for")
	maxDepth := flag.Int("depth", 3, "max depth that the builder recirsively goes into")
	flag.Parse()
	// GET the webpage
	// build proper url links

	pages := bfs(*urlFlag, *maxDepth)

	//fmt.Println(*urlFlag,"has ", len(pages),"links in it")
	
	toXml := urlset{
		Xmlns : xmlns,
	}

	for _, page := range pages {
		toXml.Urls = append(toXml.Urls, loc{page})
	}
	// filter out links
	// find all pages (BFS)

	// print out XML
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("  ", "    ")
	fmt.Println(xml.Header)
	if err := enc.Encode(toXml); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func bfs(urlString string, maxDepth int) []string{
	seen := make (map[string] struct{}) //struct uses less memory than bool--- preferred method of creating a set
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlString : struct{}{},
	}
	for i :=0 ; i< maxDepth ; i++ {
		q, nq = nq, make(map [string]struct{})
		for url, _ := range q {
			if _, ok := seen[url]; ok {
				continue
			} 
			seen[url] = struct{}{}
			for _, link :=  range getPages(url){
				nq[link] = struct{}{}
			}
		} 
	}
	var ret []string
	for url, _ := range seen {
		ret = append(ret, url)
	}
	return ret
}


func getPages(urlString string) []string {
	resp, err := http.Get(urlString)
	if err != nil {
		panic(err)
	}	
	defer resp.Body.Close()

	// Parse all links on the page
	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme : reqUrl.Scheme,
		Host: reqUrl.Host, 
	}
	base := baseUrl.String()
	//fmt.Println(base)
	return filterPages(base, buildHrefs(resp.Body, base))
}

func filterPages(base string, links []string) []string {
	var ret []string
	for _, link := range links {
		if strings.HasPrefix(link,base){
			ret = append(ret, link)
		}
	}
	return ret
}


func buildHrefs(r io.Reader, base string) []string{

	links, _ := readhtml.Parse(r)

	var hrefs []string

	for _,link := range links {
		switch {
			case strings.HasPrefix(link.Href,"/"):
				hrefs = append(hrefs, base+link.Href)
			case strings.HasPrefix(link.Href,"http"):
				hrefs = append(hrefs, link.Href)
		}
	}

	return hrefs
}





