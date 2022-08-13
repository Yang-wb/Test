package main

import (
	"flag"
	"fmt"
	"leango/crawler_distributed/rpcsupport"
	"leango/crawler_distributed/worker"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen to")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.SaverRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}
