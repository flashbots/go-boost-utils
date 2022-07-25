#!/bin/bash
set -euo pipefail

# By default, fuzz each test for one minute.
FUZZTIME=${FUZZTIME:-1m}

files=$(grep --recursive --include='**_test.go' --files-with-matches 'func Fuzz' .)

bold=$(tput bold)
normal=$(tput sgr0)

for file in $files; do
    funcs=$(grep -o 'func Fuzz\w*' $file | awk '{print $2}')
    for func in $funcs; do
        echo "${bold}[+] fuzzing $func in $file for ${FUZZTIME}${normal}"
        parent=$(dirname $file)
        go test $parent -run=$func\$ -fuzz=$func\$ -fuzztime=${FUZZTIME}
    done
done
