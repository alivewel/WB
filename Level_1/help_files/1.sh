#!/bin/bash

# создание папок для заданий

for i in {11..26}; do
    folder_name="ex$i"
    mkdir "$folder_name"
done
