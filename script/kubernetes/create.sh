#!/bin/bash
BASE_DIR=`pwd`

# must create configmap first
dirs=(./configmap ./deployment ./service)
for dir in "${dirs[@]}"
do
    cd $dir
    for filename in ./*.yaml; do
        kubectl create -f $filename
    done

    cd $BASE_DIR
done