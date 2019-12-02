package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

type pipeClient struct {
	connString    string
	socketAddress string
	delim         string
	batch         int
	buffer        []string
	gpssclient    *gpssClient
}

func makeSocketClient(socketAddress string, gpssclient *gpssClient, batch int, delim string) *pipeClient {
	client := new(pipeClient)
	client.socketAddress = socketAddress
	client.batch = batch
	client.delim = delim
	client.buffer = make([]string, client.batch)
	client.gpssclient = gpssclient

	return client
}

func (client *pipeClient) failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (client *pipeClient) socketListen() {
	// Open a socket for reading

	listen := client.socketAddress
	addr, error := net.ResolveTCPAddr("tcp", listen)
	if error != nil {
		fmt.Printf("Cannot parse \"%s\": %s\n", listen, error)
		os.Exit(1)
	}
	listener, error := net.ListenTCP("tcp", addr)
	if error != nil {
		fmt.Printf("Cannot listen: %s\n", error)
		os.Exit(1)
	}
	for { // ever...
		conn, error := listener.AcceptTCP()
		if error != nil {
			fmt.Printf("Cannot accept: %s\n", error)
			os.Exit(1)
		}
		go client.handle(conn)
	}
}

func (client *pipeClient) handle(conn net.Conn) {
	defer conn.Close()
	defer fmt.Println("")

	fmt.Printf("Connected to: %s\n", conn.RemoteAddr().String())

	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("Error reading:")
			fmt.Println(err)
			continue
		}

		buff := bytes.NewBuffer(buf)
		client.scanBuffer(buff)
	}
}

func (client *pipeClient) scanBuffer(buff *bytes.Buffer) {

	i := 0

	for true {
		line, err := buff.ReadBytes('\n')
		// File terminated, let's write what we have
		if err != nil {
			if i > 0 {
				client.delegateToGpss(client.buffer)
			}
			break
		}
		//log.Printf("line: " + string(line))
		client.buffer[i] = string(line)
		i++

		if i >= client.batch {
			client.delegateToGpss(client.buffer)
			i = 0

		}
	}

}

func (client *pipeClient) delegateToGpss(buffer []string) {

	client.gpssclient.ConnectToGreenplumDatabase()
	client.gpssclient.WriteToGreenplum(client.buffer, client.delim)
	client.gpssclient.DisconnectToGreenplumDatabase()

}
