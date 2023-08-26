#!/bin/bash

backup() {
    echo "Curren directory: $(pwd)"
    echo "Backup database..."
    TIMESTAMP=$(date +%Y%m%d%H%M%S)
    BACKUP_DIR="./internal/infra/db/postgres"
    BACKUP_FILE="$BACKUP_DIR/backup.sql"

    # verify if has postgres container
    if [ ! "$(docker exec -it pg_db psql -U postgres -c '\q' > /dev/null 2>&1)" ]; then
        echo "Waiting for postgres server to start..."
        sleep 2
    fi

    # wait database backup
    until docker exec -it pg_db pg_dump -U postgres todo_list > $BACKUP_FILE; do
        printf "."
        sleep 2
    done
    if [ $? -eq 0 ]; then
        echo "Backup completed!"
    else
        echo "Backup failed!"
        exit 1
    fi
}

# Run backup
backup
