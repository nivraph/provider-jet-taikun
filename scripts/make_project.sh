source config.sh

RESOURCE="project"
DEST=$PATH_DST$PRE$RESOURCE$END

echo 'Project name:'
read CP_PROJECT_NAME

echo 'Cloud credentials:'
read CP_PROJECT_CLOUD_CREDS

echo 'Organization:'
read CP_PROJECT_ORG

echo 'Quota Disk size:'
read CP_PROJECT_QDISK

echo 'Quota RAM size:'
read CP_PROJECT_QRAM

echo 'Kubernetes profile:'
read CP_PROJECT_KP

echo 'Flavor:'
read CP_PROJECT_FLAVOR

echo 'Image:'
read CP_PROJECT_IMG

echo 'VM Name:'
read CP_PROJECT_VM_NAME

echo 'VM Flavor:'
read CP_PROJECT_VM_FLAVOR

echo 'VM Image ID:'
read CP_PROJECT_VM_IMG_ID

echo 'VM Standalone profile:'
read CP_PROJECT_VM_SA

echo 'VM Volume size:'
read CP_PROJECT_VM_VOL_SIZE

echo 'VM Username:'
read CP_PROJECT_VM_USER

cp "$EXAMPLES_PATH"example_"$RESOURCE".yaml $DEST

sed -i "s/PROJECT_NAME/$CP_PROJECT_NAME/g" $DEST
sed -i "s/CLOUD_REF/$CP_PROJECT_CLOUD_CREDS/g" $DEST
sed -i "s/ORGANIZATION_REF/$CP_PROJECT_ORG/g" $DEST
sed -i "s/QDISK_SIZE/$CP_PROJECT_QDISK/g" $DEST
sed -i "s/QRAM_SIZE/$CP_PROJECT_QRAM/g" $DEST
sed -i "s/KP_REF/$CP_PROJECT_KP/g" $DEST
sed -i "s/PROJECT_FLAVOR/$CP_PROJECT_FLAVOR/g" $DEST
sed -i "s/IMG/$CP_PROJECT_IMG/g" $DEST
sed -i "s/VM_FLAVOR/$CP_PROJECT_VM_FLAVOR/g" $DEST
sed -i "s/VM_IMAGE_ID/$CP_PROJECT_VM_IMG_ID/g" $DEST
sed -i "s/VM_NAME/$CP_PROJECT_VM_NAME/g" $DEST
sed -i "s/VM_SA/$CP_PROJECT_VM_SA/g" $DEST
sed -i "s/VM_VOLUME_SIZE/$CP_PROJECT_VM_VOL_SIZE/g" $DEST
sed -i "s/VM_USER/$CP_PROJECT_VM_USER/g" $DEST
