package main

import (
	"flag"
	"fmt"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "URL for Sitemap")
	flag.Parse()

	fmt.Println(*urlFlag)
}
