#!/bin/bash
inotifywait ./*.go -m -e close_write 
while read events; do 
  killall db-compare-server 
  go build && ./db-compare-server &; 
  echo "app restarted"; 
done

