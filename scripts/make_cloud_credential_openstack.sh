source config.sh

RESOURCE="cloud_credential_openstack"
DEST=$PATH_DST$PRE$RESOURCE$END

echo 'Credentials name:'
read CP_CREDS_OS_NAME

echo 'User:'
read CP_CREDS_OS_USER

echo 'Password:'
read CP_CREDS_OS_PASSWORD

echo 'URL:'
read CP_CREDS_OS_URL

echo 'Domain:'
read CP_CREDS_OS_DOMAIN

echo 'Project name:'
read CP_CREDS_OS_PROJECT

echo 'Region:'
read CP_CREDS_OS_REGION

echo 'Public network:'
read CP_CREDS_OS_PUBLIC_NETWORK

echo 'Lock:'
read CP_CREDS_OS_LOCK

echo 'Availability:'
read CP_CREDS_OS_AVAILABILITY


#MAKE SECRET YAML

echo -n "$CP_CREDS_OS_PASSWORD" | base64 > passwordb
CP_CREDS_OS_PASSWORD_B=`cat passwordb`
rm passwordb

cp "$EXAMPLES_PATH"example_secret_"$RESOURCE".yaml "$PATH_DST"secret_$RESOURCE$END
sed -i "s/PASSWORD/$CP_CREDS_OS_PASSWORD_B/g" "$PATH_DST"secret_$RESOURCE$END


#MAKE OS CREDS YAML

cp "$EXAMPLES_PATH"example_"$RESOURCE".yaml $DEST

sed -i "s/OPENSTACK/$CP_CREDS_OS_NAME/g" $DEST
sed -i "s/USER/$CP_CREDS_OS_USER/g" $DEST
sed -i "s^URL^$CP_CREDS_OS_URL^g" $DEST
sed -i "s/DOMAIN/$CP_CREDS_OS_DOMAIN/g" $DEST
sed -i "s/PROJECT/$CP_CREDS_OS_PROJECT/g" $DEST
sed -i "s/REGION/$CP_CREDS_OS_REGION/g" $DEST
sed -i "s/PUBLIC_NETWORK/$CP_CREDS_OS_PUBLIC_NETWORK/g" $DEST
sed -i "s/LOCK/$CP_CREDS_OS_LOCK/g" $DEST
sed -i "s/AVAILABILITY/$CP_CREDS_OS_AVAILABILITY/g" $DEST
