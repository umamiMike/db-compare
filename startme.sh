tmux new -d -s server 'exec bash'
tmux send-keys 'go run ./server/main.go ./server/data.json' C-m
#tmux select-window -t server:0
tmux split-window -h -t 0 'exec bash'
tmux send-keys -t 1 'cd client' C-m 'npm run dev' C-m
tmux attach

