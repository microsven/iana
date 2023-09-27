package main

type country_info struct {
	chinese   string
	english   string
	shortcuts string
	code      int32
}


func region_cn2code(cn string) int32 {
	code, ok := countries_table.cn2code[cn]
	if ok {
		return code
	}
	return 0
}

func region_code2en(code int32) string {
	cty, ok := countries_table.code2en[code]
	if ok {
		return cty
	}
	return "unkown"
}

func region_shortcuts2code(sc string) int32 {
	code, ok := countries_table.shortcuts2code[sc]
	if ok {
		return code
	}
	return 0
}

func region_code2shortcuts(code int32) string {
	sc, ok := countries_table.code2shortcuts[code]
	if ok {
		return sc
	}
	return "unkown"
	
}

type country_context struct {
	cn2code        map[string]int32
	code2en        map[int32]string
	code2shortcuts  map[int32]string
	shortcuts2code map[string]int32
}

var countries_table country_context

func country_init() {
	countries_table.cn2code = make(map[string]int32)
	countries_table.shortcuts2code = make(map[string]int32)
	countries_table.code2en = make(map[int32]string)
	countries_table.code2shortcuts = make(map[int32]string)

	for _, c := range all_cuntries {
		countries_table.cn2code[c.chinese] = c.code
		countries_table.shortcuts2code[c.shortcuts] = c.code
		countries_table.code2shortcuts[c.code] = c.shortcuts
		countries_table.code2en[c.code] = c.english
	}
}

var all_cuntries = []country_info{
