# Readme

Doggo is a simple package that helps to read the details of the http going out or in from a network socket.
It should dump data for the visualizations to pick up. This data can be used to analyse the network better.
Idea is to have the ability to understand the protocol data better to figure out issues with a paticular interface. 

# Design

![blank diagram](https://user-images.githubusercontent.com/778330/44226517-38a5b780-a145-11e8-9b85-59f888c068f0.png)

# Steps to Run

 Assuming your go environment is all setup. Just do.

```
$ go get github.com/georgethomas111/doggo
$ cd $WORKSPACE/github.com/georgethomas111/doggo
$ go run cmd/doggo.go
```

 To get a detailed list of options, use.
```
$ go run cmd/doggo.go -help
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

# Components 

* network - This is the package that interacts with the system network interfaces.
          - The stats client is called which handles the data.
* stats - stats collects the stats from the network and uses the plotting libraries to create visualizations. 

# Future work

* Add a UI for stats with bootstrap.
* Clean the data and dump into a te series data base for easy querying.


