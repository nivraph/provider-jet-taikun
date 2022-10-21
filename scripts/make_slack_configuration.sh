source config.sh

RESOURCE="slack_configuration"
DEST=$PATH_DST$PRE$RESOURCE$END

echo 'Slack configuration name:'
read CP_SC_NAME

echo 'Organization:'
read CP_SC_ORG

echo 'Channel:'
read CP_SC_CHANNEL

echo 'Type:'
read CP_SC_TYPE

echo 'Url:'
read CP_SC_URL

cp "$EXAMPLES_PATH"example_"$RESOURCE".yaml $DEST

sed -i "s/SLACK_CONFIGURATION/$CP_SC_NAME/g" $DEST
sed -i "s/ORGANIZATION_REF/$CP_SC_ORG/g" $DEST
sed -i "s/CHANNEL/$CP_SC_CHANNEL/g" $DEST
sed -i "s/TYPE/$CP_SC_TYPE/g" $DEST
sed -i "s^URL^$CP_SC_URL^g" $DEST
