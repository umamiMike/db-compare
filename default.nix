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
export PGDATA="$dev/postgres"
# Place Postgres' Unix socket inside the data directory
export PGHOST="$PGDATA"
# create if doesnt exist

      
      function cleanup {
      echo "exiting the shell and killing postgres" 
      killall postgres
      }

      # echo  -e "start db?" 
      # read startdb
      # if [  $startdb = "y" ]; then
      # source ./setup-postgres.sh
      ./startme.sh
      # fi

    '';
}
