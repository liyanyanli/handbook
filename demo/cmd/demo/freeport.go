package main

import (
	"strconv"
	"net"
	"fmt"
)


// Get a free port.
func Get() (port int, err error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()

	addr := listener.Addr().String()
	_, portString, err := net.SplitHostPort(addr)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(portString)
}
func main() {
	port, err := Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(strconv.Itoa(port))
}