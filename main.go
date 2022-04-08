package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/maxsokolovsky/updrop/config"
	"github.com/maxsokolovsky/updrop/server"
)

func main() {
	key := os.Getenv("SECRET_KEY")
	addr := flag.Int("addr", 8000, "Server addr port")
	flag.Parse()

	c := config.Config{
		Addr: ":" + strconv.Itoa(*addr),
		Key:  key,
	}

	if key != "" {
		c.ServerWideKey = true
	}
	s := server.New(c)
	log.Printf("Listening on port %s\n", c.Addr)
	log.Fatal(s.ListenAndServe())
}
