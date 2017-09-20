#!/usr/local/bin/dash

stop() {
  kill -9 $(ps -e | grep [p]lay_music | awk '{ print $1 }')
  killall afplay
}

next() {
  killall afplay
}

case $1 in
  1 )
    stop
    ;;
  4 )
    next
    ;;
  * )
    ;;
esac