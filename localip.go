// © 2012 Steve McCoy. Licensed under the MIT license.

/*
The myip command prints all non-loopback IP addresses associated
with the machine that it runs on, one per line.
*/
package main

import (
	"net"
)

func GetLocalIp() ([]string, error) {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}

	for _, addr := range addrs {
		if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			ips = append(ips, ip.IP.String())
		}
	}
	return ips, nil
}

func main() {
	ips, err := GetLocalIp()
	if err != nil {
		println(err)
		return
	}

	for _, ip := range ips {
		println(ip)
	}
}
