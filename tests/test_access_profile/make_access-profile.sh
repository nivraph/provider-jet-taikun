#!/bin/bash

source data.sh
cp example_"$RESOURCE".yaml $DEST

sed -i "s/PROVIDER/$PROVIDER_NAME/g" $DEST

sed -i "s/ACCESS_PROFILE/$CP_ACCESS_NAME/g" $DEST
sed -i "s/ORGANIZATION_REF/$CP_ACCESS_ORG/g" $DEST
