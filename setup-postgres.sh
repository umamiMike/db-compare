# make the data directory if needed 
mkdir -p postgres
# necessary env vars
export PGDATA="$(pwd)/postgres"
# Place Postgres' Unix socket inside the data directory
export PGHOST="$PGDATA"

# initialize the db
initdb

# add the appropriate configuration file
cat > "$PGDATA/postgresql.conf" <<EOF
listen_addresses = 'localhost'
port = 5432
unix_socket_directories = '$PGHOST'
EOF

# create the base db
echo "CREATE DATABASE postgres;" | postgres --single -E postgres
# create the user
echo "CREATE USER postgres --createdb;" | postgres --single -E postgres

# start the instance
postgres
