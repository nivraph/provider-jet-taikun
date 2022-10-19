source config.sh

RESOURCE="billing_credential"
DEST=$PATH_DST$PRE$RESOURCE$END

echo 'Billing credential name:'
read CP_BILL_CRED_NAME

echo 'URL:'
read CP_BILL_CRED_URL

echo 'Username:'
read CP_BILL_CRED_USER

echo 'Password:'
read CP_BILL_CRED_PASSWORD

echo 'Organization:'
read CP_BILL_CRED_ORG

echo 'Lock:'
read CP_BILL_CRED_LOCK


#MAKE SECRET YAML

echo -n "$CP_BILL_CRED_PASSWORD" | base64 > passwordb
CP_BILL_CRED_PASSWORD_B=`cat passwordb`
rm passwordb

cp "$EXAMPLES_PATH"example_secret_"$RESOURCE".yaml "$PATH_DST"secret_$RESOURCE$END
sed -i "s/PASSWORD/$CP_BILL_CRED_PASSWORD_B/g" "$PATH_DST"secret_$RESOURCE$END


#MAKE OS CREDS YAML

cp "$EXAMPLES_PATH"example_"$RESOURCE".yaml $DEST

sed -i "s/BILLING_CREDENTIAL/$CP_BILL_CRED_NAME/g" $DEST
sed -i "s^URL^$CP_BILL_CRED_URL^g" $DEST
sed -i "s/USERNAME/$CP_BILL_CRED_USER/g" $DEST
sed -i "s/ORGANIZATION_REF/$CP_BILL_CRED_ORG/g" $DEST
sed -i "s/LOCK/$CP_BILL_CRED_LOCK/g" $DEST
