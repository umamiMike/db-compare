#!/bin/bash

build_cmd () {
clear
$(go build) > /dev/null && echo -e "\e[1m\e[36mapp rebuilt successfully\e[0m" || echo -e "\n\n\e[1m[91mrebuild failed\n\n"
./db-compare-server & 
disown
}

killall db-compare-server &> /dev/null

rebuild () {
killall db-compare-server &> /dev/null
build_cmd
fswatch --event Updated --one-event $(pwd) 
rebuild 
 }

rebuild

