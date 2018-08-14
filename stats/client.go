package stats

type Client interface {
	Receive(PacketStats)
}
