#!/bin/bash

WHITE='\033[0m'
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'

source data.sh

./make_"$RESOURCE".sh

echo -n '--- TESTING '
echo "$RESOURCE"

# Create
kubectl apply -f "$DEST"

# Wait and get ref
sleep "$WAITING"
source get_ref.sh

# Check if created
go run test_created.go
OUT="$?"

printf "    -> Creating...\n"
if [ "$OUT" -eq "0" ]
then
    printf "    ${GREEN}PASSED !!!${WHITE}\n"
else
    printf "    ${RED}FAILED TO CREATE...${WHITE}\n"
fi

# Destroy
kubectl delete -f "$DEST"
# Check if destroyed
go run test_destroyed.go
OUT="$?"

printf "    -> Destroying...\n"
if [ "$OUT" -eq "0" ]
then
    printf "    ${GREEN}PASSED !!!${WHITE}\n"
else
    printf "    ${RED}FAILED TO DESTROY...${WHITE}\n"
fi

rm "$DEST"
