#!/bin/bash
if [ $# -eq 0 ]
  then
    echo "Usage: $0 example_path"
    exit 1
fi

LISTING=$1
SCRIPT=${LISTING}-test.txtar
txtar-c -quote -script ${LISTING}/test.txtar $LISTING >>$SCRIPT
echo -n "${LISTING}... " && testscript $SCRIPT || exit
rm $SCRIPT
