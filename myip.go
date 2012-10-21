// © 2012 Steve McCoy. Licensed under the MIT license.

/*
The myip command prints all non-loopback IP addresses associated
with the machine that it runs on.
*/
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Oops: %v\n", err)
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			fmt.Println(ipnet.IP)
		}
	}
}
