src=$(wildcard *.go)

.PHONY:all clean

all:iana
iana: $(src)
	cp -f country.txt  country_tmp.txt
	cp -f country.go country_tmp.go
	mv country.go country_backup._go
	sed -i '1d' country_tmp.txt 
	cat country_tmp.txt >> country_tmp.go
	go build
	mv country_backup._go country.go

