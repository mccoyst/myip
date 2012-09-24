// © 2012 Steve McCoy. Licensed under the MIT license.
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Oops: %v\n", err)
		os.Exit(1)
	}

	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Oops: %v\n", err)
		os.Exit(1)
	}

	for _, a := range addrs {
		fmt.Println(a)
	}
}
