// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		var pref, answer string
		if !strings.HasPrefix(url, "http://") {
			pref = "http://"
			url = pref + url
			fmt.Println("No prefix found, added the http:// prefix ")
			fmt.Println("The new url is ", url)
		}
		fmt.Println("Sending request ...")
		resp, err := http.Get(url)

		if resp.StatusCode == 200 {
			answer = "OK"
		}

		fmt.Println("The response code is : ", resp.StatusCode, answer)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
