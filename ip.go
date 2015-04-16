package main

import (
	"net"
    "net/http"
    "strings"
)

func LocalIp() ([]string, error) {
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

func RemoteIp(r *http.Request) string {
    if r == nil {
        return ""
    }

    realIp := r.Header.Get("X-Forwarded-For")
    if realIp != "" {
        return strings.SplitN(realIp, ",", 2)[0]
    }

    realIp = r.Header.Get("X-Real-IP")
    if realIp != "" {
        return realIp
    }
    return strings.SplitN(r.RemoteAddr, ":", 2)[0]
}

func main() {
	ips, err := LocalIp()
	if err != nil {
		println(err)
		return
	}

	for _, ip := range ips {
		println(ip)
	}
}
