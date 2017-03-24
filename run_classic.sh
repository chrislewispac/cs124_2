#!/bin/bash
exec 2> ./output/classic_results.log

time ./strassen 1 4 classic.txt
time ./strassen 1 8 classic.txt
time ./strassen 1 16 classic.txt
time ./strassen 1 32 classic.txt
time ./strassen 1 64 classic.txt
time ./strassen 1 128 classic.txt
time ./strassen 1 256 classic.txt
time ./strassen 1 512 classic.txt
time ./strassen 1 1024 classic.txt
time ./strassen 1 2048 classic.txt
time ./strassen 1 4096 classic.txt
time ./strassen 1 8192 classic.txt
