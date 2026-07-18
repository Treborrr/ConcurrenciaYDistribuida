package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var remotehost string

func main() {
	gin := bufio.NewReader(os.Stdin)
	fmt.Print("Remote host: ")
	remotehost, _ = gin.ReadString('\n')
	remotehost = strings.TrimSpace(remotehost)
	send(1000)
}

func send(num int) {
	conn, _ := net.Dial("tcp", remotehost)
	defer conn.Close()
	fmt.Fprintf(conn, "%d\n", num)
}
