package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {

	// calling http.Get() which gives a response and an error value.
	resp, err := http.Get(url)

	//error handling
	if err != nil {
		fmt.Printf("%s is Down!\n", url)
	} else {
		// a good practice is to close the Response Body if there is one
		// If you forget to close it there will be a resource leak.
		defer resp.Body.Close()
		fmt.Printf("%s -> Status Code: %d \n", url, resp.StatusCode)

		// fetching the page if HTTP Status Code is 200 (success)
		if resp.StatusCode == 200 {
			// The resp.Body implements the io.Reader interface
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			//Creating the file's name
			file := strings.Split(url, "//")[1]
			file += ".txt"
			fmt.Printf("Writing response Body to %s\n", file)

			// Writing the response Body to File
			// If the file doesn't exist WriteFile() creates it and if it already exists
			// the function will truncate it before writing to file.
			err = ioutil.WriteFile(file, bodyBytes, 0664)

			if err != nil {
				log.Fatal(err)
			}

		}
	}

	wg.Done()
}

func main() {
	urls := []string{"https://www.golang.org", "https://www.google.com", "https://www.medium.com", "https://www.zivarcreations.com"}

	var wg sync.WaitGroup

	wg.Add(len(urls))

	for _, url := range urls {
		go checkAndSaveBody(url, &wg)
		fmt.Println(strings.Repeat("#", 20))
	}
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())
	wg.Wait()
}
