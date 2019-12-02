package main

import "net"
import "fmt"

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	//for {
	// read in input from stdin
	//reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 10; i++ {
		fmt.Print("sending line: " + "1;Renaldo;Bulmer;rbulmer0@nymag.com;Male\n")
		//text, _ := reader.ReadString('\n')
		text := "1;Renaldo;Bulmer;rbulmer0@nymag.com;Male"
		// send to socket
		fmt.Fprintf(conn, text)
		// listen for reply
		//message, _ := bufio.NewReader(conn).ReadString('\n')
	}
	//fmt.Print("Message from server: " + message)
	//}
}
