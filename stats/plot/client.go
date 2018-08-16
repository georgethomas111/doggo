package plot

import (
	"net/http"
	"time"

	"github.com/georgethomas111/doggohttp/stats"
	"github.com/google/gopacket/layers"
	chart "github.com/wcharczuk/go-chart"
)

type Client struct {
	ackCount  int
	packCount int
	l         *Line
}

func New(port string) *Client {
	l := &Line{
		XName:  "Time",
		YName:  "Percent",
		Width:  1280,
		Height: 748,
	}

	c := &Client{
		l: l,
	}

	go c.Listen(port)
	return c
}

func (c *Client) Receive(ps stats.PacketStats) {
	c.plotTCP(ps.TCP())
}

func (c *Client) plotTCP(tcp *layers.TCP) {
	if tcp == nil {
		return
	}
	c.packCount++
	timestamp := time.Now()
	if tcp.ACK {
		c.ackCount++
	}
	percent := (float64(c.ackCount) / float64(c.packCount)) * 100
	c.l.X = append(c.l.X, timestamp)
	c.l.Y = append(c.l.Y, percent)
}

func (c *Client) DrawAckPercent(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", chart.ContentTypePNG)
	c.l.DrawLine(res)
}

func (c *Client) Listen(port string) {
	http.HandleFunc("/", c.DrawAckPercent)
	http.ListenAndServe(port, nil)
}
