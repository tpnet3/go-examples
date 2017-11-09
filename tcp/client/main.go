package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	//Connect TCP
	conn, err := net.Dial("tcp", "127.0.0.1:6000")

	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	defer func() {
		conn.Close()
		fmt.Println("Close connection.")
	}()

	fmt.Println("Connect: " + conn.LocalAddr().String())

	go receive(conn)
	send(conn)
}

func send(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadBytes('\n')

		_, err := conn.Write(text)

		if err != nil {
			fmt.Println("Error: ", err.Error())
			break
		}

		fmt.Println("<-- ", string(text[:len(text)-1]))
	}
}

func receive(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		text, err := reader.ReadBytes('\n')

		if err != nil {
			fmt.Println("Error: ", err.Error())
			conn.Close()
			//panic(err)
			break
		}

		//conn.Read(text)
		fmt.Println("--> ", string(text[:len(text)-1]))
	}
}
