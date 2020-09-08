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
  inotify-tools

];

buildPhase = ''

'';
 shellHook = ''
  export PATH="$PWD/node_modules/.bin/:$PATH"
  export PGDATA="$(pwd)/postgres"
# Place Postgres' Unix socket inside the data directory
  export PGHOST="$PGDATA"
glibcLocales=$(nix-build --no-out-link "<nixpkgs>" -A glibcLocales)
echo $glibcLocales
export LOCALE_ARCHIVE_2_27="${glibcLocales}/lib/locale/locale-archive"

 '';
}
