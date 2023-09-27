package main
/*
import (
	"fmt"
	"math/rand"
)

const (
	TEST_COUNT = 10000000
)

type test_exept struct {
	code_from_iana    int32
	code_from_xdb     int32
	country_from_iana string
	country_from_xdb  string
}

func main() {
	country_init()
	iana_init(".")
	locator_init("ip2region.xdb")

	checks := make(map[uint32]struct{})
	for i := 0; i < TEST_COUNT; i += 1 {
		checks[rand.Uint32()] = struct{}{}
	}

	fmt.Printf("test %v entries\n", len(checks))

	except := make(map[uint32]struct{})
	results := make(map[uint32]*test_exept)
	for k := range checks {
		code_from_iana := iana_search(k)
		code_from_xdb, _ := locate_addr(k)
		results[k] = &test_exept{
			code_from_iana:    code_from_iana,
			code_from_xdb:     code_from_xdb,
			country_from_iana: region_code2en(code_from_iana),
			country_from_xdb:  region_code2en(code_from_xdb),
		}

		if code_from_iana != code_from_xdb {
			except[k] = struct{}{}
		}
	}

	fmt.Printf("mismatch %v entries percent %v\n", len(except), float64(len(except))/float64(len(checks)))

	for k := range except {
		v := results[k]
		fmt.Printf("%v.%v.%v.%v iana code %v : %v  country xdb : %v\n",
			(k>>24)&0xff,
			(k>>16)&0xff,
			(k>>8)&0xff,
			(k>>0)&0xff,
			v.code_from_iana, region_code2en(v.code_from_iana), region_code2en(v.code_from_xdb))
	}
}
*/
