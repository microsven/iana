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
	{"安道尔", "Andorra", "AD", 20},
	{"阿联酋", "United Arab Emirates", "AE", 784},
	{"阿富汗", "Afghanistan", "AF", 4},
	{"安提瓜和巴布达", "Antigua and Barbuda", "AG", 28},
	{"安圭拉", "Anguilla", "AI", 660},
	{"阿尔巴尼亚", "Albania", "AL", 8},
	{"亚美尼亚", "Armenia", "AM", 51},
	{"安哥拉", "Angola", "AO", 24},
	{"南极洲", "Antarctica", "AQ", 10},
	{"阿根廷", "Argentina", "AR", 32},
	{"美属萨摩亚", "American Samoa", "AS", 16},
	{"奥地利", "Austria", "AT", 40},
	{"澳大利亚", "Australia", "AU", 36},
	{"阿鲁巴", "Aruba", "AW", 533},
	{"奥兰群岛", "Ahvenanmaa", "AX", 248},
	{"阿塞拜疆", "Azerbaijan", "AZ", 31},
	{"波黑", "Bosnia and Herzegovina", "BA", 70},
	{"巴巴多斯", "Barbados", "BB", 52},
	{"孟加拉", "Bangladesh", "BD", 50},
	{"比利时", "Belgium", "BE", 56},
	{"布基纳法索", "Burkina", "BF", 854},
	{"保加利亚", "Bulgaria", "BG", 100},
	{"巴林", "Bahrain", "BH", 48},
	{"布隆迪", "Burundi", "BI", 108},
	{"贝宁", "Benin", "BJ", 204},
	{"圣巴泰勒米岛", "Saint Barthélemy", "BL", 652},
	{"百慕大", "Bermuda", "BM", 60},
	{"文莱", "Brunei", "BN", 96},
	{"玻利维亚", "Bolivia", "BO", 68},
	{"荷兰加勒比区", "Caribbean Netherlands", "BQ", 535},
	{"巴西", "Brazil", "BR", 76},
	{"巴哈马", "Bahamas", "BS", 44},
	{"不丹", "Bhutan", "BT", 64},
	{"布韦岛", "Bouvet Island", "BV", 74},
	{"博茨瓦纳", "Botswana", "BW", 72},
	{"白俄罗斯", "Belarus", "BY", 112},
	{"伯利兹", "Belize", "BZ", 84},
	{"加拿大", "Canada", "CA", 124},
	{"科科斯群岛", "Cocos Islands", "CC", 166},
	{"中非", "Central African Republic", "CF", 140},
	{"瑞士", "Switzerland", "CH", 756},
	{"智利", "Chile", "CL", 152},
	{"喀麦隆", "Cameroon", "CM", 120},
	{"哥伦比亚", "Colombia", "CO", 170},
	{"哥斯达黎加", "Costa Rica", "CR", 188},
	{"古巴", "Cuba", "CU", 192},
	{"佛得角", "Cape Verde", "CV", 132},
	{"圣诞岛", "Christmas Island", "CX", 162},
	{"塞浦路斯", "Cyprus", "CY", 196},
	{"捷克", "Czech", "CZ", 203},
	{"德国", "Germany", "DE", 276},
	{"吉布提", "Djibouti", "DJ", 262},
	{"丹麦", "Denmark", "DK", 208},
	{"多米尼克", "Dominica", "DM", 212},
	{"多米尼加", "Dominican Republic", "DO", 214},
	{"阿尔及利亚", "Algeria", "DZ", 12},
	{"厄瓜多尔", "Ecuador", "EC", 218},
	{"爱沙尼亚", "Estonia", "EE", 233},
	{"埃及", "Egypt", "EG", 818},
	{"西撒哈拉", "Western Sahara", "EH", 732},
	{"厄立特里亚", "Eritrea", "ER", 232},
	{"西班牙", "Spain", "ES", 724},
	{"芬兰", "Finland", "FI", 246},
	{"斐济群岛", "Fiji", "FJ", 242},
	{"马尔维纳斯群岛", "Falkland Islands", "FK", 238},
	{"密克罗尼西亚联邦", "Federated States of Micronesia", "FM", 583},
	{"法罗群岛", "Faroe Islands", "FO", 234},
	{"法国", "France", "FR", 250},
	{"加蓬", "Gabon", "GA", 266},
	{"格林纳达", "Grenada", "GD", 308},
	{"格鲁吉亚", "Georgia", "GE", 268},
	{"法属圭亚那", "French Guiana", "GF", 254},
	{"加纳", "Ghana", "GH", 288},
	{"直布罗陀", "Gibraltar", "GI", 292},
	{"格陵兰", "Greenland", "GL", 304},
	{"几内亚", "Guinea", "GN", 324},
	{"瓜德罗普", "Guadeloupe", "GP", 312},
	{"赤道几内亚", "Equatorial Guinea", "GQ", 226},
	{"希腊", "Greece", "GR", 300},
	{"南乔治亚岛和南桑威奇群岛", "South Georgia and the South Sandwich Islands", "GS", 239},
	{"危地马拉", "Guatemala", "GT", 320},
	{"关岛", "Guam", "GU", 316},
	{"几内亚比绍", "Guinea Bissau", "GW", 624},
	{"圭亚那", "Guyana", "GY", 328},
	{"香港", "Hong Kong", "HK", 344},
	{"赫德岛和麦克唐纳群岛", "Heard Island and McDonald Islands", "HM", 334},
	{"洪都拉斯", "Honduras", "HN", 340},
	{"克罗地亚", "Croatia", "HR", 191},
	{"海地", "Haiti", "HT", 332},
	{"匈牙利", "Hungary", "HU", 348},
	{"印尼", "Indonesia", "ID", 360},
	{"爱尔兰", "Ireland", "IE", 372},
	{"以色列", "Israel", "IL", 376},
	{"马恩岛", "Isle of Man", "IM", 833},
	{"印度", "India", "IN", 356},
	{"英属印度洋领地", "British Indian Ocean Territory", "IO", 86},
	{"伊拉克", "Iraq", "IQ", 368},
	{"伊朗", "Iran", "IR", 364},
	{"冰岛", "Iceland", "IS", 352},
	{"意大利", "Italy", "IT", 380},
	{"泽西岛", "Jersey", "JE", 832},
	{"牙买加", "Jamaica", "JM", 388},
	{"约旦", "Jordan", "JO", 400},
	{"日本", "Japan", "JP", 392},
	{"柬埔寨", "Cambodia", "KH", 116},
	{"基里巴斯", "Kiribati", "KI", 296},
	{"科摩罗", "Comoros", "KM", 174},
	{"科威特", "Kuwait", "KW", 414},
	{"开曼群岛", "Cayman Islands", "KY", 136},
	{"黎巴嫩", "Lebanon", "LB", 422},
	{"列支敦士登", "Liechtenstein", "LI", 438},
	{"斯里兰卡", "Sri Lanka", "LK", 144},
	{"利比里亚", "Liberia", "LR", 430},
	{"莱索托", "Lesotho", "LS", 426},
	{"立陶宛", "Lithuania", "LT", 440},
	{"卢森堡", "Luxembourg", "LU", 442},
	{"拉脱维亚", "Latvia", "LV", 428},
	{"利比亚", "Libya", "LY", 434},
	{"摩洛哥", "Morocco", "MA", 504},
	{"摩纳哥", "Monaco", "MC", 492},
	{"摩尔多瓦", "Moldova", "MD", 498},
	{"黑山", "Montenegro", "ME", 499},
	{"法属圣马丁", "Saint Martin France", "MF", 663},
	{"马达加斯加", "Madagascar", "MG", 450},
	{"马绍尔群岛", "Marshall islands", "MH", 584},
	{"马其顿", "Macedonia", "MK", 807},
	{"马里", "Mali", "ML", 466},
	{"缅甸", "Myanmar", "MM", 104},
	{"澳门", "Macao", "MO", 446},
	{"马提尼克", "Martinique", "MQ", 474},
	{"毛里塔尼亚", "Mauritania", "MR", 478},
	{"蒙塞拉特岛", "Montserrat", "MS", 500},
	{"马耳他", "Malta", "MT", 470},
	{"马尔代夫", "Maldives", "MV", 462},
	{"马拉维", "Malawi", "MW", 454},
	{"墨西哥", "Mexico", "MX", 484},
	{"马来西亚", "Malaysia", "MY", 458},
	{"纳米比亚", "Namibia", "NA", 516},
	{"尼日尔", "Niger", "NE", 562},
	{"诺福克岛", "Norfolk Island", "NF", 574},
	{"尼日利亚", "Nigeria", "NG", 566},
	{"尼加拉瓜", "Nicaragua", "NI", 558},
	{"荷兰", "Netherlands", "NL", 528},
	{"挪威", "Norway", "NO", 578},
	{"尼泊尔", "Nepal", "NP", 524},
	{"瑙鲁", "Nauru", "NR", 520},
	{"阿曼", "Oman", "OM", 512},
	{"巴拿马", "Panama", "PA", 591},
	{"秘鲁", "Peru", "PE", 604},
	{"法属波利尼西亚", "French polynesia", "PF", 258},
	{"巴布亚新几内亚", "Papua New Guinea", "PG", 598},
	{"菲律宾", "Philippines", "PH", 608},
	{"巴基斯坦", "Pakistan", "PK", 586},
	{"波兰", "Poland", "PL", 616},
	{"皮特凯恩群岛", "Pitcairn Islands", "PN", 612},
	{"波多黎各", "Puerto Rico", "PR", 630},
	{"巴勒斯坦", "Palestinian", "PS", 275},
	{"帕劳", "Palau", "PW", 585},
	{"巴拉圭", "Paraguay", "PY", 600},
	{"卡塔尔", "Qatar", "QA", 634},
	{"留尼汪", "Réunion", "RE", 638},
	{"罗马尼亚", "Romania", "RO", 642},
	{"塞尔维亚", "Serbia", "RS", 688},
	{"俄罗斯", "Russian", "RU", 643},
	{"卢旺达", "Rwanda", "RW", 646},
	{"所罗门群岛", "Solomon Islands", "SB", 90},
	{"塞舌尔", "Seychelles", "SC", 690},
	{"苏丹", "Sudan", "SD", 729},
	{"瑞典", "Sweden", "SE", 752},
	{"新加坡", "Singapore", "SG", 702},
	{"斯洛文尼亚", "Slovenia", "SI", 705},
	{"斯瓦尔巴群岛和", "SJM Svalbard", "SJ", 744},
	{"斯洛伐克", "Slovakia", "SK", 703},
	{"塞拉利昂", "Sierra Leone", "SL", 694},
	{"圣马力诺", "San Marino", "SM", 674},
	{"塞内加尔", "Senegal", "SN", 686},
	{"索马里", "Somalia", "SO", 706},
	{"苏里南", "Suriname", "SR", 740},
	{"南苏丹", "South Sudan", "SS", 728},
	{"圣多美和普林西比", "Sao Tome Principe", "ST", 678},
	{"萨尔瓦多", "Salvador", "SV", 222},
	{"叙利亚", "Syria", "SY", 760},
	{"斯威士兰", "Swaziland", "SZ", 748},
	{"特克斯和凯科斯群岛", "Turks Caicos Islands", "TC", 796},
	{"乍得", "Chad", "td", 148},
	{"多哥", "Togo", "TG", 768},
	{"泰国", "Thailand", "TH", 764},
	{"托克劳", "Tokelau", "TK", 772},
	{"东帝汶", "Timor Leste", "TP", 626},
	{"突尼斯", "Tunisia", "TN", 788},
	{"汤加", "Tonga", "TO", 776},
	{"土耳其", "Turkey", "TR", 792},
	{"图瓦卢", "Tuvalu", "TV", 798},
	{"坦桑尼亚", "Tanzania", "TZ", 834},
	{"乌克兰", "Ukraine", "UA", 804},
	{"乌干达", "Uganda", "UG", 800},
	{"美国", "United States of America", "US", 840},
	{"乌拉圭", "Uruguay", "UY", 858},
	{"梵蒂冈", "Vatican", "VA", 336},
	{"委内瑞拉", "Venezuela", "VE", 862},
	{"英属维尔京群岛", "British Virgin Islands", "VG", 92},
	{"美属维尔京群岛", "United States Virgin Islands", "VI", 850},
	{"越南", "Vietnam", "VN", 704},
	{"瓦利斯和富图纳", "Wallis and Futuna", "WF", 876},
	{"萨摩亚", "Samoa", "WS", 882},
	{"也门", "Yemen", "YE", 887},
	{"马约特", "Mayotte", "YT", 175},
	{"南非", "South Africa", "ZA", 710},
	{"赞比亚", "Zambia", "ZM", 894},
	{"津巴布韦", "Zimbabwe", "ZW", 716},
	{"中国", "China", "CN", 156},
	{"刚果（布）", "Republic of the Congo", "CG", 178},
	{"刚果（金）", "Democratic Republic of the Congo", "CD", 180},
	{"莫桑比克", "Mozambique", "MZ", 508},
	{"根西岛", "Guernsey", "GG", 831},
	{"冈比亚", "Gambia", "GM", 270},
	{"北马里亚纳群岛", "Northern Mariana Islands", "MP", 580},
	{"埃塞俄比亚", "Ethiopia", "ET", 231},
	{"新喀里多尼亚", "New Caledonia", "NC", 540},
	{"瓦努阿图", "Vanuatu", "VU", 548},
	{"法属南部领地", "French Southern Territories", "TF", 260},
	{"纽埃", "Niue", "NU", 570},
	{"美国本土外小岛屿", "United States Minor Outlying Islands", "UM", 581},
	{"库克群岛", "Cook Islands", "CK", 184},
	{"英国", "Great Britain", "GB", 826},
	{"特立尼达和多巴哥", "Trinidad Tobago", "TT", 780},
	{"圣文森特和格林纳丁斯", "St. Vincent and the Grenadines", "VC", 670},
	{"台湾", "Taiwan", "TW", 158},
	{"新西兰", "New Zealand", "NZ", 554},
	{"沙特阿拉伯", "Saudi Arabia", "SA", 682},
	{"老挝", "Laos", "LA", 418},
	{"朝鲜", "North Korea", "KP", 408},
	{"韩国", "South Korea", "KR", 410},
	{"葡萄牙", "Portugal", "PT", 620},
	{"吉尔吉斯斯坦", "Kyrgyzstan", "KG", 417},
	{"哈萨克斯坦", "Kazakhstan", "KZ", 398},
	{"塔吉克斯坦", "Tajikistan", "TJ", 762},
	{"土库曼斯坦", "Turkmenistan", "TM", 795},
	{"乌兹别克斯坦", "Uzbekistan", "UZ", 860},
	{"圣基茨和尼维斯", "St.Kitts Nevis", "KN", 659},
	{"圣皮埃尔和密克隆", "Saint-Pierre and Miquelon", "PM", 666},
	{"圣赫勒拿", "St.Helena Dependencies", "SH", 654},
	{"圣卢西亚", "St.Lucia", "LC", 662},
	{"毛里求斯", "Mauritius", "MU", 480},
	{"科特迪瓦", "Coate dIvoire", "CI", 384},
	{"肯尼亚", "Kenya", "KE", 404},
	{"蒙古", "Mongolia", "MN", 496},
	{"随机", "Random", "RD", 0},
	{"保留", "Reserved", "RV", 65536}}
