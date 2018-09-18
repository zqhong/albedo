#!/bin/bash

SERVER="albedo"
BASE_DIR=$PWD/build
INTERVAL=2

ARGS=""

function start()
{
	if [ "`pgrep $SERVER -u $UID`" != "" ];then
		echo "$SERVER already running"
		exit 1
	fi

	nohup $BASE_DIR/$SERVER $ARGS  &>/dev/null &

	echo "sleeping..." &&  sleep $INTERVAL

	# check status
	if [ "`pgrep $SERVER -u $UID`" == "" ];then
		echo "$SERVER start failed"
		exit 1
	fi
}

function status()
{
	if [ "`pgrep $SERVER -u $UID`" != "" ];then
		echo $SERVER is running
	else
		echo $SERVER is not running
	fi
}

function stop()
{
	if [ "`pgrep $SERVER -u $UID`" != "" ];then
		kill -9 `pgrep $SERVER -u $UID`
	fi

	echo "sleeping..." &&  sleep $INTERVAL

	if [ "`pgrep $SERVER -u $UID`" != "" ];then
		echo "$SERVER stop failed"
		exit 1
	fi
}

function restart()
{
	if [ "`pgrep $SERVER -u $UID`" != "" ];then
		kill -s HUP `pgrep $SERVER -u $UID`
		echo "$SERVER restart successful"
    else
        echo "$SERVER is not running, restart fail"
        exit 1
	fi
}

case "$1" in
	'start')
	start
	;;
	'stop')
	stop
	;;
	'status')
	status
	;;
	'restart')
	restart
	;;
	*)
	echo "usage: $0 {start|stop|restart|status}"
	exit 1
	;;
esac