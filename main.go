package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://github.com/caliwyr/Proxy-List/raw/main/proxy-list/data-with-geolocation.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	var result []any
	err = json.Unmarshal(body, &result)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	// https://proxy-seller.com/tools/proxy-checker/

	for _, v := range result {
		m, ok := v.(map[string]any)
		if !ok {
			continue
		}
		ip, ok := m["ip"].(string)
		if !ok {
			continue
		}
		port, ok := m["port"].(float64)
		if !ok {
			continue
		}
		geo, ok := m["geolocation"].(map[string]any)
		if !ok {
			continue
		}
		countryCode, ok := geo["countryCode"]
		if !ok {
			continue
		}
		if countryCode != "TR" {
			continue
		}
		fmt.Printf("%v:%v\n", ip, port)
	}
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
