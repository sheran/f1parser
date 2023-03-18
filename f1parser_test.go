package f1parser

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestBareToml(t *testing.T) {
	list := LoadToml("filters")
	url := "https://www.motorsportmagazine.com/articles/single-seaters/f1/mph-dominant-red-bull-could-be-curbed-by-f1-rule-change-should-it/?v=0f177369a3b7"
	bodytoml, err := Process(url, list)
	if err != nil {
		panic(err)
	}
	if bodytoml != nil {
		fmt.Println(bodytoml.Title)
		fmt.Println(bodytoml.Body)
	} else {
		fmt.Println("empty body")
	}
}

func TestHTML(t *testing.T) {
	resp, err := http.Get("https://www.planetf1.com/news/charles-leclerc-grid-penalty-saudi-arabia/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	var pattern strings.Builder
	doc.Find("div.ciam-article-pf1 > p").Each(func(i int, s *goquery.Selection) {
		if s.Children().Length() > 3 {
			for _, child := range s.Children().Nodes {
				pattern.WriteString(child.Data)
			}
		}
	})
	fmt.Println(pattern.String())
}
