#!/bin/bash
OPT=$1
help() {
    echo "Usage: ./start.sh [options]"
    echo "Options:"
    echo "  --help: show help"
    echo "  --cron: run cron job"
    echo "    > for stop the cron job, run: ./cron.sh -s"
    echo " --build: build docker image"
    exit 0
}


# Run docker compose, if option --build is passed, build image
if [ $OPT == "--build" ]; then
    echo "[docker] Building image..."
    docker-compose up --build -d
else
    echo "[docker] Starting container..."
    docker-compose up -d
fi

waitContainer() {
    echo "[test] Waiting for container to start..."
    until docker exec -i go_app go version > /dev/null 2>&1; do
        sleep 0.2
    done
    if [ $? -eq 0 ]; then
        echo "[test] - Container started!"
    else
        echo "[test] - Container failed!"
        exit 1
    fi
}

# Run all tests into container
runTests() {
    waitContainer
    echo "[test] Waiting for tests to finish..."
    until docker exec -i go_app go test ./internal/application/usecase/...; do
        sleep 0.2
    done
    if [ $? -eq 0 ]; then
        echo "[test] - Tests passed!"
    else
        echo "[test] - Tests failed!"
        exit 1
    fi
}

case $OPT in
--help) help;;
--cron) chmod +x ./cron.sh && ./cron.sh;;
*) ;;
esac

# BASEDIR=$(pwd)
# Run script to create dabase
./internal/infra/db/postgres/create-database.sh
# Run server
echo "- Waiting for postgres server to start..."
sleep 2
# execute in second plan
go run cmd/todolist/main.go