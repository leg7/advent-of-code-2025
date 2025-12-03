#!/bin/sh

for i in $(seq 12); do
	if test -d "$i"; then
		printf "\nDAY $i\n\n"
		cd "$i" || exit

		printf "PART 1\n\n"
		go run ./1 < example.txt
		printf "\n"

		printf "PART 2\n\n"
		go run ./2 < example.txt
		printf "\n"

		cd - || exit
	fi
done
