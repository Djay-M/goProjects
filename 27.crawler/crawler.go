package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

const links = "https://www.amazon.in/s?bbn=1389401031&rh=n%3A1389401031%2Cp_36%3A1318507031&dc&qid=1729450514&rnid=1318502031&ref=lp_1389401031_nr_p_36_4"

//"https://www.flipkart.com/search?q=mobiles&as=on&as-show=on&otracker=AS_Query_TrendingAutoSuggest_1_0_na_na_na&otracker1=AS_Query_TrendingAutoSuggest_1_0_na_na_na&as-pos=1&as-type=TRENDING&suggestionId=mobiles&requestId=f1ff39f0-1278-4f3c-9787-79c589208080"

// type productData struct {
// 	OriginalLink *url.URL `json:"originalLink"`
// 	Links []string `json:"links"`
// }

func handleErrors(funcName string, err error) {
	if err != nil {
		log.Fatal("Error in: ", funcName, "\n ERROR :::", err)
	}
}

func main() {
	fmt.Println("Starting the Web Crawler")
	productLinks := make(map[string][]string)
	var linkData []string
	c := colly.NewCollector()

	// called before an HTTP request is triggered
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	// triggered when the scraper encounters an error
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong: ", err)
	})

	// fired when the server responds
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	c.OnHTML("a", func(h *colly.HTMLElement) {
		hrefLink := h.Attr("href")

		if len(hrefLink) > 0 {
			linkData = append(linkData, hrefLink)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		productLinks[r.Request.URL.String()] = linkData
	})

	c.Visit(links)

	// fmt.Println("%+v", productLinks)

	jsonData, err := json.Marshal(productLinks)
	handleErrors("Marshal", err)
	os.WriteFile("productLinks.json", jsonData, 0644)
}
