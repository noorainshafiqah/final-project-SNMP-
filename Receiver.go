package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", "192.168.235.128:8081")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\n\t =================================================")
	fmt.Print("\n\t #\t\t SNMP Trap Monitor \t\t#")
	fmt.Print("\n\t =================================================\n")

	fmt.Print("\n\t ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Print("\n\t |  --> t,	Total available RAM in kilobytes  |")
	fmt.Print("\n\t |  --> u,	Used available RAM in kilobytes   |")
	fmt.Print("\n\t |  --> f,	Free available RAM in kilobytes   |")
	fmt.Print("\n\t ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")

	for {
		fmt.Print("\n Choose your command: ")
		text, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, text + "\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("\nResults from server: \n\t" + message + "\n")
	}

}

