package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/georgethomas111/doggo/pkg/api"
	"github.com/georgethomas111/doggo/pkg/db/memdb"
	"github.com/georgethomas111/doggo/pkg/heartbeat"
	"github.com/georgethomas111/doggo/pkg/network"
	"github.com/georgethomas111/doggo/pkg/service"
	"github.com/georgethomas111/doggo/pkg/service/bee"
	"github.com/georgethomas111/doggo/pkg/stats"
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

func services(portStr string, db bee.DB, hb int, s *service.Stop) []stats.Client {
	//	c := plot.New(portStr)
	b := bee.New(db)
	h := heartbeat.New(time.Millisecond*time.Duration(hb), []heartbeat.Application{b})
	s.Add(h)

	return []stats.Client{b}
}

func waitForStop() {

}

func addTowaitForStop(apps []interface{ Close() }) {
}

func handleSniff(intName string, portStr string, db bee.DB, hb int, s *service.Stop) error {
	n, err := network.New(intName, services(portStr, db, hb, s))
	if err != nil {
		return err
	}
	s.Add(n)
	return nil
}

func main() {
	var intName = flag.String("interface", "wlan0", "The network interface to sniff.")
	var ls = flag.Bool("ls", false, "List interfaces")
	var port = flag.String("port", ":8080", "Port to look at the UI.")
	var jPort = flag.String("jport", ":9000", "Port to listen for api web requests.")
	var heartbeat = flag.Int("heartbeat", 1000, "The sampling heartbeat in ms")

	flag.Parse()

	if *ls {
		handleLS()
		return
	}

	stop := &service.Stop{}
	db := memdb.New()

	err := handleSniff(*intName, *port, db, *heartbeat, stop)
	if err != nil {
		log.Println("Sniff error ", err.Error())
		return
	}

	// This will wait to receive api requests.
	api := api.New(*jPort, db)
	stop.Add(api)
	stop.Wait()

}
