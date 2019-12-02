package main

import (
	"log"
	"strconv"
)

func main() {

	log.Printf("Starting the connector and reading properties in the properties.ini file")
	/* Reading properties from ./properties.ini */
	prop, _ := ReadPropertiesFile("./properties.ini")
	port, _ := strconv.Atoi(prop["GreenplumPort"])
	batch, _ := strconv.Atoi(prop["Batch"])

	log.Printf("Properties read: Connecting to the Grpc server specified")

	/* Connect to the grpc server specified */
	gpssClient := MakeGpssClient(prop["GpssAddress"], prop["GreenplumAddress"], int32(port), prop["GreenplumUser"], prop["GreenplumPassword"], prop["Database"], prop["SchemaName"], prop["TableName"])
	gpssClient.ConnectToGrpcServer()

	log.Printf("Connected to the grpc server")
	log.Printf("Listening connections to" + prop["SocketAddress"])

	socket := makeSocketClient(prop["SocketAddress"], gpssClient, batch, prop["Delim"])
	socket.socketListen()

}
