source config.sh

RESOURCE="alerting_profile"
DEST=$PATH_DST$PRE$RESOURCE$END

echo 'Alerting profile name:'
read CP_ALERT_NAME

echo 'Email:'
read CP_ALERT_EMAIL

echo 'Organization:'
read CP_ALERT_ORG

echo 'Reminder:'
read CP_ALERT_REMINDER

cp example_"$RESOURCE".yaml $DEST

sed -i "s/ALERTING/$CP_ALERT_NAME/g" $DEST
sed -i "s/EMAIL/$CP_ALERT_EMAIL/g" $DEST
sed -i "s/ORGANIZATION_REF/$CP_ALERT_ORG/g" $DEST
sed -i "s/REMINDER/$CP_ALERT_REMINDER/g" $DEST
