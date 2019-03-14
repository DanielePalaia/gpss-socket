package main

import (
	"bytes"
	"io"
	"log"
	"os"
)

type pipeClient struct {
	connString string
	pipePath   string
	delim      string
	batch      int
	buffer     []string
	gpssclient *gpssClient
}

func makePipeClient(pipePath string, gpssclient *gpssClient, delim string) *pipeClient {
	client := new(pipeClient)
	client.pipePath = pipePath
	client.batch = 100
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

func (client *pipeClient) readPipe() {
	// Open named pipe for reading
	log.Println("Opening named pipe: " + client.pipePath + " for reading")
	for true {
		log.Print("waiting for someone to write something in the pipe")

		stdout, _ := os.OpenFile(client.pipePath, os.O_RDONLY, 0600)
		var buff bytes.Buffer
		io.Copy(&buff, stdout)
		stdout.Close()

		log.Printf("pipe written: decomposing the message and sending to gpss in batches: ")
		client.scanBuffer(buff)
	}
}

func (client *pipeClient) scanBuffer(buff bytes.Buffer) {

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
