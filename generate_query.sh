#!/bin/bash
echo "start generate query code"
echo "build generator"

# build 
go build -o generate_query ./src/database/define
echo "build success"

# clean
echo "start clean up query code"
rm -f ./src/model/*.gen.go
rm -f ./src/query/gen*.go
rm -f ./src/query/*.gen.go

# generate
./generate_query

# clean up
rm -f ./generate_query
rm -f ./MoneyManager