// https://tour.golang.org/methods/18

package main

import "fmt"

type IPAddr [4]byte

func (ipadr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipadr[0], ipadr[1], ipadr[2], ipadr[3])
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

