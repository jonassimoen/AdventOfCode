#!/bin/bash

i=$1
if [[ $i -lt 10 ]]; then
  i="0$i"
fi


if [[ -z $2 ]]; then
  time go run day${i}/main.go -i day${i}/test_in
  echo "--"
  time go run day${i}/main.go -i day${i}/in
else
  time go run day${i}/main.go -i day${i}/$2
fi
