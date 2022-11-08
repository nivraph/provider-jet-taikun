kubectl get managed > out_raw
grep "$RESOURCE_NAME" "out_raw" > out

sed -i "s/.* \([0-9]\+\)\(\/[0-9]\+\)\? .*/\1/g" "out"

REF=`cat out`
export RESOURCE_REF=$REF

rm out*
