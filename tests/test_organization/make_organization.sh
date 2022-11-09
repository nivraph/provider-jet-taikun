#!/bin/bash

source data.sh
cp example_"$RESOURCE".yaml $DEST

sed -i "s/ORGANIZATION/$CP_ORG_NAME/g" $DEST
sed -i "s/DESCRIPTION/$CP_ORG_DESCRIPTION/g" $DEST
sed -i "s/DISCOUNT/$CP_ORG_DISCOUNT/g" $DEST
sed -i "s/CITY/$CP_ORG_CITY/g" $DEST
sed -i "s/BILLING_EMAIL/$CP_ORG_BEMAIL/g" $DEST
sed -i "s/EMAIL/$CP_ORG_EMAIL/g" $DEST
sed -i "s/PHONE/$CP_ORG_PHONE/g" $DEST
