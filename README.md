# README

Doggo http is a simple package that helps to read the details of the http going out or in from a network socket.

This uses https://github.com/google/gopacket to do the implementation.

# Components 

* network - This is the package that interacts with the system network interfaces.
    * interfaces.go - list of network interfaces.
    * interfaces_error.go - the errors that the interfaces can give.
    * interfaces_test.go - tests for the network interfaces.

* network/httpdecoder - decodes HTTP data from an os interface that is passed.

# TODO

* Use the interfaces package to get an interface and dump some bytes from it.
