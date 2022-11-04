source config.sh

RESOURCE="project_user_attachment"
DEST=$PATH_DST$PRE$RESOURCE$END

echo 'Project user attachment name:'
read CP_PROJUSER_NAME

echo 'Project:'
read CP_PROJUSER_PROJ

echo 'User:'
read CP_PROJUSER_USER

cp "$EXAMPLES_PATH"example_"$RESOURCE".yaml $DEST

sed -i "s^USER_ATTACH^$CP_PROJUSER_NAME^g" $DEST
sed -i "s^PROJECT_REF^$CP_PROJUSER_PROJ^g" $DEST
sed -i "s^USER_REF^$CP_PROJUSER_USER^g" $DEST
