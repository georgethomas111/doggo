package stats

// Client is any client that accepts an object with PacketStats.
type Client interface {
	Receive(PacketStats)
}
