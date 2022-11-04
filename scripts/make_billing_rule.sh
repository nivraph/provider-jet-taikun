source config.sh

RESOURCE="billing_rule"
DEST=$PATH_DST$PRE$RESOURCE$END

echo 'Billing rule name:'
read CP_BILL_RULE_NAME

echo 'Billing credentials:'
read CP_BILL_RULE_CREDS

echo 'Type:'
read CP_BILL_RULE_TYPE

echo 'Price:'
read CP_BILL_RULE_PRICE

cp "$EXAMPLES_PATH"example_"$RESOURCE".yaml $DEST

sed -i "s^BILLING_RULE^$CP_BILL_RULE_NAME^g" $DEST
sed -i "s^BILLING_CRED_REF^$CP_BILL_RULE_CREDS^g" $DEST
sed -i "s^TYPE^$CP_BILL_RULE_TYPE^g" $DEST
sed -i "s^PRICE^$CP_BILL_RULE_PRICE^g" $DEST
