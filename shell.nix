with import <nixpkgs> {};

stdenv.mkDerivation {
  name = "db-compare-umbrella";

  postgresql = postgresql_10;
  PROJECT_ROOT = builtins.toString ./.;
  LOCALE_ARCHIVE_2_27 = "${glibcLocales}/lib/locale/locale-archive";

  buildInputs = with pkgs; [
  bash
  go
  goimports
  fswatch
  vgo2nix
  nodejs
  postgresql
];

shellHook = ''
  export GOPATH="/home/mike/dev/.go"
  export GOCACHE=""
  export PATH="$PROJECT_ROOT/react-client/node_modules/.bin/:$PATH"

  export PGDATA="$PROJECT_ROOT/data/postgres_data"
  export PGHOST="$PROJECT_ROOT/data/postgres"
  export LOG_PATH="$PROJECT_ROOT/data/postgres/LOG"
  export PGDATABASE=postgres
  export DATABASE_URL="postgresql?host=$PGHOST"

 '';
}
