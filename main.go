package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/maxsokolovsky/updrop/config"
	"github.com/maxsokolovsky/updrop/server"
)

func main() {
	port := flag.Int("port", 8000, "server port")
	flag.Parse()

	c := config.Config{
		Addr: ":" + strconv.Itoa(*port),
	}

	s := server.New(c)
	log.Printf("Listening on %s\n", c.Addr)
	log.Fatal(s.ListenAndServe())
}
