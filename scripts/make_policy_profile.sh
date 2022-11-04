source config.sh

RESOURCE="policy_profile"
DEST=$PATH_DST$PRE$RESOURCE$END

echo 'Policy profile name:'
read CP_POLICY_PROFILE

echo 'Organization:'
read CP_POLICY_ORG

echo 'Allowed repos:'
read CP_POLICY_REPOS

echo 'Forbidden tags:'
read CP_POLICY_TAGS

echo 'Lock:'
read CP_POLICY_LOCK

echo 'Forbid http ingress:'
read CP_POLICY_HTTP

echo 'Forbid node port:'
read CP_POLICY_NODE

echo 'Require probe:'
read CP_POLICY_PROBE

echo 'Unique ingress:'
read CP_POLICY_INGRESS

echo 'Unique service selector:'
read CP_POLICY_SELECTOR

cp "$EXAMPLES_PATH"example_"$RESOURCE".yaml $DEST

sed -i "s^POLICY_PROFILE^$CP_POLICY_PROFILE^g" $DEST
sed -i "s^ORGANIZATION_REF^$CP_POLICY_ORG^g" $DEST
sed -i "s^REPO^$CP_POLICY_REPOS^g" $DEST
sed -i "s^TAG^$CP_POLICY_TAGS^g" $DEST
sed -i "s^LOCK^$CP_POLICY_LOCK^g" $DEST
sed -i "s^HTTP^$CP_POLICY_HTTP^g" $DEST
sed -i "s^NODE^$CP_POLICY_NODE^g" $DEST
sed -i "s^PROBE^$CP_POLICY_PROBE^g" $DEST
sed -i "s^INGRESS^$CP_POLICY_INGRESS^g" $DEST
sed -i "s^SELECTOR^$CP_POLICY_SELECTOR^g" $DEST
