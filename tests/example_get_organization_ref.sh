#!/bin/bash

WAITING=938
COUNT=0
COUNT_LIMIT=$(($WAITING / $JUMP_TIME))
REF=""
while [ "$REF" = "" ];
do
    COUNT=$(($COUNT + 1))
    if [ $COUNT -gt $COUNT_LIMIT ];
    then
        echo "TIMEOUT!"
        kubectl delete -f "$DEST"
        rm out*
        rm "$DEST"
        exit 1
    fi

    kubectl get managed > out_raw
    grep "test-$TAIKUN_USER-org-attach" "out_raw" > out
    CONTENT=`cat out`

    READY=`grep -o True out | wc -l`

    sed -i "s/.* \([0-9]\+\)\(\/[0-9]\+\)\? .*/\1/g" "out"

    CONTENT_AFTER=`cat out`

    if [ $READY -eq 2 -a "$CONTENT" != "$CONTENT_AFTER" ];
    then
        RESOURCES=`cat out | tr " " "\n"`
        for id in $RESOURCES; do
            REF="$id"
        done
    else
        REF=""
    fi

    sleep $JUMP_TIME
done

export TESTS_ORGANIZATION_REF=$REF

rm out*
