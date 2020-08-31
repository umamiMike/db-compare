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
  inotify-tools

]
    ++ optional stdenv.isLinux inotify-tools # For file_system on Linux.
    ++ optionals stdenv.isDarwin (with darwin.apple_sdk.frameworks; [
      # For file_system on macOS.
      CoreFoundation
      CoreServices
    ]);

 shellHook = ''
    export PGDATA="$(pwd)/postgres"
    # Place Postgres' Unix socket inside the data directory
    export PGHOST="$PGDATA"
    echo $(pwd) entered & shell loaded
 '';
}
