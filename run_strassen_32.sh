#!/bin/bash
exec 2> ./output/strassen_results_32.log

time ./strassen 32 4 strassen.txt
time ./strassen 32 8 strassen.txt
time ./strassen 32 16 strassen.txt
time ./strassen 32 32 strassen.txt
time ./strassen 32 64 strassen.txt
time ./strassen 32 128 strassen.txt
time ./strassen 32 256 strassen.txt
time ./strassen 32 512 strassen.txt
time ./strassen 32 1024 strassen.txt
time ./strassen 32 2048 strassen.txt
time ./strassen 32 4096 strassen.txt
