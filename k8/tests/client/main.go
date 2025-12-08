package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
)

type Res struct {
	Version int    `json:"version"`
	Id      string `json:"id"`
}

func main() {
	ids := make([]string, 0)
	for {
		xd, err := http.Get("http://192.168.49.2:32642/test")
		if err != nil {
			panic(err)
		}
		var pog Res
		if err := json.NewDecoder(xd.Body).Decode(&pog); err == nil {
			if slices.Contains(ids, pog.Id) {
				continue
			}
			ids = append(ids, pog.Id)
			if len(ids) > 1 {
				fmt.Printf("%v", ids)
			}
		}

	}

}
