#!/bin/bash

source_folder="ex11"

for i in {12..26}; do
    destination_folder="ex$i"
    cp -r "$source_folder"/* "$destination_folder"
done
