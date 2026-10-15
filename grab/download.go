package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cavaliergopher/grab/v3"
)

func main() {
	// Get the args
	if len(os.Args) < 2 {
		fmt.Println("Usage: download <url>")
		os.Exit(1)
	}
	url := os.Args[1]
	// create client
	client := grab.NewClient()
	req, err := grab.NewRequest(".", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating grab request for %s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("Request is %#v\n", req.HTTPRequest)

	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating http request for %s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("httpReq is %#v\n", httpReq)

	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	// check for errors
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	fmt.Printf("Download saved to ./%v \n", resp.Filename)
}
