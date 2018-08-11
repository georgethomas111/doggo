# Readme

Doggo http is a simple package that helps to read the details of the http going out or in from a network socket.
It should dump data for the visualizations to pick up in a sampled manner. This data can be used to analyse the network better.
Idea is to have the ability to understand the protocol data better to figure out issues with a paticular interface. 

# Design

![blank diagram 1](https://user-images.githubusercontent.com/778330/43988406-f5ec2cee-9ce9-11e8-9924-5a742c18e53d.png)


# Components 

* network - This is the package that interacts with the system network interfaces.
    * interfaces.go - list of network interfaces.
    * interfaces_error.go - the errors that the interfaces can give.
    * interfaces_test.go - tests for the network interfaces.

* network/httpdecoder - decodes HTTP data from an os interface that is passed.

# Future work

* Use statsd to dump the data about different layers into it. 
    * tcp packet type check
    
* Add UI for graphana.
