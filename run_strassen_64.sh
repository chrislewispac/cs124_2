#!/bin/bash
exec 2> ./output/strassen_results_64.log

time ./strassen 64 4 strassen.txt
time ./strassen 64 8 strassen.txt
time ./strassen 64 16 strassen.txt
time ./strassen 64 32 strassen.txt
time ./strassen 64 64 strassen.txt
time ./strassen 64 128 strassen.txt
time ./strassen 64 256 strassen.txt
time ./strassen 64 512 strassen.txt
time ./strassen 64 1024 strassen.txt
time ./strassen 64 2048 strassen.txt
time ./strassen 64 4096 strassen.txt
