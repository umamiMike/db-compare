# init the db with 
if [[ ! -d "$PGDATA" ]]; then
	# If the data directory doesn't exist, create an empty one, and...
	initdb
	# ...configure it to listen only on the Unix socket, and...
	cat >> "$PGDATA/postgresql.conf" <<-EOF
		listen_addresses = 'localhost'
                port = 5432
		unix_socket_directories = '$PGHOST'
	EOF
	# ...create a database using the name Postgres defaults to.
	echo "CREATE DATABASE postgres;" | postgres --single -E postgres
	echo "CREATE USER postgres --createdb;" | postgres --single -E postgres

fi
postgres &

