package crwal

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
	"log"
	"regexp"
	"strings"
)

const (
	url = "https://www.youtube.com/watch?v=7i7xflBgS60"
)

func ClashConfig() {
	c := colly.NewCollector()
	rp, err := proxy.RoundRobinProxySwitcher("http://192.168.31.62:7890")
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)

	lzy := c.Clone()

	downLoad := c.Clone()

	downLoad.OnResponse(func(r *colly.Response) {
		err := r.Save("./xxx.zip")
		if err != nil {
			fmt.Printf("download error %v", err)
			return
		}
	})

	downLoad.OnRequest(func(r *colly.Request) {
		fmt.Println("downLoad visiting", r.URL)
	})

	lzy.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	lzy.OnResponse(func(r *colly.Response) {
		fmt.Printf("Response %s: %d bytes\n", r.Request.URL, len(r.Body))
		return
	})

	lzy.OnHTML("*", func(e *colly.HTMLElement) {
		fmt.Printf(e.Text)
		//s := regexpUrl(e.Text, "\"https://develope.lanzoug.com/file/.*?\"")
		//if s != "" {
		//	lzy.Visit(strings.Replace(s, "\"", "", -1))
		//}
	})

	lzy.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error %s: %v\n", r.Request.URL, err)
	})

	c.OnHTML("html body", func(e *colly.HTMLElement) {
		s := regexpUrl(e.Text, "\"https://wwz.lanzouq.com/.*?\"")
		if s != "" {
			lzy.Visit(strings.Replace(s, "\"", "", -1))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Response %s: %d bytes\n", r.Request.URL, len(r.Body))
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error %s: %v\n", r.Request.URL, err)
	})

	c.Visit(url)
}

func regexpUrl(text string, patten string) string {
	compile := regexp.MustCompile(patten)
	subMatch := compile.FindStringSubmatch(text)
	if subMatch != nil {
		return subMatch[0]
	}
	return ""
}
