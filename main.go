package main

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/qbxt/ping-benson-client/constants"
	"time"
)

func main() {
	requester := gorequest.New().Timeout(2 * time.Second)
	for {
		requester.Post(fmt.Sprintf(constants.URL_TEMPLATE, constants.TOKEN, 69)).End(printStatus)
		fmt.Printf("Sent heartbeat: ID %s, Ping %d\n", constants.TOKEN, 69)
		time.Sleep(20 * time.Second)
	}
}

func printStatus(resp gorequest.Response, body string, errs []error) {
	fmt.Printf("Got response %s\n", resp.Status)
}