#!/bin/bash

for LISTING in */; do
    ./run_testscript.sh $LISTING
done