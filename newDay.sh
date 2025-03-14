#!/bin/bash

INPUT_TEMPLATE="package %s

import _ \"embed\"

//go:embed test-input
var TestInput string

//go:embed input
var Input string"

DAY_TEMPLATE="package %s

func Solution(input string){}
"

if [[ "$#" -ne 1 ]]; then
    echo "need day number"
    exit
fi

if [[ "$1" -lt 1 ]]; then
    echo "want day number > 0"
    exit
fi

echo 

num=$(printf "%02d" $1)
day="day$num"

if [[ ! -d "./$day" ]]; then
    echo "making dir"
    mkdir "./$day"
fi

if [[ ! -e "./$day/$day.go" ]]; then
    echo "making go files"
    printf "$INPUT_TEMPLATE" "$day" > "./$day/input.go"
    printf "$DAY_TEMPLATE" "$day" > "./$day/$day.go"
fi