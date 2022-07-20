package crwal

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

const (
	url = "https://www.youtube.com/watch?v=7i7xflBgS60&t=322s"
)

func ClashConfig() {
	c := colly.NewCollector()
	//lanzouq := c.Clone()
	c.OnHTML("#description > yt-formatted-string", func(element *colly.HTMLElement) {
		element.ForEach("a:nth-child", func(i int, element *colly.HTMLElement) {
			text := element.Text
			fmt.Println(text)
			if strings.Contains(text, "https://wwz.lanzouq.com/") {
				fmt.Println(element.Attr("href"))
				//lanzouq.Visit(element.Attr("href"))
				return
			}
		})
	})
	c.Visit(url)
}
