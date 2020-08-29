#!/bin/bash
# Place the data directory inside the project directory
      # function cleanup {
      # echo "exiting the shell and killing postgres" 
      # killall postgres
      # }

# new window 
#    -c  is the starting directory
root="$(pwd)"

tmux kill-session -t db-compare-dev
tmux new-session -d -s db-compare-dev 

tmux new-window -t db-compare-dev:2 -n server -c $root/db-compare-server
tmux send-keys -t db-compare-dev:2 './dev-build.sh' Enter

tmux new-window -t db-compare-dev:3 -n react-client -c $root/react-client
tmux send-keys -t db-compare-dev:3 'npm run start' Enter 

tmux new-window -t db-compare-dev:4 -n postgres-db -c $root
tmux send-keys  -t db-compare-dev:4 "./setup-postgres.sh" Enter 

# tmux attach -t db-compare-dev
