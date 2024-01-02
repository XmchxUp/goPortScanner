package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./program <hostname>")
		os.Exit(0)
	}

	hostname := os.Args[1]
	ips, err := net.LookupIP(hostname)
	if err != nil {
		fmt.Printf("DNS lookup failed for %s: %v\n", hostname, err)
	}

	fmt.Printf("Scanning target: %s\n", hostname)
	fmt.Printf("Scanning started at: %s\n", time.Now().String())

	var wg sync.WaitGroup

	for _, ip := range ips {
		fmt.Printf("Scanning IP: %s\n", ip)
		for i := 1; i <= 65535; i++ {
			wg.Add(1)
			go func(ip string, port int) {
				defer wg.Done()
				conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), time.Second)
				if err == nil {
					fmt.Printf("Port: %d is open.\n", port)
					conn.Close()
				}
			}(ip.String(), i)
		}
	}
	wg.Wait()

}
