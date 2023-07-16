#!/bin/sh

cat ripencc.original  | grep "assigned\|allocated" | grep ipv4 | awk -F '|' '{print $2 "|" $4 "|"  $5}' > ripencc
cat lacnic.original  | grep "assigned\|allocated" | grep ipv4 | awk -F '|' '{print $2 "|" $4 "|"  $5}' > lacnic
cat arin.original  | grep "assigned\|allocated" | grep ipv4 | awk -F '|' '{print $2 "|" $4 "|"  $5}' > arin
cat afrinic.original  | grep "assigned\|allocated" | grep ipv4 | awk -F '|' '{print $2 "|" $4 "|"  $5}' > afrinic

