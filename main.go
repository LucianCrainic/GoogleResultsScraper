package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"math/rand"
	"net/url"
	"github.com/PuerkitoBio/goquery"
) 

var googleDomains = map[string]string{
	// empty for now
}

type SearchResult struct{
	ResultRank int
	ResultURL string
	ResultTitle string
	ResultDesc string
}

var userAgents = []string{
	// empty for now
}
//* randomUserAgent func
func randomUserAgent() string {
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}
//* buildGoogleUrls func 
func buildGoogleUrls(searchTerm, countryCode, languageCode string, pages, count int)([]string, error){
	toScrape := []string{}
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	if googleBase, found := googleDomains[countryCode]; found{
		for i := 0; i<pages ; i++{
			start := i*count
			scrapeURL := fmt.Sprints("%s%s&num=%d&hl=%s&start=%d&filter=0",googlegoogleBase, searchTerm, count, languageCode, start)
		}
	} else {
		err := fmt.Errorf("country (%s) is currently not supported", countryCode)
		return nil, err
	} 
	return toScrape, nil
}
//! GoogleScrape func
func GoogleScrape(searchTerm, countryCode, languageCode string, pages, count)([]SearchResult, err){
	results := []SearchResult{}
	resultCounter := 0
	googlePages, err := buildGoogleUrls(searchTerm, countryCode,languageCode, pages, count)
	if err !=nil {
		return nil, err
	}

	for _, page := range googlePages{
		res, err := scrapeClientRequest(page, proxyString)
	}
}
func main() {
	res, err := GoogleScrape("Sapienza", "com", "en", 1, 30)
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}
}