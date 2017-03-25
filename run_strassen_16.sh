#!/bin/bash
exec 2> ./output/strassen_results_16.log

time ./strassen 16 4 strassen.txt
time ./strassen 16 8 strassen.txt
time ./strassen 16 16 strassen.txt
time ./strassen 16 32 strassen.txt
time ./strassen 16 64 strassen.txt
time ./strassen 16 128 strassen.txt
time ./strassen 16 256 strassen.txt
time ./strassen 16 512 strassen.txt
time ./strassen 16 1024 strassen.txt
time ./strassen 16 2048 strassen.txt
time ./strassen 16 4096 strassen.txt
