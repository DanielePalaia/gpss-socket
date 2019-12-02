# gpss-socket

gpss integration with linux tcp socket.

Small varation of:
https://github.com/DanielePalaia/gpss-pipe

We are now receving info from a socket and push them back to Greenplum using GPSS.

This component is implementing a tcp server. The tcp server wait for new connections (in multithreading) and then wait for data coming from the opened socket. </br>

1) a gpss server needs to be initialized and working </br>

3) Let's create the destination table in greenplum (whatever table is fine if it is coherent with the input fields) ex. </br>

   **CREATE TABLE people(id int, name varchar(1000), surname varchar(1000), email varchar(1000), gender varchar(10));** </br>

3) Configure the program properties file that needs to be in the path you are running the software (./bin/linux/properties.ini), where specify the path of the pipe created and the delim set as input field separator (; in case of csv. **SocketAddress** is the port where the tcp server will listen) </br>

   **GpssAddress=10.91.51.23:50007</br>**
   **GreenplumAddress=10.91.51.23</br>**
   **GreenplumPort=5533</br>**
   **GreenplumUser=gpadmin</br>**
   **GreenplumPasswd=</br>**
   **Database=test</br>**
   **SchemaName=public</br>**
   **TableName=people</br>**
   **SocketAddress=:8080</br>**
   **Delim=;</br>**
   **Batch=5</br>**

4) Run the software (./bin/macosx/gpss-socket or ./bin/linux/gpss-socket) </br>

   **Danieles-MBP:macosx dpalaia$ ./gpss-socket**</br>
   **2019/12/02 14:38:46 Starting the connector and reading properties in the properties.ini file**</br>
   **2019/12/02 14:38:46 Properties read: Connecting to the Grpc server specified**</br>
   **2019/12/02 14:38:46 Connected to the grpc server**</br>
   **2019/12/02 14:38:46 Listening connections to:8080**</br>

5) In order to test the software, run the client binary as well which will send 10 rows like this in the socket (it will create a tcp connection with the server):</br>

   **1;Renaldo;Bulmer;rbulmer0@nymag.com;Male**</br>

    ./client/bin/osx/client</br>
    ./client/bin/linux/client</br>
   
   **Danieles-MBP:osx dpalaia$ ./client**</br>
   **sending line: 1;Renaldo;Bulmer;rbulmer0@nymag.com;Male**</br>
   **sending line: 1;Renaldo;Bulmer;rbulmer0@nymag.com;Male**</br>
   **sending line: 1;Renaldo;Bulmer;rbulmer0@nymag.com;Male**</br>
   **sending line: 1;Renaldo;Bulmer;rbulmer0@nymag.com;Male**</br>
   **sending line: 1;Renaldo;Bulmer;rbulmer0@nymag.com;Male**</br>
   **sending line: 1;Renaldo;Bulmer;rbulmer0@nymag.com;Male**</br>
   **sending line: 1;Renaldo;Bulmer;rbulmer0@nymag.com;Male**</br>
   **sending line: 1;Renaldo;Bulmer;rbulmer0@nymag.com;Male**</br>
   **sending line: 1;Renaldo;Bulmer;rbulmer0@nymag.com;Male**</br>
   **sending line: 1;Renaldo;Bulmer;rbulmer0@nymag.com;Male**</br>
 
 The server is able to handle multiple connections in parallel

6) you should see some logs in the socketgpss screen and the table populated with 10 elements </br>
see the server logs:</br>

   **Connected to: 127.0.0.1:53056**</br>
   **2019/12/02 15:02:03 connecting to a greenplum database**</br>
   **2019/12/02 15:02:03 Beginning to write to greenplum**</br>
   **2019/12/02 15:02:03 table informations**</br>
   **2019/12/02 15:02:03 prepare for writing**</br>
   **Result:  SuccessCount:5**</br>
   **2019/12/02 15:02:04 disconnecting to a greenplum database**</br>
   **2019/12/02 15:02:04 connecting to a greenplum database**</br>
   **2019/12/02 15:02:04 Beginning to write to greenplum**</br>
   **2019/12/02 15:02:04 table informations**</br>
   **2019/12/02 15:02:04 prepare for writing**</br>
   **Result:  SuccessCount:5**</br>
