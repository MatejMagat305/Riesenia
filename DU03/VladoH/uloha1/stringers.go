package main

import (
"strconv"
"fmt"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() (res string) {
	res = strconv.Itoa(int(ip[0]))
	for ind := 1; ind < len(ip); ind++{
		res += "." + strconv.Itoa(int(ip[ind]))
	}
	return
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
