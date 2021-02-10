package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

var count = 0

func handleConnection(c net.Conn, count int) {
	//fmt.Print("client. ")
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}
		counter := strconv.Itoa(count)
		fmt.Println("client= ", counter, ": ", temp)

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("server= ", counter, ": ")
		text, _ := reader.ReadString('\n')
		//user input is sent to the TCP server over the network
		fmt.Fprintf(c, text+"\n")
		//
		// c.Write([]byte(string(counter)))
	}
	c.Close()
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		count++
		go handleConnection(c, count)

	}
}