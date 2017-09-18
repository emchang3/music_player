#!/usr/local/bin/dash

stop() {
  kill -9 $(ps -e | grep [p]lay_music | awk '{ print $1 }')
  killall afplay
}

cont() {
  killall -19 afplay
}

pause() {
  killall -17 afplay
}

next() {
  killall afplay
}

case $1 in
  1 )
    stop
    ;;
  2 )
    cont
    ;;
  3 )
    pause
    ;;
  4 )
    next
    ;;
  * )
    ;;
esac