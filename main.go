package main

import (
	"fmt"
	"github.com/go-ping/ping"
	"github.com/parnurzeal/gorequest"
	"github.com/qbxt/ping-benson-client/constants"
	"time"
)

func main() {
	requester := gorequest.New().Timeout(2 * time.Second)
	pinger, err := ping.NewPinger("www.google.com")
	if err != nil {
		panic(err)
	}
	pinger.SetPrivileged(true)
	for {
		pinger.Count = 1
		if err := pinger.Run(); err != nil {
			panic(err)
		}
		stats := pinger.Statistics()
		requester.Post(fmt.Sprintf(constants.URL_TEMPLATE, constants.TOKEN, stats.AvgRtt / time.Millisecond)).End(printStatus)
		fmt.Printf("Sent heartbeat: ID %s, Ping %d\n", constants.TOKEN, stats.AvgRtt / time.Millisecond)
		time.Sleep(20 * time.Second)
	}
}

func printStatus(resp gorequest.Response, body string, errs []error) {
	fmt.Printf("Got response %s\n", resp.Status)
}