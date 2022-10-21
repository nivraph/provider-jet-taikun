source config.sh

RESOURCE="standalone_profile"
DEST=$PATH_DST$PRE$RESOURCE$END

echo 'Standalone profile name:'
read CP_SA_NAME

echo 'Public key:'
read CP_SA_PUBLIC_KEY

echo "$CP_SA_PUBLIC_KEY"

echo 'Organization:'
read CP_SA_ORG

cp "$EXAMPLES_PATH"example_"$RESOURCE".yaml $DEST

sed -i "s/STANDALONE/$CP_SA_NAME/g" $DEST
sed -i "s^PUBLIC_KEY^$CP_SA_PUBLIC_KEY^g" $DEST
sed -i "s/ORGANIZATION_REF/$CP_SA_ORG/g" $DEST
