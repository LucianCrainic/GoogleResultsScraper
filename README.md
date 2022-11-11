# GoogleResultScraper
A simple web scraper made with GOLang and [GoQuery](https://github.com/PuerkitoBio/goquery). 
> Goal of the scraper is to scrape the pages generated from a google search, user can select the search term and amount of pages to be displayed
## Instructions
Simple instructions on how to use the web scraper.

Clone the repository on your local computer.
```
git clone https://github.com/LucianCrainic/GoogleResultsScraper.git
```
open the directory and open the main.go file, look for the main function (bottom of the file).
Inside the main function we call the GoogleScrape function wich takes the following parameters :
- searchTerm => term that you want to be sent to google search. 
  - Example: Tom Cruise
- countryCode => country code for websites. 
  - Example: com
- languageCode => language code for websites. 
  - Example: en
- proxyString => proxy if you want to use one, not necessary. 
  - Example: nil
- pages => number of google pages that you want to scrape, the first page only or more. 
  - Example : 1
- count : => number of responses from the scraper. 
  - Example: 10
- backoff => Sleep time for each request. 
  - Example: 10

```go
func main() {
	res, err := GoogleScrape("Tom Cruise", "com", "en", nil, 1, 10, 10)
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}
}
```

Sample output after you run the main.go file, it may take a few seconds for it to return.
```
{1 https://en.wikipedia.org/wiki/Tom_Cruise  }
{2 https://en.wikipedia.org/wiki/Tom_Cruise_filmography  }
{3 https://it.wikipedia.org/wiki/Tom_Cruise  }
{4 https://www.imdb.com/name/nm0000129/  }
{5 http://www.tomcruise.com/  }
{6 https://www.instagram.com/tomcruise/  }
{7 https://twitter.com/tomcruise  }
{8 https://www.facebook.com/officialtomcruise/  }
{9 https://www.mymovies.it/persone/tom-cruise/9665/  }
{10 https://www.themoviedb.org/person/500-tom-cruise  }
```
