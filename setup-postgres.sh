#!/bin/env nix-shell
# inits if non existant and starts the process, stops on shell exit
   if [[ ! -d $PGHOST ]]; then
   echo "creating pghost directory"
      mkdir -p $PGHOST
    fi

    if [[ ! -d $PGDATA ]]; then
      echo 'Initializing postgresql database...'
      initdb $PGDATA --auth=trust >/dev/null
    fi

pg_ctl start -l $LOG_PATH -o "-c listen_addresses= -c unix_socket_directories=$PGHOST"

 trap "pg_ctl stop" EXIT
