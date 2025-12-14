#!/bin/bash

for i in {01..12}; do
  if [[ $i -lt 10 ]]; then
    i="0$i"
  fi
  if ! [[ -d day${i} ]]; then
    mkdir day${i}
    touch day${i}/in day${i}/test_in
    cp template.go day${i}/main.go
    echo ${i}
  fi
done