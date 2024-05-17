package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var port string

	flag.StringVar(&port, "p", "80", "port range to scan")
	flag.Parse()

	if ip := flag.Arg(0); ip != "" {
		dt := time.Now()
		fmt.Printf("Starting portarum at %s\n", dt.Format(time.UnixDate))
		fmt.Printf("portarum scan report for %s : [%s]\n", ip, port)
		fmt.Println("Port\tState")

		portRange := strings.Split(port, "-")
		startPort, err := strconv.Atoi(portRange[0])
		if err != nil {
			fmt.Printf("error getting start port; %v", err)
		}
		endPort, err := strconv.Atoi(portRange[1])
		if err != nil {
			fmt.Printf("error getting end port; %v", err)
		}

		for p := startPort; p <= endPort; p++ {
			isOpen := scanPort("tcp", "localhost", p)
			if isOpen {
				fmt.Printf("%d\t%s\t\n", p, "open")
			} else {
				fmt.Printf("%d\t%s\t\n", p, "closed")
			}
		}
	} else {
		fmt.Println("error; IP address not provided")
		os.Exit(1)
	}
}

func scanPort(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
