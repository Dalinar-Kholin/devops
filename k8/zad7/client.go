package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	url := "http://192.168.49.2:31459/?i=35"

	i := 0

	for {
		client := &http.Client{
			Timeout: 40 * time.Millisecond,
		}

		resp, err := client.Get(url)
		if err != nil {
			fmt.Println("request error:", err)
			continue
		}

		fmt.Printf("status: %d, %d\n", resp.StatusCode, i)
		i++
	}
}
