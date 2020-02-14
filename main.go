package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port = flag.Int("p", 8000, "port to serve on")
	dir  = flag.String("d", ".", "the directory of static file to host")
)

func main() {
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir(*dir)))
	log.Printf("HTTP server start on port: %v\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *port), nil))
}
