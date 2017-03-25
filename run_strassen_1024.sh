#!/bin/bash
exec 2> ./output/strassen_results_1024.log

time ./strassen 1024 4 strassen.txt
time ./strassen 1024 8 strassen.txt
time ./strassen 1024 16 strassen.txt
time ./strassen 1024 32 strassen.txt
time ./strassen 1024 64 strassen.txt
time ./strassen 1024 128 strassen.txt
time ./strassen 1024 256 strassen.txt
time ./strassen 1024 512 strassen.txt
time ./strassen 1024 1024 strassen.txt
time ./strassen 1024 2048 strassen.txt
time ./strassen 1024 4096 strassen.txt
