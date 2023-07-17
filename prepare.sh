#!/bin/sh

cat ripencc.original  | grep "assigned\|allocated" | grep ipv4 | awk -F '|' '{print $2 "|" $4 "|"  $5}' > ripencc
cat lacnic.original  | grep "assigned\|allocated" | grep ipv4 | awk -F '|' '{print $2 "|" $4 "|"  $5}' > lacnic
cat arin.original  | grep "assigned\|allocated" | grep ipv4 | awk -F '|' '{print $2 "|" $4 "|"  $5}' > arin
cat afrinic.original  | grep "assigned\|allocated" | grep ipv4 | awk -F '|' '{print $2 "|" $4 "|"  $5}' > afrinic
cat apnic.original  | grep "assigned\|allocated" | grep ipv4 | awk -F '|' '{print $2 "|" $4 "|"  $5}' > apnic

cat ripencc.original  | grep "reserved" | grep ipv4 | awk -F '|' '{print "RSV|" $4 "|"  $5}' > reserved
cat lacnic.original  | grep "reserved" | grep ipv4 | awk -F '|' '{print "RSV|" $4 "|"  $5}' >> reserved
cat arin.original  | grep "reserved" | grep ipv4 | awk -F '|' '{print "RSV|" $4 "|"  $5}' >> reserved
cat afrinic.original  | grep "reserved" | grep ipv4 | awk -F '|' '{print "RSV|" $4 "|"  $5}' >> reserved
cat apnic.original  | grep "reserved" | grep ipv4 | awk -F '|' '{print "RSV|" $4 "|"  $5}' >> reserved 
