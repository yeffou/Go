package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func fetchURL(ctx context.Context, url string, wg *sync.WaitGroup) {
	defer wg.Done()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Printf("Failed to create request for %s: %v\n", url, err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Printf("Timeout reached for URL: %s\n", url)
		} else {
			fmt.Printf("Failed to fetch URL: %s, error: %v\n", url, err)
		}
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response for URL: %s, error: %v\n", url, err)
		return
	}

	content := string(body)
	if len(content) > 100 {
		content = content[:100] + "..."
	}
	fmt.Printf("Fetched %s, Content: %s\n", url, content)
}

func main() {
	urls := []string{
		"http://example.com",
		"http://httpbin.org/delay/2",
		"http://httpbin.org/delay/5",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetchURL(ctx, url, &wg)
	}
	wg.Wait()
	fmt.Println("All fetch operations completed.")
}
