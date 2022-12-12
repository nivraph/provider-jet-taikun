#!/bin/bash

export JUMP_TIME=2

cp example_required-test-organization.yaml required-test-organization.yaml
cp example_required-test-user.yaml required-test-user.yaml
cp example_required-test-kubernetes-profile.yaml required-test-kubernetes-profile.yaml
cp example_required-test-cloud-credential.yaml required-test-cloud-credential.yaml

cp example_get_cloud_ref.sh get_cloud_ref.sh
cp example_get_organization_ref.sh get_organization_ref.sh
cp example_get_user_ref.sh get_user_ref.sh


find required* -type f -exec sed -i "s^PROVIDER^$PROVIDER_NAME^g" {} \; &> /dev/null
#find test_* -type f -exec sed -i "s^PROVIDER^$PROVIDER_NAME^g" {} \; &> /dev/null

#find get_* -type f -exec sed -i "s^USER_TEST^$TAIKUN_USER^g" {} \; &> /dev/null
find required* -type f -exec sed -i "s^USER_TEST^$TAIKUN_USER^g" {} \; &> /dev/null
#find test_* -type f -exec sed -i "s^USER_TEST^$TAIKUN_USER^g" {} \; &> /dev/null

#echo -n "test-USER_TEST-org-attachach" > ref

# CREATE SECRET FOR OPENSTACK CLOUD PASSWORD
cp example_secret_cloud.yaml secret_cloud.yaml
echo -n "$OPENSTACK_CLOUD_PASSWORD" | base64 > passwordb
OPENSTACK_CLOUD_PASSWORD_B=`cat passwordb`
rm passwordb
sed -i "s/PASSWORD/$OPENSTACK_CLOUD_PASSWORD_B/g" secret_cloud.yaml
kubectl apply -f secret_cloud.yaml



# CREATE SECRET FOR PROMETHEUS PASSWORD

#cp example_secret_showback.yaml secret_showback.yaml
#echo -n "$PROMETHEUS_PASSWORD" | base64 > passwordb
#PROMETHEUS_PASSWORD_B=`cat passwordb`
#rm passwordb
#sed -i "s/PASSWORD/$PROMETHEUS_PASSWORD_B/g" secret_showback.yaml
#kubectl apply -f secret_showback.yaml



# CONFIGURE YAML FOR CLOUD CREDENTIAL

sed -i "s/USER/$OPENSTACK_CLOUD_USER/g" required-test-cloud-credential.yaml



# APPLY THE REQUIRED RESOURCES

kubectl apply -f required-test-organization.yaml
kubectl apply -f required-test-user.yaml

#kubectl apply -f required-test-billing-credential.yaml
#kubectl apply -f required-test-billing-rule.yaml
#kubectl apply -f required-test-showback-credential.yaml

#kubectl apply -f required-test-standalone-profile.yaml
kubectl apply -f required-test-kubernetes-profile.yaml

#kubectl apply -f required-test-project.yaml


# GET SOME USEFUL REFS OF RESOURCES

source get_organization_ref.sh
#source get_standalone_ref.sh
#source get_project_ref.sh
source get_user_ref.sh

#cp example-cloud-credential.yaml required-test-cloud-credential.yaml
#sed -i "s/ORG_REF/$TESTS_ORGANIZATION_REF/g" required-test-cloud-credential.yaml

kubectl apply -f required-test-cloud-credential.yaml

source get_cloud_ref.sh

#sleep 35

