package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var (
	clients = []net.Conn{}
)

func main() {

	l, err := net.Listen("tcp", "127.0.0.1:6000")

	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	defer l.Close()

	go listen(l)

	send()
}

func listen(l net.Listener) {
	for {
		c, err := l.Accept()
		fmt.Println("Connect: " + c.RemoteAddr().String())

		if err != nil {
			fmt.Println("Error: ", err.Error())
			break
		}

		clients = append(clients, c)
		go receive(c)
	}
}

func send() {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadBytes('\n')

		for i := 0; i < len(clients); i++ {
			_, err := clients[i].Write(text)

			if err != nil {
				fmt.Println("Error: ", err.Error())
				break
			}

			fmt.Println("--> ", string(text[:len(text)-1]))
		}
	}
}

func receive(c net.Conn) {
	for {
		reader, err := bufio.NewReader(c).ReadBytes('\n')

		if err != nil {
			fmt.Println("Error: " + err.Error())
			break
		}

		//c.Read(reader)

		str := string(reader[:len(reader)-1])
		fmt.Println("<-- ", str)

		/*
			c.Write(reader)
			fmt.Println("--> ", str)
		*/
	}
}
