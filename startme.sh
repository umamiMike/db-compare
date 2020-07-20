root="$(pwd)"

tmux new-session -d -s db-compare
tmux attach-session db-compare

tmux set -g pane-border-status top
tmux set -g pane-border-format "#{pane_index} #{pane_current_command}"
tmux rename-window 'db-compare'

tmux send-keys 'cd ./server/ && go build && ./server' 'C-m'
ls ./
tmux split-window -v 
tmux send-keys  'cd client &&  npm run dev' 'C-m'
ls ./
tmux split-window -h 
tmux send-keys 'cd ./db-compare-react-client &&  npm run start' 'C-m'
