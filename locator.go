package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

const (
	LOCATOR_NUM = 4
)

type LocatorContext struct {
	searcher *xdb.Searcher
}

var locator_ctx LocatorContext

var (
	locate_err_notv4  = errors.New("not ipv4")
	locate_err_search = errors.New("search error")
)

func locate_addr(addr uint32) (int32, error) {
	var code int32
	var strs []string
	var result string

	result, err := locator_ctx.searcher.Search(addr)
	if err != nil {
		goto err_out
	}

	strs = strings.Split(result, "|")
	if len(strs) > 0 {
		code = region_cn2code(strs[0])
	} else {
		err = locate_err_search
	}

err_out:
	return code, err
}

func locator_init(file string) {
	cbuf, err := xdb.LoadContentFromFile(file)
	if err != nil {
		panic(fmt.Sprintf("failed to load content from `%s`: %s\n", file, err))
	}

	searcher, err := xdb.NewWithBuffer(cbuf)
	if err != nil {
		panic(fmt.Sprintf("failed to create searcher with content: %s\n", err))
	}

	locator_ctx.searcher = searcher
}
