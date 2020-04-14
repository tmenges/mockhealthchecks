#! /bin/bash

INITIAL_DELAY="${INITIAL_DELAY:-10}"
LIVE_TIME="${LIVE_TIME:-20}"
READY_TIME="${READY_TIME:-$LIVE_TIME}"

echo "Starting process"
sleep $INITIAL_DELAY

echo "Running webserver"
webserver &
webserver_pid=$!

sleep $LIVE_TIME

echo "Killing webserver"
kill $webserver_pid

echo "Done"
exit 0