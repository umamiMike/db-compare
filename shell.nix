{ pkgs ? import <nixpkgs> {} }:

with pkgs;

let
  postgresql = postgresql_10;

in
mkShell {

  buildInputs = [
  bash
  go
  goimports
  fswatch
  vgo2nix
  tmux
  nodejs
  postgresql
];

shellHook = ''
  export GOPATH="$(pwd)/.go"
  exportGOCACHE=""
  export PATH="$(pwd)/react-client/node_modules/.bin/:$PATH"
  export PGDATA="$(pwd)/postgres"
  export PGHOST="$PGDATA"
  zsh

 '';
}
