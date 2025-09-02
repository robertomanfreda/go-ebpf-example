package main

import (
	"ebpf-test/gen"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	c, err := gen.Start("enp3s0")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	log.Printf("Counting incoming packets on %s..", c.Interface().Name)

	tick := time.NewTicker(time.Second)
	defer tick.Stop()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	for {
		select {
		case <-tick.C:
			n, err := c.Count()
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Received %d packets", n)
		case <-stop:
			log.Print("Received signal, exiting..")
			return
		}
	}
}
