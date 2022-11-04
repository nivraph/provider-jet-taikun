source config.sh

RESOURCE="showback_credential"
DEST=$PATH_DST$PRE$RESOURCE$END

echo 'Showback credential name:'
read CP_SHOWBACK_CRED_NAME

echo 'URL:'
read CP_SHOWBACK_CRED_URL

echo 'Username:'
read CP_SHOWBACK_CRED_USER

echo 'Password:'
read CP_SHOWBACK_CRED_PASSWORD

echo 'Organization:'
read CP_SHOWBACK_CRED_ORG

echo 'Lock:'
read CP_SHOWBACK_CRED_LOCK


#MAKE SECRET YAML

echo -n "$CP_SHOWBACK_CRED_PASSWORD" | base64 > passwordb
CP_SHOWBACK_CRED_PASSWORD_B=`cat passwordb`
rm passwordb

cp "$EXAMPLES_PATH"example_secret_"$RESOURCE".yaml "$PATH_DST"secret_$RESOURCE$END
sed -i "s/PASSWORD/$CP_SHOWBACK_CRED_PASSWORD_B/g" "$PATH_DST"secret_$RESOURCE$END


#MAKE OS CREDS YAML

cp "$EXAMPLES_PATH"example_"$RESOURCE".yaml $DEST

sed -i "s^SHOWBACK_CREDENTIAL^$CP_SHOWBACK_CRED_NAME^g" $DEST
sed -i "s^URL^$CP_SHOWBACK_CRED_URL^g" $DEST
sed -i "s^USER^$CP_SHOWBACK_CRED_USER^g" $DEST
sed -i "s^ORGANIZATION_REF^$CP_SHOWBACK_CRED_ORG^g" $DEST
sed -i "s^LOCK^$CP_SHOWBACK_CRED_LOCK^g" $DEST
