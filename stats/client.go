package stats

type Client interface {
	Receive(map[string]int)
}
