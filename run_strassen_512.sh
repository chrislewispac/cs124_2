#!/bin/bash
exec 2> ./output/strassen_results_512.log

time ./strassen 512 4 strassen.txt
time ./strassen 512 8 strassen.txt
time ./strassen 512 16 strassen.txt
time ./strassen 512 32 strassen.txt
time ./strassen 512 64 strassen.txt
time ./strassen 512 128 strassen.txt
time ./strassen 512 256 strassen.txt
time ./strassen 512 512 strassen.txt
time ./strassen 512 1024 strassen.txt
time ./strassen 512 2048 strassen.txt
time ./strassen 512 4096 strassen.txt
