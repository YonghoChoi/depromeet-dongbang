#!/bin/bash
BASE_DIR=`pwd`

dirs=(./service ./deployment ./configmap)
for dir in "${dirs[@]}"
do
    cd $dir
    for filename in ./*.yaml; do
        kubectl delete -f $filename
    done

    cd $BASE_DIR
done
