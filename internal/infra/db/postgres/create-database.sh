#!/bin/bash
# wait for postgres server to start
# read variables from .env file
export $(grep -v '^#' .env | xargs)
if [ $DB_NAME = "" ] || [ $DB_USER = "" ] || [ $DB_PASSWORD = "" ]; then
    echo "Database variables not found!"
    printenv | grep DB
    exit 1
fi

# verify if container is running
until docker exec pg_db psql -U $DB_USER -c '\q' > /dev/null 2>&1; do 
    echo "Waiting for postgres server to start..."
    sleep 2
done
# verify if has backup
if [ -f ./internal/infra/db/postgres/backup.sql ]; then
    echo "Restoring database..."
    docker exec -i pg_db psql -U $DB_USER -d $DB_NAME < ./internal/infra/db/postgres/backup.sql
    if [ $? -eq 0 ]; then
        echo "Database restored!"
    else
        echo "Database restore failed!"
        exit 1
    fi
else
    echo "Creating database..."
    docker exec -i pg_db psql -U $DB_USER -d $DB_NAME < ./internal/infra/db/postgres/create-tables.sql
    if [ $? -eq 0 ]; then
        echo "Database created!"
    else
        echo "Database creation failed!"
        exit 1
    fi
fi