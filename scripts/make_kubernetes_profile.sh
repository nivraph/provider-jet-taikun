source config.sh

RESOURCE="kubernetes_profile"
DEST=$PATH_DST$PRE$RESOURCE$END

echo 'Kubernetes profile name:'
read CP_KP_NAME

echo 'Organization:'
read CP_KP_ORG

cp "$EXAMPLES_PATH"example_"$RESOURCE".yaml $DEST

sed -i "s/KUBERNETES_PROFILE/$CP_KP_NAME/g" $DEST
sed -i "s/ORGANIZATION_REF/$CP_KP_ORG/g" $DEST
