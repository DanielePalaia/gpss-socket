# gpss-pipe
gpss integration with pipes

This component is listening forever to a pipe (when EOF is reached it ask to gpss to write the info). </br>

1) a gpss server needs to be initialized and working </br>

2) Let's create a named pipe in Linux: </br>
mkfifo mypipe <br>

3) Let's create the destination table in greenplum (whatever table is fine if it is coherent with the input fields) ex. </br>

CREATE TABLE people(id int, name varchar(1000), surname varchar(1000), email varchar(1000), gender varchar(10)); </br>

4) Configure the program properties (./bin/linux/properties.ini), where put the path of the pipe created and the delim set as input field separator (; in case of csv) </br>

**GpssAddress=10.91.51.23:50007</br>**
**GreenplumAddress=10.91.51.23</br>**
**GreenplumPort=5533</br>**
**GreenplumUser=gpadmin</br>**
**GreenplumPasswd=</br>**
**Database=test</br>**
**SchemaName=public</br>**
**TableName=companies</br>**
**PipePath=./mypipe</br>**
**Delim=;</br>**

5) Run the software (./bin/macosx/pipegpss or ./bin/linux/pipegpss) </br>

Danieles-MacBook-Pro:bin dpalaia$ ./pipegpss</br>
**2019/03/14 15:58:11 Starting the connector and reading properties in the properties.ini file</br>**
**2019/03/14 15:58:11 Properties read: Connecting to the Grpc server specified</br>**
**2019/03/14 15:58:11 Connected to the grpc server</br>**
**2019/03/14 15:58:11 delegating to pipe client</br>**
**2019/03/14 15:58:11 Opening named pipe: ./mypipe for reading</br>**
**2019/03/14 15:58:11 waiting for someone to write something in the pipe</br></br>**

6) submit the example csv file provided in the pipe (./bin/macosx/data.csv): it contains 1000 elements </br></br>

**cat data.csv >> mypipe </br>**

7) you should see some logs in the pipegpss screen and the table populated with 1K elements </br>

**test=# select count(*) from people;</br>**
** count </br>**
**-------</br>**
**  1000</br>**
**(1 row)</br>**
