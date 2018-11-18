package stats

// Client is any client that accepts an object with Packet.
type Client interface {
	Receive(Packet)
}
