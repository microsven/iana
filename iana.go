package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*

url: "http://ftp.apnic.net/stats/afrinic/delegated-afrinic-latest",
url: "http://ftp.apnic.net/stats/lacnic/delegated-lacnic-latest",
url: "http://ftp.apnic.net/stats/apnic/delegated-apnic-latest",
url: "http://ftp.apnic.net/stats/ripe-ncc/delegated-ripencc-latest"
url: "http://ftp.apnic.net/stats/arin/delegated-arin-extended-latest",

cat ripencc | grep allocated | grep ipv4 | awk -F '|' '{print $2 "|" $4 "|"  $5}'
*/

type addr_mask struct {
	mask uint32
	ents map[uint32]*route_entry
}

type route_entry struct {
	country string
	code    int32
}

type iana_context struct {
	dir   string
	masks [4]addr_mask
	areas []string
}

var iana_ctx = iana_context{
	areas: []string{
		"afrinic",
		"lacnic",
		"apnic",
		"ripencc",
		"arin",
		"reserved",
	},
}

func iana_init(dir string) {
	for i := 0; i < 4; i += 1 {
		iana_ctx.masks[i].ents = make(map[uint32]*route_entry)
	}

	iana_ctx.masks[0].mask = 0xff000000
	iana_ctx.masks[1].mask = 0xffff0000
	iana_ctx.masks[2].mask = 0xffffff00
	iana_ctx.masks[3].mask = 0xffffffff

	for _, area := range iana_ctx.areas {
		fname := dir + "/" + area
		iana_entry_from_file(fname)
	}

}

func iana_entry_add(addr uint32, num uint32, country string) {
	masklen := 4
	n0 := 0
	for ; num&0xff == 0; num = num >> 8 {
		masklen -= 1
		n0 += 8
	}
	for i := uint32(0); i < num; i += 1 {
		new_addr := addr + (i << n0)
		iana_ctx.masks[masklen-1].ents[new_addr] = &route_entry{
			country: country,
		}
/*
		fmt.Printf("add addr-%v.%v.%v.%v/%v\n",
			(new_addr>>24)&0xff,
			(new_addr>>16)&0xff,
			(new_addr>>8)&0xff,
			new_addr&0xff, masklen*8)
*/
	}
}

func iana_search(addr uint32) int32 {
	for i := 4; i > 0; i -= 1 {
		index := i - 1
		mask := iana_ctx.masks[index]
		addr_masked := addr & mask.mask
		ent, ok := mask.ents[addr_masked]
		if ok {
			return region_shortcuts2code(ent.country)
		}
	}

	return 0
}

func iana_entry_from_file(file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	br := bufio.NewReader(f)
	for {
		line, err := br.ReadString('\n')
		if err != nil {
			break
		}

		strs := strings.Split(line, "|")
		if len(strs) != 3 {
			fmt.Printf("line format error, %v\n", line)
			continue
		}

		var a1, a2, a3, a4 uint32
		var number uint32

		n, err := fmt.Sscanf(strs[1], "%v.%v.%v.%v", &a1, &a2, &a3, &a4)
		if err != nil || n != 4 {
			fmt.Printf("addr format error, n : %v, err : %v\n", n, err)
			continue
		}
		n, err = fmt.Sscanf(strs[2], "%v", &number)
		if err != nil || n != 1 {
			fmt.Printf("number format error, n : %v, err : %v\n", n, err)
			continue
		}
		//fmt.Printf("add %v|%v|%v\n", strs[0], strs[1], strs[2])
		iana_entry_add((a1<<24)|(a2<<16)|(a3<<8)|a4, number, strs[0])
	}
}
