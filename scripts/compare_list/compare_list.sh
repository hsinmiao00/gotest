#!/bin/bash
COMM=`comm -12 ./l1 ./l2 | wc -l | awk '{print $1}'`
PREV=`wc -l ./l1 | awk '{print $1}'`
RATIO=`awk "BEGIN {printf \"%.2f\n\", $COMM/$PREV}"`
if  (( $(bc <<< "$RATIO < 0.8") ))
then
    echo "ERROR - comm: "$COMM", old list: "$PREV
else
    echo "list ok!!"
fi

