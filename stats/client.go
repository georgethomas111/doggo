package stats

type Client interface {
	Receive(interface{})
}
