package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"os"
)

type Fact struct {
	Quotes string `json:"quotes"`
}

func main() {
	allQuotes := make([]Fact, 0)

	collector := colly.NewCollector()

	collector.OnHTML("body > div.globalwrapper > div.page_wrapper > div > section.row > article > section > section > section.body > span > p ", func(element *colly.HTMLElement) {

		quotesList := element.Text
		fact := Fact{
			Quotes: quotesList,
		}
		allQuotes = append(allQuotes, fact)
	})
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting: ", request.URL.String())
	})

	err := collector.Visit("https://parade.com/937586/parade/life-quotes/")
	if err != nil {
		return
	}

	encode := json.NewEncoder(os.Stdout)
	encode.SetIndent("", " ")
	err = encode.Encode(allQuotes)
	if err != nil {
		return
	}
}
