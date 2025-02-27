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
	cert := flag.String("cert", "", "path to cert")
	key := flag.String("key", "", "path to private key")
	flag.Parse()

	if *cert == "" {
		log.Fatal("-cert is required")
	}
	if *key == "" {
		log.Fatal("-key is required")
	}

	c := config.Config{
		Addr: ":" + strconv.Itoa(*port),
	}
	s := server.New(c)
	log.Printf("Listening on %s.\n", c.Addr)
	log.Fatal(s.ListenAndServeTLS(*cert, *key))
}
