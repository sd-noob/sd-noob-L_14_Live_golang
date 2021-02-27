package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	nl, err := net.Listen("tcp", ":8080") //ip:port
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) // 1 = error type which means terminated
	}

	conn, err := nl.Accept() //layer 5 session layer
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) // 1 = error type which means terminated
	}

	remortAdder := conn.RemoteAddr().String()
	fmt.Println(remortAdder)
	conn.Write([]byte("Welcome we have recieved your sms"))
	conn.Close()
	nl.Close()

}
