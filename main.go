package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
)

var urlsV4 = []string{
	"https://api.ipify.org?format=json",
	"https://api4.my-ip.io/ip.json",
	"https://ip4.seeip.org/json",
	"https://ipinfo.io",
}

var urlsV6 = []string{
	"https://api64.ipify.org?format=json",
	"https://api6.my-ip.io",
	"https://api6.my-ip.io/ip.json",
	"https://ipinfo.io",
}

func logToFile(file string, message string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	ip := "| " + message
	logger := log.New(f, "", log.LstdFlags)
	logger.Println(ip)
}

func notify(title string, message string, logFile string) {

	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		panic(err)
	}

	err = beeep.Notify(title, message, "")
	if err != nil {
		panic(err)
	}

	if logFile != "" {
		logToFile(logFile, message)
	}

}

func request(urls []string) string {

	var rjson map[string]interface{}
	var ip string

	for _, url := range urls {
		res, err := http.Get(url)
		if err != nil {
			continue
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			continue
		}

		json.Unmarshal([]byte(body), &rjson)
		ip = fmt.Sprintf("%v", rjson["ip"])

		if len(strings.TrimSpace(ip)) != 0 {
			break
		}

		defer res.Body.Close()
	}
	return ip
}

func doPrint(toPrint string) {
	if toPrint == "all" {
		fmt.Println(request(urlsV4))
		fmt.Println(request(urlsV6))
	} else if toPrint == "v4" {
		fmt.Println(request(urlsV4))
	} else if toPrint == "v6" {
		fmt.Println(request(urlsV6))
	} else {
		fmt.Println(request(urlsV4))
	}
}

func doNotify(toNotify string, log string, delay int) {
	for {
		time.Sleep(time.Minute * time.Duration(delay))

		if toNotify == "all" {
			notify("New IPv4:", request(urlsV4), log)
			notify("New IPv6:", request(urlsV6), log)
		} else if toNotify == "v4" {
			notify("New IPv4:", request(urlsV4), log)
		} else if toNotify == "v6" {
			notify("New IPv6:", request(urlsV6), log)
		} else {
			notify("New IPv4:", request(urlsV4), log)
		}

		if delay < 5 {
			fmt.Println("Sorry, I cannot use short delay during the APIs request limitation.")
			fmt.Println("Recommended to use +5 minutes delay.")
			break
		}
	}
}

func main() {
	v4 := flag.Bool("4", false, "Show my public IPv4.")
	v6 := flag.Bool("6", false, "Show my public IPv6.")
	all := flag.Bool("a", false, "Show my public IPv4 & IPv6.")
	ntfy := flag.Bool("n", false, "Show desktop notifications. (Continues check every 10 min)")
	log := flag.String("l", "", "Log new IP address (use with '-n' flag).")
	delay := flag.Int("d", 10, "Run in loop to notify every X minutes.")
	flag.Parse()

	switch {
	case *ntfy:
		if *all || *v4 && *v6 {
			doNotify("all", *log, *delay)
		} else if *v4 {
			doNotify("v4", *log, *delay)
		} else if *v6 {
			doNotify("v6", *log, *delay)
		} else {
			doNotify("v4", *log, *delay)
		}
	case *all || *v4 && *v6:
		doPrint("all")
	case *v4:
		doPrint("v4")
	case *v6:
		doPrint("v6")
	default:
		doPrint("v4")
	}
}
