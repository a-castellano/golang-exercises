// Fetch prints the content found at a URL.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	var processed_url string

	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			processed_url = "http://" + url
		} else {
			processed_url = url
		}
		resp, err := http.Get(processed_url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		} else {
			fmt.Println(resp.Status)
		}
	}
}
