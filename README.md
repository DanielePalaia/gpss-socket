# gpss-socket

gpss integration with tcp socket. </br>
In this scenario sockets are of type tcp but we can have with the same principle upd and unix sockets. </br>

Small varation of:
https://github.com/DanielePalaia/gpss-pipe

We are now receving info from a socket and push them back to Greenplum using GPSS.

This component is implementing a tcp server. The tcp server wait for new connections (in multithreading) and then wait for data coming from the opened socket. </br>

1. a gpss server needs to be initialized and working </br>

2. Let's create the destination table in greenplum (whatever table is fine if it is coherent with the input fields) ex. </br>
```
   CREATE TABLE people(id int, name varchar(1000), surname varchar(1000), email varchar(1000), gender varchar(10));
```
3. Configure the program properties file that needs to be in the path you are running the software (./bin/linux/properties.ini), where specify the port socket to listen and the delim set as input field separator (; in case of csv. **SocketAddress** is the port where the tcp server will listen) </br>

```
   GpssAddress=10.91.51.23:50007
   GreenplumAddress=10.91.51.23
   GreenplumPort=5533
   GreenplumUser=gpadmin
   GreenplumPasswd=
   Database=test
   SchemaName=public
   TableName=people
   SocketAddress=:8080
   Delim=;
   Batch=5
```   

4. Run the software (./bin/macosx/gpss-socket or ./bin/linux/gpss-socket) </br>

```
   Danieles-MBP:macosx dpalaia$ ./gpss-socket
   2019/12/02 14:38:46 Starting the connector and reading properties in the properties.ini file
   2019/12/02 14:38:46 Properties read: Connecting to the Grpc server specified
   2019/12/02 14:38:46 Connected to the grpc server
   2019/12/02 14:38:46 Listening connections to:8080
```

5. In order to test the software, run the client binary as well which will connect to the server and send 10 rows like this in the socket (it will create a tcp connection with the server):</br>

   **1;Renaldo;Bulmer;rbulmer0@nymag.com;Male**</br>

    ./client/bin/osx/client</br>
    ./client/bin/linux/client</br>

```
   Danieles-MBP:osx dpalaia$ ./client
   sending line: 1;Renaldo;Bulmer;rbulmer0@nymag.com;Male
   sending line: 1;Renaldo;Bulmer;rbulmer0@nymag.com;Male
   sending line: 1;Renaldo;Bulmer;rbulmer0@nymag.com;Male
   sending line: 1;Renaldo;Bulmer;rbulmer0@nymag.com;Male
   sending line: 1;Renaldo;Bulmer;rbulmer0@nymag.com;Male
   sending line: 1;Renaldo;Bulmer;rbulmer0@nymag.com;Male
   sending line: 1;Renaldo;Bulmer;rbulmer0@nymag.com;Male
   sending line: 1;Renaldo;Bulmer;rbulmer0@nymag.com;Male
   sending line: 1;Renaldo;Bulmer;rbulmer0@nymag.com;Male
   sending line: 1;Renaldo;Bulmer;rbulmer0@nymag.com;Male

```
 The server is able to handle multiple connections in parallel

6. you should see some logs in the socketgpss screen and the table populated with 10 elements </br>
see the server logs:</br>

```
   Connected to: 127.0.0.1:53056
   2019/12/02 15:02:03 connecting to a greenplum database
   2019/12/02 15:02:03 Beginning to write to greenplum
   2019/12/02 15:02:03 table informations
   2019/12/02 15:02:03 prepare for writing
   Result:  SuccessCount:5
   2019/12/02 15:02:04 disconnecting to a greenplum database
   2019/12/02 15:02:04 connecting to a greenplum database
   2019/12/02 15:02:04 Beginning to write to greenplum
   2019/12/02 15:02:04 table informations
   2019/12/02 15:02:04 prepare for writing
   Result:  SuccessCount:5
```
