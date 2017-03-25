#!/bin/bash
exec 2> ./output/strassen_results_8.log

time ./strassen 8 4 strassen.txt
time ./strassen 8 8 strassen.txt
time ./strassen 8 16 strassen.txt
time ./strassen 8 32 strassen.txt
time ./strassen 8 64 strassen.txt
time ./strassen 8 128 strassen.txt
time ./strassen 8 256 strassen.txt
time ./strassen 8 512 strassen.txt
time ./strassen 8 1024 strassen.txt
time ./strassen 8 2048 strassen.txt
time ./strassen 8 4096 strassen.txt
