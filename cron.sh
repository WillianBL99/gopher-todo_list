#!/bin/bash
# default time to run the cron script
TIME=12

help() {
    echo "Usage: ./cron.sh [options]"
    echo "Options:"
    echo "  -s: stop the cron script"
    echo "  -t <time>: set the time in hours to run the cron script"
    echo "  -o: run the cron routine once"
    echo "  -h: show help"
    exit 0
}

case $1 in
-s)
    echo "Stopping cron job..."
    kill -9 $(ps aux | grep '[c]ron.sh' | awk '{print $2}')
    exit 0
    ;;
-t)
    if [ -z "$2" ]; then
        echo "Error: time is required!"
        exit 1
    fi
    TIME=$2
    echo "Time set to $TIME hours"
    ;;
-h)
    help
    ;;
*);;
esac

MIN=60
HOUR=$((60*$MIN))

# show time remaining in background for the next backup
remaining() {
    local start_time=$(date +%s)
    echo "Cron started at $(date -u -d @$start_time +%H:%M:%S)"
    echo "Press Ctrl+C to stop"
    while true; do
        local next_time=$(($start_time+$TIME*$HOUR))
        local time_format=$(date -u -d @$(($next_time-$(date +%s))) +%H:%M:%S)
        printf "\rTime remaining to next backup: $time_format"
        sleep 1
    done
}

# function to stop the cron script
cleanup() {
    echo
    echo "Stopping cron job..."
    kill $bg_pid  # send a signal to the process in the background
    wait $bg_pid  # wait for the process to actually exit
    exit 1
}

trap cleanup INT

execute() {
    chmod +x ./internal/infra/db/postgres/backup-database.sh
    ./internal/infra/db/postgres/backup-database.sh
}

if [ "$1" == "-o" ]; then
    execute
    exit 0
fi

while true; do
    #./internal/infra/db/postgres/backup-database.sh
    execute
    remaining &
    bg_pid=$!  # Salva o PID do processo em segundo plano
    sleep $(($TIME*$HOUR))
    kill $bg_pid
    wait $bg_pid
done