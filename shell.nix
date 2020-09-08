{ pkgs ? import <nixpkgs> {} }:

with pkgs;

let
  postgresql = postgresql_10;

in
mkShell {
  # LOCALE_ARCHIVE_2_27 = "${glibcLocales}/lib/locale/locale-archive"; 

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

buildPhase = ''

'';
 shellHook = ''
  glibcLocales=$(nix-build --no-out-link "<nixpkgs>" -A glibcLocales)
  export LOCALE_ARCHIVE_2_27="${glibcLocales}/lib/locale/locale-archive"

  export GOPATH="$(pwd)/.go"
  exportGOCACHE=""
  export PATH="$(pwd)/react-client/node_modules/.bin/:$PATH"
  export PGDATA="$(pwd)/postgres"
  # Place Postgres' Unix socket inside the data directory
  export PGHOST="$PGDATA"

 '';
}
