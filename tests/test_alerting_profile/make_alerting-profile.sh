#!/bin/bash

source data.sh
cp example_"$RESOURCE".yaml $DEST

sed -i "s/PROVIDER/$PROVIDER_NAME/g" $DEST

sed -i "s/ALERTING/$CP_ALERT_NAME/g" $DEST
sed -i "s/EMAIL/$CP_ALERT_EMAIL/g" $DEST
sed -i "s/ORGANIZATION_REF/$CP_ALERT_ORG/g" $DEST
sed -i "s/REMINDER/$CP_ALERT_REMINDER/g" $DEST
