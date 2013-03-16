package main

import (
	"log"
)

func main() {
	log.Println("hello")

	// You can see flags define in http://golang.org/pkg/log/#pkg-constants
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Println("hello")

	// default format
	log.SetFlags(log.LstdFlags)

	// set prefix string
	log.SetPrefix("server1 ")
	log.Println("hello")
}
