package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

var file = flag.String("file", "", "path to file to read (newline delimited)")
var debug = flag.Bool("debug", false, "")

func main() {
	flag.Parse()

	f, err := ioutil.ReadFile(*file)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	urls := strings.Split(string(f), "\n")
	var validProtocol = regexp.MustCompile(`^http`)

	for _, url := range urls {
		if len(url) > 0 {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				if !validProtocol.MatchString(url) {
					if *debug {
						fmt.Printf("URL %s does not have protocol, using http\n", url)
					}
					url = "http://" + url
				}
				resp, err := http.Get(url)
				if err != nil {
					fmt.Printf("URL %s failed with:\n%s", url, err)
				} else {
					if *debug {
						fmt.Printf("%+v\n", resp)
					}
					if resp.StatusCode >= 300 {
						fmt.Printf("%s returned with %d\n", url, resp.StatusCode)
					}
				}
			}(url)
		}
	}
	wg.Wait()
}
