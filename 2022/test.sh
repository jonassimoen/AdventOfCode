#!/bin/bash

i=$1
if [[ $i -lt 10 ]]; then
  i="0$i"
fi

p=$2
if [[ -n $3 ]]; then
  f="test_in"
else
  echo "Final puzzle"
  f="in"
fi

time go run day${i}/main.go -p ${p} -i day${i}/${f}