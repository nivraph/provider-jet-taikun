source config.sh

RESOURCE="backup_policy"
DEST=$PATH_DST$PRE$RESOURCE$END

echo 'Backup policy name:'
read CP_BACKUP_POLICY_NAME

echo 'Cron period:'
read CP_BACKUP_POLICY_PERIOD

echo 'Project:'
read CP_BACKUP_POLICY_PROJECT

echo 'Included namespace:'
read CP_BACKUP_POLICY_INC

cp "$EXAMPLES_PATH"example_"$RESOURCE".yaml $DEST

sed -i "s^BACKUP_POLICY^$CP_BACKUP_POLICY_NAME^g" $DEST
sed -i "s^PERIOD^$CP_BACKUP_POLICY_PERIOD^g" $DEST
sed -i "s^PROJECT_REF^$CP_BACKUP_POLICY_PROJECT^g" $DEST
sed -i "s^INC_NAMESPACE^$CP_BACKUP_POLICY_INC^g" $DEST
