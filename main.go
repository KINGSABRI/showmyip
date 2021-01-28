package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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

func main() {
	urlsV4 := []string{
		"https://api.ipify.org?format=json",
		"https://api4.my-ip.io/ip.json",
		"https://ip4.seeip.org/json",
		"https://ipinfo.io",
	}

	urlsV6 := []string{
		"https://api64.ipify.org?format=json",
		"https://api6.my-ip.io",
		"https://api6.my-ip.io/ip.json",
		"https://ipinfo.io",
	}

	v4 := flag.Bool("v4", true, "Show my public IPv4.")
	v6 := flag.Bool("v6", false, "Show my public IPv6.")
	all := flag.Bool("a", false, "Show my public IPv4 & IPv6.")
	flag.Parse()

	// if *all || *v4 && *v6 {
	// 	fmt.Println(request(urlsV4))
	// 	fmt.Println(request(urlsV6))
	// } else if *v6 {
	// 	fmt.Println(request(urlsV6))
	// } else if *v4 {
	// 	fmt.Println(request(urlsV4))
	// } else {
	// 	fmt.Println(request(urlsV4))
	// }

	switch {
	case *all || *v4 && *v6:
		fmt.Println(request(urlsV4))
		fmt.Println(request(urlsV6))
	case *v6:
		fmt.Println(request(urlsV6))
	case *v4:
		fmt.Println(request(urlsV4))
	default:
		fmt.Println(request(urlsV4))
	}
}
