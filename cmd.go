package main

import "log"
import "flag"

func main() {
	var extIF, intIF string
	flag.StringVar(&extIF, "ext", "eth0", "internal IPv6 interface")
	flag.StringVar(&intIF, "int", "sit0", "external IPv6 interface")
	flag.Parse()

	log.Printf("Launching 6rdrtr, external interface %s, internal interface %s", extIF, intIF)

	log.Fatalln("Just kidding, there's literally nothing here.")
}
