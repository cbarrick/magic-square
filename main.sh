#!/bin/bash

count=10
timeout=60

echo "timeout,$timeout"
echo "count,$count"

echo "Generator,Order,Difficulty,GA_time,CSP_time"

echo -n "siamese1,3,easy,"
echo -n `go run ./ga_trial.go -g siamese1 -o 3 -d easy -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g siamese1 -o 3 -d easy -n $count -t $timeout`
echo

echo -n "siamese1,3,med,"
echo -n `go run ./ga_trial.go -g siamese1 -o 3 -d med -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g siamese1 -o 3 -d med -n $count -t $timeout`
echo
echo -n "siamese1,3,hard,"
echo -n `go run ./ga_trial.go -g siamese1 -o 3 -d hard -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g siamese1 -o 3 -d hard -n $count -t $timeout`
echo

echo -n "siamese1,5,easy,"
echo -n `go run ./ga_trial.go -g siamese1 -o 5 -d easy -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g siamese1 -o 5 -d easy -n $count -t $timeout`
echo
echo -n "siamese1,5,med,"
echo -n `go run ./ga_trial.go -g siamese1 -o 5 -d med -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g siamese1 -o 5 -d med -n $count -t $timeout`
echo
echo -n "siamese1,5,hard,"
echo -n `go run ./ga_trial.go -g siamese1 -o 5 -d hard -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g siamese1 -o 5 -d hard -n $count -t $timeout`
echo

echo -n "siamese1,7,easy,"
echo -n `go run ./ga_trial.go -g siamese1 -o 7 -d easy -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g siamese1 -o 7 -d easy -n $count -t $timeout`
echo
echo -n "siamese1,7,med,"
echo -n `go run ./ga_trial.go -g siamese1 -o 7 -d med -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g siamese1 -o 7 -d med -n $count -t $timeout`
echo
echo -n "siamese1,7,hard,"
echo -n `go run ./ga_trial.go -g siamese1 -o 7 -d hard -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g siamese1 -o 7 -d hard -n $count -t $timeout`
echo

echo -n "siamese2,3,easy,"
echo -n `go run ./ga_trial.go -g siamese2 -o 3 -d easy -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g siamese2 -o 3 -d easy -n $count -t $timeout`
echo
echo -n "siamese2,3,med,"
echo -n `go run ./ga_trial.go -g siamese2 -o 3 -d med -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g siamese2 -o 3 -d med -n $count -t $timeout`
echo
echo -n "siamese2,3,hard,"
echo -n `go run ./ga_trial.go -g siamese2 -o 3 -d hard -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g siamese2 -o 3 -d hard -n $count -t $timeout`
echo

echo -n "siamese2,5,easy,"
echo -n `go run ./ga_trial.go -g siamese2 -o 5 -d easy -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g siamese2 -o 5 -d easy -n $count -t $timeout`
echo
echo -n "siamese2,5,med,"
echo -n `go run ./ga_trial.go -g siamese2 -o 5 -d med -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g siamese2 -o 5 -d med -n $count -t $timeout`
echo
echo -n "siamese2,5,hard,"
echo -n `go run ./ga_trial.go -g siamese2 -o 5 -d hard -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g siamese2 -o 5 -d hard -n $count -t $timeout`
echo

echo -n "siamese2,7,easy,"
echo -n `go run ./ga_trial.go -g siamese2 -o 7 -d easy -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g siamese2 -o 7 -d easy -n $count -t $timeout`
echo
echo -n "siamese2,7,med,"
echo -n `go run ./ga_trial.go -g siamese2 -o 7 -d med -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g siamese2 -o 7 -d med -n $count -t $timeout`
echo
echo -n "siamese2,7,hard,"
echo -n `go run ./ga_trial.go -g siamese2 -o 7 -d hard -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g siamese2 -o 7 -d hard -n $count -t $timeout`
echo

echo -n "king,5,easy,"
echo -n `go run ./ga_trial.go -g king -o 5 -d easy -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g king -o 5 -d easy -n $count -t $timeout`
echo
echo -n "king,5,med,"
echo -n `go run ./ga_trial.go -g king -o 5 -d med -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g king -o 5 -d med -n $count -t $timeout`
echo
echo -n "king,5,hard,"
echo -n `go run ./ga_trial.go -g king -o 5 -d hard -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g king -o 5 -d hard -n $count -t $timeout`
echo

echo -n "king,7,easy,"
echo -n `go run ./ga_trial.go -g king -o 7 -d easy -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g king -o 7 -d easy -n $count -t $timeout`
echo
echo -n "king,7,med,"
echo -n `go run ./ga_trial.go -g king -o 7 -d med -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g king -o 7 -d med -n $count -t $timeout`
echo
echo -n "king,7,hard,"
echo -n `go run ./ga_trial.go -g king -o 7 -d hard -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g king -o 7 -d hard -n $count -t $timeout`
echo

echo -n "fours,4,easy,"
echo -n `go run ./ga_trial.go -g fours -o 4 -d easy -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g fours -o 4 -d easy -n $count -t $timeout`
echo
echo -n "fours,4,med,"
echo -n `go run ./ga_trial.go -g fours -o 4 -d med -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g fours -o 4 -d med -n $count -t $timeout`
echo
echo -n "fours,4,hard,"
echo -n `go run ./ga_trial.go -g fours -o 4 -d hard -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g fours -o 4 -d hard -n $count -t $timeout`
echo

echo -n "fours,8,easy,"
echo -n `go run ./ga_trial.go -g fours -o 8 -d easy -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g fours -o 8 -d easy -n $count -t $timeout`
echo
echo -n "fours,8,med,"
echo -n `go run ./ga_trial.go -g fours -o 8 -d med -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g fours -o 8 -d med -n $count -t $timeout`
echo
echo -n "fours,8,hard,"
echo -n `go run ./ga_trial.go -g fours -o 8 -d hard -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g fours -o 8 -d hard -n $count -t $timeout`
echo

echo -n "lux,6,easy,"
echo -n `go run ./ga_trial.go -g lux -o 6 -d easy -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g lux -o 6 -d easy -n $count -t $timeout`
echo
echo -n "lux,6,med,"
echo -n `go run ./ga_trial.go -g lux -o 6 -d med -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g lux -o 6 -d med -n $count -t $timeout`
echo
echo -n "lux,6,hard,"
echo -n `go run ./ga_trial.go -g lux -o 6 -d hard -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g lux -o 6 -d hard -n $count -t $timeout`
echo

echo -n "lux,10,easy,"
echo -n `go run ./ga_trial.go -g lux -o 10 -d easy -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g lux -o 10 -d easy -n $count -t $timeout`
echo
echo -n "lux,10,med,"
echo -n `go run ./ga_trial.go -g lux -o 10 -d med -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g lux -o 10 -d med -n $count -t $timeout`
echo
echo -n "lux,10,hard,"
echo -n `go run ./ga_trial.go -g lux -o 10 -d hard -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g lux -o 10 -d hard -n $count -t $timeout`
echo

echo -n "generate,3,hard,"
echo -n `go run ./ga_trial.go -g generate -o 3 -d hard -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g generate -o 3 -d hard -n $count -t $timeout`
echo
echo -n "generate,4,hard,"
echo -n `go run ./ga_trial.go -g generate -o 4 -d hard -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g generate -o 4 -d hard -n $count -t $timeout`
echo
echo -n "generate,5,hard,"
echo -n `go run ./ga_trial.go -g generate -o 5 -d hard -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g generate -o 5 -d hard -n $count -t $timeout`
echo
echo -n "generate,6,hard,"
echo -n `go run ./ga_trial.go -g generate -o 6 -d hard -n $count -t $timeout`
echo -n ,
echo -n `./csp_trial.pl -g generate -o 6 -d hard -n $count -t $timeout`
echo
