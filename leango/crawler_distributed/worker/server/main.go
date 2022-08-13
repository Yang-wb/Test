package main

import (
	"leango/crawler_distributed/config"
	"leango/crawler_distributed/rpcsupport"
	"log"
)

func main() {
	log.Fatal(rpcsupport.SaverRpc(":%d", config.WorkerPort0))
}
