package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/georgethomas111/doggo/db/memdb"
	"github.com/georgethomas111/doggo/heartbeat"
	"github.com/georgethomas111/doggo/network"
	"github.com/georgethomas111/doggo/service/bee"
	"github.com/georgethomas111/doggo/service/plot"
	"github.com/georgethomas111/doggo/stats"
)

func handleLS() {
	interfaces, err := network.LS()
	if err != nil {
		log.Println("LS error ", err.Error())
		return
	}

	for _, i := range interfaces {
		fmt.Println(i)
	}
	return
}

func services(portStr string, db bee.DB, hb int) []stats.Client {
	c := plot.New(portStr)
	b := bee.New(db)
	heartbeat.New(time.Millisecond*time.Duration(hb), []heartbeat.Application{b})

	return []stats.Client{c, b}
}

func handleSniff(intName string, portStr string, db bee.DB, hb int) error {
	n, err := network.New(intName, services(portStr, db, hb))
	if err != nil {
		return err
	}
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Waiting for interrupt")
	<-stop
	n.Close()
	fmt.Println("Received interrupt. Bye use me again.")
	return nil
}

func main() {
	var intName = flag.String("interface", "wlan0", "The network interface to sniff.")
	var ls = flag.Bool("ls", false, "List interfaces")
	var port = flag.String("port", ":8080", "Port to listen for web requests. eg :8080")
	var heartbeat = flag.Int("heartbeat", 1000, "The sampling heartbeat in ms")
	//	var jPort = flag.String("jport", ":8081", "Port to listen for api web requests. eg :8081")

	flag.Parse()

	if *ls {
		handleLS()
		return
	}

	err := handleSniff(*intName, *port, memdb.New(), *heartbeat)
	if err != nil {
		log.Println("Sniff error ", err.Error())
		return
	}

}
