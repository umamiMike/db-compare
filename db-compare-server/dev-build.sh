#!/bin/bash
cd cmd/db-compare-server
go get -u

build_cmd () {
clear
$(go build -o ./dev/server) > /dev/null && printf "\e[1m\e[36mapp rebuilt successfully\e[0m\n\n" || printf "\n\n\e[1m[91mrebuild failed\n\n"
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
