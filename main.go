package main

import (
	"os"
	"fmt"
	"net/http"
	"time"
)

const usage string = "<adapter> <watch uri> <report ui>"

func main() {
	//Fetch start arguments.
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("Not enough arguments!\nusage:", usage)
		return
	}	

	adapterName := args[0]
	watchUri := args[1]
	reportUri := args[2]

	//List of all available adapters.
	adapters := map[string]Adapter{
		"shoutcast2": ShoutcastAdapter{},
		"icecast": IcecastAdapter{},
	}

	//Make sure that the adapter exists.
	if !adaptersContains(adapters, adapterName) {
		fmt.Println("Adapter", adapterName, "not found!")
		return
	}

	adapter := adapters[adapterName]
	var previousResult *http.Response = nil
	for {
		resp, err := http.Get(watchUri)
		if err != nil {
			fmt.Println("Connnection error for uri", "<"+watchUri+">", err.Error())
			time.Sleep(15 * time.Second)
			continue
		}

		if resp.StatusCode == http.StatusOK {
			if previousResult == nil {
				triggerEvent(reportUri, resp)
			} else {
				pr := make([]byte, previousResult.ContentLength)
				r := make([]byte, resp.ContentLength)
				previousResult.Body.Read(pr)
				resp.Body.Read(r)
				changed := adapter.Compare(pr, r)
				if changed { triggerEvent(reportUri, resp) }
			}

			previousResult = resp
		} else {
			fmt.Println("Request url", "<"+watchUri+">", "returned", resp.Status)
		}

		time.Sleep(2 * time.Second)
	}
}

func triggerEvent(reportUri string, response *http.Response) {
	resp, err := http.Post(reportUri, "text/plain", response.Body)
	if err != nil || resp.StatusCode != http.StatusOK {
		if err != nil {
			fmt.Println("Error:", err.Error())
			time.Sleep(15 * time.Second)
			return
		}
		fmt.Println("Webhook returned response", resp.Status, "verify API key.")
		time.Sleep(15 * time.Second)
	}
	return
}

func adaptersContains(adapterList map[string]Adapter, name string) bool {
	for k, _ := range adapterList {
		if k == name { return true }
	}
	return false
}