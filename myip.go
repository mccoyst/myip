// © 2012 Steve McCoy. Licensed under the MIT license.

/*
The myip command prints all non-loopback IP addresses associated
with the machine that it runs on, one per line.
*/
package main

import (
	"net"
	"os"
)

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			os.Stdout.WriteString(ipnet.IP.String() + "\n")
		}
	}
}
