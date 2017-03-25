#!/bin/bash
exec 2> ./output/strassen_results_256.log

time ./strassen 256 4 strassen.txt
time ./strassen 256 8 strassen.txt
time ./strassen 256 16 strassen.txt
time ./strassen 256 32 strassen.txt
time ./strassen 256 64 strassen.txt
time ./strassen 256 128 strassen.txt
time ./strassen 256 256 strassen.txt
time ./strassen 256 512 strassen.txt
time ./strassen 256 1024 strassen.txt
time ./strassen 256 2048 strassen.txt
time ./strassen 256 4096 strassen.txt
