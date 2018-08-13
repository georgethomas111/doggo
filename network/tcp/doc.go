// tcp is a package which aggregates data from an interfacce and returns the
// number of packets lost in each trigger compared to the previous
// trigger. It saves the data locally until the call to the next trigger
// which makes it zero again.
package tcp
