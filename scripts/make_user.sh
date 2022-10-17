source config.sh

RESOURCE="user"
DEST=$PATH_DST$PRE$RESOURCE$END

echo 'User name:'
read CP_USER_NAME

echo 'Email:'
read CP_USER_EMAIL

echo 'Role:'
read CP_USER_ROLE

echo 'Organization:'
read CP_USER_ORG

cp "$EXAMPLES_PATH"example_"$RESOURCE".yaml $DEST

sed -i "s/USER/$CP_USER_NAME/g" $DEST
sed -i "s/EMAIL/$CP_USER_EMAIL/g" $DEST
sed -i "s/ROLE/$CP_USER_ROLE/g" $DEST
sed -i "s/ORGANIZATION_REF/$CP_USER_ORG/g" $DEST
