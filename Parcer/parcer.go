package Parcer

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func GetWordDescription(word string, siteWithExpl string)string {
	/*Makes GET request to the site with word
	descriptions and calls parcer to parce gotten html*/
	
	request := siteWithExpl + word
	resp, _ := http.Get(request)
	defer resp.Body.Close()
	return GetParcer(resp)
}

func GetParcer(page *http.Response)string {
	/*Parces gotten html page. If gotten parced text 
	is empty it sends a special exception in a string form*/

	doc, err := goquery.NewDocumentFromReader(page.Body)
	if err != nil{
		log.Panic(err)
	}
	parcedHTML := doc.Find("html body #c #container div#article div div p")
	if parcedHTML.Text() != "" { return parcedHTML.Text() }
	return "❌Тлумачення цього слова не знайдене"
}