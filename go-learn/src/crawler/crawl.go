package crawler

import "fmt"

type fakeResult struct {
	body string
	urls []string
}

type fakeFetcher map[string]*fakeResult

func (f *fakeFetcher) Fetcher(url string) (body string, urls []string, err error) {

	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = &fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

func StartCrawl() {

	Crawl("http://golang.org/", 4, fetcher)
}

type Fetcher interface {
	Fetcher(url string) (body string, urls []string, err error)
}

type result struct {
	url, body string
	urls []string
	err error
	depth int
}

func Crawl(url string, depth int, fetcher Fetcher) {

	results := make(chan *result)
	fetched := make(map[string]bool)
	fetch := func(url string, depth int) {
		body, urls, err := fetcher.Fetcher(url)
		results <- &result{url, body, urls, err, depth}
	}

	go fetch(url, depth)

	for fetching := 1; fetching > 0; fetching-- {
		res := <- results
		if res.err != nil {
			fmt.Println(res.err)
			continue
		}
		fmt.Printf("found： %s %q\n", res.url, res.body)

		if res.depth > 0 {
			for _, u := range res.urls{
				if !fetched[u] {
					fetching++
					go fetch(u, res.depth - 1)
					fetched[u] = true
				}
			}

		}
	}
	close(results)
}



