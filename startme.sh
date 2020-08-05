root="$(pwd)"

tmux new-session -d -s db-compare
tmux attach-session db-compare

tmux set -g pane-border-status top
tmux set -g pane-border-format "#{pane_index} #{pane_current_command}"
tmux rename-window 'db-compare'
tmux new-window -n server
tmux send-keys 'cd ./db-compare-server/ && go build && ./db-compare-server' 'C-m'
tmux new-window -n react-client
tmux send-keys  'cd react-client &&  npm run start' 'C-m'
