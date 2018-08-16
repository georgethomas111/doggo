package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/georgethomas111/doggohttp/network"
	"github.com/georgethomas111/doggohttp/stats/plot"
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

func handleSniff(intName string, portStr string) error {
	c := plot.New(portStr)
	n, err := network.New(intName, c)
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
	var intName = flag.String("interface", "wlan0", "The interface to sniff.")
	var ls = flag.Bool("ls", false, "List interfaces")
	var port = flag.String("port", ":8080", "Port to listen for web requests. eg :8080")

	flag.Parse()

	if *ls {
		handleLS()
		return
	}

	err := handleSniff(*intName, *port)
	if err != nil {
		log.Println("Sniff error ", err.Error())
		return
	}

}
