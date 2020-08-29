  $(killall db-compare-server)
clear
  go build && echo " \e[1m\e[36mapp rebuilt successfully\e[0m" || echo -e "\n\n\e[1m[91mrebuild failed\n\n"
  ./db-compare-server & 
inotifywait ./ -q -m -e close_write | while read events; do 
clear
  $(killall db-compare-server)
  go build && echo " \e[1m\e[36mapp rebuilt successfully\e[0m" || echo -e "\n\n\e[1m[91mrebuild failed\n\n"
  ./db-compare-server & 
done
