source config.sh

RESOURCE="backup_credential"
DEST=$PATH_DST$PRE$RESOURCE$END

echo 'Backup credential name:'
read CP_BACKUP_CRED_NAME

echo 'Access key ID:'
read CP_BACKUP_CRED_ACCESS_KEY

echo 'Secret access key:'
read CP_BACKUP_CRED_SECRET_KEY

echo 'Endpoint:'
read CP_BACKUP_CRED_ENDPOINT

echo 'Region:'
read CP_BACKUP_CRED_REGION

cp "$EXAMPLES_PATH"example_"$RESOURCE".yaml $DEST

sed -i "s^BACKUP_CREDENTIAL^$CP_BACKUP_CRED_NAME^g" $DEST
sed -i "s^ACCESS_KEY_ID^$CP_BACKUP_CRED_ACCESS_KEY^g" $DEST
sed -i "s^ENDPOINT^$CP_BACKUP_CRED_ENDPOINT^g" $DEST
sed -i "s^REGION^$CP_BACKUP_CRED_REGION^g" $DEST
