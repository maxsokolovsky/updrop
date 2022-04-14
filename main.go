package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/maxsokolovsky/updrop/config"
	"github.com/maxsokolovsky/updrop/server"
)

func main() {
	addr := flag.Int("addr", 8000, "Server addr port")
	flag.Parse()

	c := config.Config{
		Addr: ":" + strconv.Itoa(*addr),
	}

	s := server.New(c)
	log.Printf("Listening on port %s\n", c.Addr)
	log.Fatal(s.ListenAndServe())
}
