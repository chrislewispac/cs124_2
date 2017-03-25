#!/bin/bash
exec 2> ./output/strassen_results_4.log

time ./strassen 4 4 strassen.txt
time ./strassen 4 8 strassen.txt
time ./strassen 4 16 strassen.txt
time ./strassen 4 32 strassen.txt
time ./strassen 4 64 strassen.txt
time ./strassen 4 128 strassen.txt
time ./strassen 4 256 strassen.txt
time ./strassen 4 512 strassen.txt
time ./strassen 4 1024 strassen.txt
time ./strassen 4 2048 strassen.txt
time ./strassen 4 4096 strassen.txt
