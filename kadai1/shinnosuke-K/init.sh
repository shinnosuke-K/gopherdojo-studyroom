#!/bin/sh

# create dir except exists
exist() {
    if [ ! -e "$1" ]; then
      mkdir "$1"
    fi
}

exist "dummyImg"
exist "dummyImg/jpg"
exist "dummyImg/jpeg"
exist "dummyImg/png"


# create sample image files for test
create() {
  for i in $(seq 1 3)
  do
    cp standard.png "$1"/"$i"."$2"
  done
}

create "dummyImg/jpg" "jpg"
create "dummyImg/jpeg" "jpeg"
create "dummyImg/png" "png"
create "dummyImg" "gif"

dupCreate() {
  for i in $(seq 1 3)
  do
    cp standard.png "$1"/"$i"."$2"."$2"
  done
}

dupCreate "dummyImg/jpg" "jpg"
dupCreate "dummyImg/jpeg" "jpeg"
dupCreate "dummyImg/png" "png"
dupCreate "dummyImg" "gif"