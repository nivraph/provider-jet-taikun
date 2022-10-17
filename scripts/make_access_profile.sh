source config.sh

RESOURCE="access_profile"
DEST=$PATH_DST$PRE$RESOURCE$END

echo 'Access profile name:'
read CP_ACCESS_NAME

echo 'Organization:'
read CP_ACCESS_ORG

cp "$EXAMPLES_PATH"example_"$RESOURCE".yaml $DEST

sed -i "s/ACCESS_PROFILE/$CP_ACCESS_NAME/g" $DEST
sed -i "s/ORGANIZATION_REF/$CP_ACCESS_ORG/g" $DEST
