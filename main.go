package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/gocolly/colly"
)

type Vacancy struct {
	Name        string
	Price       string
	Link        string
	Description string
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	var vacancy Vacancy
	searchText := ""
	fmt.Printf("Вакансия: ", searchText)
	fmt.Scanf("%s", &searchText)
	HHcollector := colly.NewCollector(
		colly.AllowedDomains("hh.ru", "khimki.hh.ru"),
	)

	// Saving each vacancy into an object
	HHcollector.OnHTML(".vacancy-serp-item", func(element *colly.HTMLElement) {
		vacancy.Link = strings.Join(element.ChildAttrs("a.bloko-link", "href"), "")
		vacancy.Name = element.ChildText("span.g-user-content")
		vacancy.Price = element.ChildText(".vacancy-serp-item__sidebar")
		fmt.Println(vacancy.Name)
	})

	// Visiting next page after scrapping current
	HHcollector.OnHTML("span.bloko-button_pressed + span > a[href]", func(element *colly.HTMLElement) {
		URL := element.Attr("href")
		element.Request.Visit(URL)
	})

	HHcollector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
	})

	HHcollector.OnError(func(_ *colly.Response, err error) {
		fmt.Println(err)
	})

	HHcollector.Visit("https://hh.ru/search/vacancy?area=1&text=" + searchText)
}
