# Readme

Doggo is a simple package that helps to read the details of the http going out or in from a network socket.
It should dump data for the visualizations to pick up. This data can be used to analyse the network better.
Idea is to have the ability to understand the protocol data better to figure out issues with a paticular interface. 

# Design

![blank diagram](https://user-images.githubusercontent.com/778330/44226517-38a5b780-a145-11e8-9b85-59f888c068f0.png)

# Steps to Run

 Assuming your go environment is all setup.

```
$ go get github.com/georgethomas111/doggo
$ cd $WORKSPACE/github.com/georgethomas111/doggo
$ go build cmd/doggo.go
```

 Usage options can be obtained by

```
$ ./doggo -help
```

 Help options

```
Usage of ./doggo:
  -interface string
        The interface to sniff. (default "wlan0")
  -ls
        List interfaces
  -port string
        Port to listen for web requests. eg :8080 (default ":8080")
```

 To list available interfaces to listen on.
```
$ go run cmd/doggo.go -ls
```

 Sample Output

```
lo
wlan0
arcbr0
veth_android
```

 To listen on wlan0 network interface.

```
$ sudo ./doggo
```
 sudo is required as the interface requires root permissions to read from. The graphs can be viewed by opening a browser and going to  
 http://localhost:8080

 Sample output
```
Waiting for interrupt
^CReceived interrupt. Bye use me again.
```

![screenshot 2018-08-16 at 11 44 39 am](https://user-images.githubusercontent.com/778330/44228294-d26f6380-a149-11e8-8930-f066ac89f98b.png)


 To listen on a different interface veth_andoid and to view the results 
```
$ sudo ./doggo -interface veth_android

```
 To change the port the web server is listening to use the following 
```
$ sudo ./doggo -interface veth_android -port :8081

```
 This will make sure the webpage is available at http://localhost:8081


# Components 

* network - This is the package that interacts with the system network interfaces.
          - The stats client is called which handles the data.
* stats - stats collects the stats from the network and uses the plotting libraries to create visualizations. 

# Future work

* Add a UI for stats with bootstrap.
* Clean the data and dump into a te series data base for easy querying.


