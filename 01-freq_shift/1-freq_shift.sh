#!/bin/sh
echo 0`tr -d ', \r\n' < input.txt` | bc
