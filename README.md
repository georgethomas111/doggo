# README

Doggo http is a simple package that helps to read the details of the http going out or in from a network socket.
It should dump data for the visualizations to pick up in a sampled manner. This data can be used to analyse the network better.
Idea is to have the ability to understand the protocol data better to figure out issues with a paticular interface. 

# Components 

* network - This is the package that interacts with the system network interfaces.
    * interfaces.go - list of network interfaces.
    * interfaces_error.go - the errors that the interfaces can give.
    * interfaces_test.go - tests for the network interfaces.

* network/httpdecoder - decodes HTTP data from an os interface that is passed.

# TODO

* Use statsd to dump the data about different layers into it. 

* Use the interfaces package to get an interface and dump some bytes from it.
   * Create a socket and bind the interface we found earlier.
   * Use ioctl commands if required to point it to the buffer information.

