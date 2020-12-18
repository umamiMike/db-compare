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
  export GOBIN="$GOPATh/.go/bin"
  export PATH=$GOPATh/.go/bin:$PATH
  exportGOCACHE=""
  export PATH="$(pwd)/react-client/node_modules/.bin/:$PATH"
  export PGDATA="$(pwd)/postgres"
  export PGHOST="$PGDATA"
  zsh

 '';
}
