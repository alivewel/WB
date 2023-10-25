#!/bin/bash

for i in {11..26}; do
    folder_name="ex$i"
    cd "$folder_name" || exit
    go mod init "$folder_name"
    cd ..
done
