package main

import (
	"fmt"
	"flag"
)

func main() {
	country_init()
	var addr string
	iana_init("/home/giggs/pangolin/iana")
	flag.StringVar(&addr, "addr", "", "address")
	flag.Parse()
	var a1, a2, a3, a4 uint32
	fmt.Sscanf(addr, "%v.%v.%v.%v", &a1, &a2, &a3, &a4)
	fmt.Printf("%v.%v.%v.%v belongs to %v\n", a1, a2, a3, a4, iana_search((a1 << 24) | (a2 << 16) | (a3 << 8) | a1))
}
