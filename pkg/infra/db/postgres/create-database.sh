#!/bin/bash
# wait for postgres server to start
until docker exec pg_db psql -U postgres -c '\q' > /dev/null 2>&1; do 
    echo "Waiting for postgres server to start..."
    sleep 2
done
# verify if has backup
if [ -f ./pkg/infra/db/postgres/backup.sql ]; then
    echo "Restoring database..."
    docker exec -i pg_db psql -U postgres -d postgres < ./pkg/infra/db/postgres/backup.sql
    if [ $? -eq 0 ]; then
        echo "Database restored!"
    else
        echo "Database restore failed!"
        exit 1
    fi
else
    echo "Creating database..."
    docker exec -i pg_db psql -U postgres -d postgres < ./pkg/infra/db/postgres/create-tables.sql
    if [ $? -eq 0 ]; then
        echo "Database created!"
    else
        echo "Database creation failed!"
        exit 1
    fi
fi