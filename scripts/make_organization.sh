source config.sh

RESOURCE="organization"
DEST=$PATH_DST$PRE$RESOURCE$END

echo 'Organization name:'
read CP_ORG_NAME

echo 'Description:'
read CP_ORG_DESCRIPTION

echo 'Discount rate:'
read CP_ORG_DISCOUNT

echo 'City:'
read CP_ORG_CITY

echo 'Billing email:'
read CP_ORG_BEMAIL

echo 'Email:'
read CP_ORG_EMAIL

echo 'Phone:'
read CP_ORG_PHONE

cp "$EXAMPLES_PATH"example_"$RESOURCE".yaml $DEST

sed -i "s/ORGANIZATION/$CP_ORG_NAME/g" $DEST
sed -i "s/DESCRIPTION/$CP_ORG_DESCRIPTION/g" $DEST
sed -i "s/DISCOUNT/$CP_ORG_DISCOUNT/g" $DEST
sed -i "s/CITY/$CP_ORG_CITY/g" $DEST
sed -i "s/BILLING_EMAIL/$CP_ORG_BEMAIL/g" $DEST
sed -i "s/EMAIL/$CP_ORG_EMAIL/g" $DEST
sed -i "s/PHONE/$CP_ORG_PHONE/g" $DEST
