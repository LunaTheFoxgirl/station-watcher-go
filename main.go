package main

import (
	"os"
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"bytes"
	"log"
	"io"
)

const usage string = "<adapter> <watch uri> <report uri>"

var (
	Info    *log.Logger
	Error   *log.Logger
)


func main() {
	initLogger(os.Stdout, os.Stderr)

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
	var previousResult []byte = nil

	for {
		resp, err := http.Get(watchUri)
		if err != nil {
			Error.Println("Connnection error for uri", "<"+watchUri+">", err.Error())
			time.Sleep(15 * time.Second)
			continue
		}

		if resp.StatusCode == http.StatusOK {
			newResult, _ := ioutil.ReadAll(resp.Body);

			if previousResult == nil {
				triggerEvent(reportUri, newResult)
			} else {
				responsesMatch := adapter.Compare(previousResult, newResult)
				if !responsesMatch { triggerEvent(reportUri, newResult) }
			}

			previousResult = newResult
		} else {
			Error.Println("Request url", "<"+watchUri+">", "returned", resp.Status)
		}

		time.Sleep(2 * time.Second)
	}
}

func initLogger(infoHandle, errorHandle io.Writer) {
	Info = log.New(infoHandle, "",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"",
		log.Ldate|log.Ltime|log.Lshortfile)
}


func triggerEvent(reportUri string, response []byte) {

	Info.Println("Change in content: triggering web hook")

	responseReader := ioutil.NopCloser(bytes.NewBuffer(response))
	resp, err := http.Post(reportUri, "text/plain", responseReader)
	if err != nil || resp.StatusCode != http.StatusOK {
		if err != nil {
			Error.Println("Webhook Response Error:", err.Error())
			time.Sleep(15 * time.Second)
			return
		}
		Error.Println("Webhook returned response", resp.Status, "verify API key.")
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