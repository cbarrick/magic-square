#!/bin/sh

echo "Generator,Order,Difficulty,GA_time,CSP_time"

printf "siamese1,3,easy,%f,%f\n" `go run ./ga_trial.go -g siamese1 -o 3 -d easy` `./csp_trial.pl -g siamese1 -o 3 -d easy`
printf "siamese1,3,med,%f,%f\n" `go run ./ga_trial.go -g siamese1 -o 3 -d med` `./csp_trial.pl -g siamese1 -o 3 -d med`
printf "siamese1,3,hard,%f,%f\n" `go run ./ga_trial.go -g siamese1 -o 3 -d hard` `./csp_trial.pl -g siamese1 -o 3 -d hard`

printf "siamese1,5,easy,%f,%f\n" `go run ./ga_trial.go -g siamese1 -o 5 -d easy` `./csp_trial.pl -g siamese1 -o 5 -d easy`
printf "siamese1,5,med,%f,%f\n" `go run ./ga_trial.go -g siamese1 -o 5 -d med` `./csp_trial.pl -g siamese1 -o 5 -d med`
printf "siamese1,5,hard,%f,%f\n" `go run ./ga_trial.go -g siamese1 -o 5 -d hard` `./csp_trial.pl -g siamese1 -o 5 -d hard`

printf "siamese1,7,easy,%f,%f\n" `go run ./ga_trial.go -g siamese1 -o 7 -d easy` `./csp_trial.pl -g siamese1 -o 7 -d easy`
printf "siamese1,7,med,%f,%f\n" `go run ./ga_trial.go -g siamese1 -o 7 -d med` `./csp_trial.pl -g siamese1 -o 7 -d med`
printf "siamese1,7,hard,%f,%f\n" `go run ./ga_trial.go -g siamese1 -o 7 -d hard` `./csp_trial.pl -g siamese1 -o 7 -d hard`

printf "siamese2,3,easy,%f,%f\n" `go run ./ga_trial.go -g siamese2 -o 3 -d easy` `./csp_trial.pl -g siamese2 -o 3 -d easy`
printf "siamese2,3,med,%f,%f\n" `go run ./ga_trial.go -g siamese2 -o 3 -d med` `./csp_trial.pl -g siamese2 -o 3 -d med`
printf "siamese2,3,hard,%f,%f\n" `go run ./ga_trial.go -g siamese2 -o 3 -d hard` `./csp_trial.pl -g siamese2 -o 3 -d hard`

printf "siamese2,5,easy,%f,%f\n" `go run ./ga_trial.go -g siamese2 -o 5 -d easy` `./csp_trial.pl -g siamese2 -o 5 -d easy`
printf "siamese2,5,med,%f,%f\n" `go run ./ga_trial.go -g siamese2 -o 5 -d med` `./csp_trial.pl -g siamese2 -o 5 -d med`
printf "siamese2,5,hard,%f,%f\n" `go run ./ga_trial.go -g siamese2 -o 5 -d hard` `./csp_trial.pl -g siamese2 -o 5 -d hard`

printf "siamese2,7,easy,%f,%f\n" `go run ./ga_trial.go -g siamese2 -o 7 -d easy` `./csp_trial.pl -g siamese2 -o 7 -d easy`
printf "siamese2,7,med,%f,%f\n" `go run ./ga_trial.go -g siamese2 -o 7 -d med` `./csp_trial.pl -g siamese2 -o 7 -d med`
printf "siamese2,7,hard,%f,%f\n" `go run ./ga_trial.go -g siamese2 -o 7 -d hard` `./csp_trial.pl -g siamese2 -o 7 -d hard`

printf "fours,4,easy,%f,%f\n" `go run ./ga_trial.go -g fours -o 4 -d easy` `./csp_trial.pl -g fours -o 4 -d easy`
printf "fours,4,med,%f,%f\n" `go run ./ga_trial.go -g fours -o 4 -d med` `./csp_trial.pl -g fours -o 4 -d med`
printf "fours,4,hard,%f,%f\n" `go run ./ga_trial.go -g fours -o 4 -d hard` `./csp_trial.pl -g fours -o 4 -d hard`

printf "fours,8,easy,%f,%f\n" `go run ./ga_trial.go -g fours -o 8 -d easy` `./csp_trial.pl -g fours -o 8 -d easy`
printf "fours,8,med,%f,%f\n" `go run ./ga_trial.go -g fours -o 8 -d med` `./csp_trial.pl -g fours -o 8 -d med`
printf "fours,8,hard,%f,%f\n" `go run ./ga_trial.go -g fours -o 8 -d hard` `./csp_trial.pl -g fours -o 8 -d hard`

printf "lux,6,easy,%f,%f\n" `go run ./ga_trial.go -g lux -o 6 -d easy` `./csp_trial.pl -g lux -o 6 -d easy`
printf "lux,6,med,%f,%f\n" `go run ./ga_trial.go -g lux -o 6 -d med` `./csp_trial.pl -g lux -o 6 -d med`
printf "lux,6,hard,%f,%f\n" `go run ./ga_trial.go -g lux -o 6 -d hard` `./csp_trial.pl -g lux -o 6 -d hard`

printf "lux,10,easy,%f,%f\n" `go run ./ga_trial.go -g lux -o 10 -d easy` `./csp_trial.pl -g lux -o 10 -d easy`
printf "lux,10,med,%f,%f\n" `go run ./ga_trial.go -g lux -o 10 -d med` `./csp_trial.pl -g lux -o 10 -d med`
printf "lux,10,hard,%f,%f\n" `go run ./ga_trial.go -g lux -o 10 -d hard` `./csp_trial.pl -g lux -o 10 -d hard`
