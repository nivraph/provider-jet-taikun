source config.sh

RESOURCE="organization_billing_rule_attachment"
DEST=$PATH_DST$PRE$RESOURCE$END

echo 'Org bill rule attachment name:'
read CP_ORG_BILL_RULE_NAME

echo 'Billing rule:'
read CP_ORG_BILL_RULE_BILLING

echo 'Organization:'
read CP_ORG_BILL_RULE_ORG

cp "$EXAMPLES_PATH"example_"$RESOURCE".yaml $DEST

sed -i "s^ORG_BILL_RULE_ATTACH^$CP_ORG_BILL_RULE_NAME^g" $DEST
sed -i "s^BILL_RULE_REF^$CP_ORG_BILL_RULE_BILLING^g" $DEST
sed -i "s^ORGANIZATION_REF^$CP_ORG_BILL_RULE_ORG^g" $DEST
