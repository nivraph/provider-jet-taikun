#!/bin/bash

source data.sh
cp example_"$RESOURCE".yaml $DEST

sed -i "s/PROVIDER/$PROVIDER_NAME/g" $DEST

sed -i "s/KUBERNETES_PROFILE/$CP_KP_NAME/g" $DEST
sed -i "s/ORGANIZATION_REF/$CP_KP_ORG/g" $DEST
