#!/bin/bash
# Place the data directory inside the project directory
      # function cleanup {
      # echo "exiting the shell and killing postgres" 
      # killall postgres
      # }

# new window 
#    -c  is the starting directory
root="$(pwd)"
kill -INT `head -1 $(pwd)/postgres/postmaster.pid`
tmux kill-session -t db-compare-dev
sleep 2
tmux new-session -d -s db-compare-dev 

tmux new-window -t db-compare-dev:0 -n server -c $root
tmux send-keys -t db-compare-dev:0 'cd $(pwd)/db-compare-server/ && ./dev-build.sh' Enter

tmux new-window -t db-compare-dev:1 -n react-client -c $root
tmux send-keys -t db-compare-dev:1 'cd $(pwd)/react-client && npm install && npm run start' Enter 

tmux new-window -t db-compare-dev:2 -n postgres-db -c $root
tmux send-keys  -t db-compare-dev:2 "./setup-postgres.sh" Enter 

tmux attach -t db-compare-dev
echo "processes started"
