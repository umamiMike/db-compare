tmux new-session -d -s db-compare
tmux attach-session db-compare

tmux set -g pane-border-status top
tmux set -g pane-border-format "#{pane_index} #{pane_current_command}"
tmux rename-window 'db-compare'

tmux send-keys 'cd ./server/ && go build && ./server' 'C-m'
tmux split-window -h 
tmux send-keys -t 1 'cd client' C-m 'npm run dev' 'C-m'
