#!/bin/bash
exec 2> ./output/strassen_results_128.log

time ./strassen 128 4 strassen.txt
time ./strassen 128 8 strassen.txt
time ./strassen 128 16 strassen.txt
time ./strassen 128 32 strassen.txt
time ./strassen 128 64 strassen.txt
time ./strassen 128 128 strassen.txt
time ./strassen 128 256 strassen.txt
time ./strassen 128 512 strassen.txt
time ./strassen 128 1024 strassen.txt
time ./strassen 128 2048 strassen.txt
time ./strassen 128 4096 strassen.txt
