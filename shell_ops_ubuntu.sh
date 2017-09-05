#!/bin/dash

destroy() {
  gopid=$(pgrep "alarm")
  gostatus=$?
  if [ $gostatus = "0" ]
  then
    kill $gopid
    echo $gopid
  else
    echo "NOGO"
  fi
}

build() {
  go build
  echo $?
}

start() {
  ./alarm
}

started() {
  gopid=$(pgrep "alarm")
  gostatus=$?
  if [ $gostatus = "0" ]
  then
    echo $gopid
  else
    echo "NOGO"
  fi
}

case $1 in
  1 )
    destroy
    ;;
  2 )
    build
    ;;
  3 )
    start
    ;;
  4 )
    started
    ;;
  * )
    ;;
esac
