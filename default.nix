{ pkgs ? import <nixpkgs> {} }:

with pkgs;

let
  inherit (lib) optional optionals;

  postgresql = postgresql_10;
  go = go_1_13;

in
mkShell {
  buildInputs = [
  go
  tmux
  nodejs
  postgresql

]
    ++ optional stdenv.isLinux inotify-tools # For file_system on Linux.
    ++ optionals stdenv.isDarwin (with darwin.apple_sdk.frameworks; [
      # For file_system on macOS.
      CoreFoundation
      CoreServices
    ]);

    # Put the PostgreSQL databases in the project diretory.
 shellHook = ''
# Place the data directory inside the project directory
      # function cleanup {
      # echo "exiting the shell and killing postgres" 
      # killall postgres
      # }


root="$(pwd)"

tmux new-session -d -s db-compare-dev
tmux set -g pane-border-status top
tmux set -g pane-border-format "#{pane_index} #{pane_current_command}"

tmux new-window -t db-compare-dev -n umbrella

tmux new-window -t db-compare-dev -n server
tmux send-keys 'cd ./db-compare-server/ && go build && ./db-compare-server' 'C-m'

tmux new-window -t db-compare-dev -n react-client
tmux send-keys  'cd react-client &&  npm run start' 'C-m'

tmux new-window -t db-compare-dev -n postgres-db
tmux send-keys  'source setup-postgres.sh' 'C-m'

tmux attach -t db-compare-dev
    '';
}
