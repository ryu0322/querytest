package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "http://meigenkakugen.net/%E7%AF%84%E9%A6%AC%E5%8B%87%E6%AC%A1%E9%83%8E/"

	doc, err := goquery.NewDocument(url)

	if err != nil {
		fmt.Printf("error URL access")
		return
	}

	doc.Find("div.graybox").Each(func(_ int, sel *goquery.Selection) {
		serifu := sel.Text()
		fmt.Println(serifu)
	})
}
