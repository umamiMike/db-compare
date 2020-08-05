# init the db  

export PGDATA="$(pwd)/postgres"
# Place Postgres' Unix socket inside the data directory
export PGHOST="$PGDATA"

chmod +x setup-postgres.sh

cat > "$PGDATA/postgresql.conf" <<EOF
listen_addresses = 'localhost'
port = 5432
unix_socket_directories = '$PGHOST'
EOF

# echo "init the db? y " 
# read initdb
# if [  $initdb = "y" ]; then
  initdb
  # ...configure it to listen only on the Unix socket, and...
  echo "CREATE DATABASE postgres;" | postgres --single -E postgres
  echo "CREATE USER postgres --createdb;" | postgres --single -E postgres
# fi
postgres
