#!/bin/bash

WHITE='\033[0m'
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'

source data.sh

./make_"$RESOURCE".sh

printf "\n=== TESTING "
printf "$RESOURCE\n"

# Create
kubectl apply -f "$DEST"
printf "    -> Creating...\n"

# Wait and get ref
source get_ref.sh

# Check if created
go run test_created.go
OUT="$?"

if [ "$OUT" -eq "0" ]
then
    printf "    ${GREEN}PASSED !!!${WHITE}\n"
else
    printf "    ${RED}FAILED TO CREATE...${WHITE}\n"
fi

# Destroy
kubectl delete -f "$DEST"
printf "    -> Destroying...\n"

# Check if destroyed
go run test_destroyed.go
OUT="$?"

if [ "$OUT" -eq "0" ]
then
    printf "    ${GREEN}PASSED !!!${WHITE}\n"
else
    printf "    ${RED}FAILED TO DESTROY...${WHITE}\n"
fi

printf "\n"

rm "$DEST"
