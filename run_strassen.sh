#!/bin/bash
exec 2> ./output/strassen_results.log

time ./strassen 2 4 strassen.txt
time ./strassen 2 8 strassen.txt
time ./strassen 2 16 strassen.txt
time ./strassen 2 32 strassen.txt
time ./strassen 2 64 strassen.txt
time ./strassen 2 128 strassen.txt
time ./strassen 2 256 strassen.txt
time ./strassen 2 512 strassen.txt
time ./strassen 2 1024 strassen.txt
time ./strassen 2 2048 strassen.txt
time ./strassen 2 4096 strassen.txt
time ./strassen 2 8192 strassen.txt
