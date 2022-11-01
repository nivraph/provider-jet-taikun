apiVersion: cloudcredential.taikun.jet.crossplane.io/v1alpha1
kind: Credential
metadata:
  name: cloud-credential
spec:
  forProvider:
    name: "cloud-credential"
    type: "openstack"
    user: "user"
    passwordSecretRef:
      key: "password"
      name: "openstack-credentials"
      namespace: "namespace"
    organizationIdRef:
        name: "organization-name"
    url: "url"
    domain: "domain"
    projectName: "project"
    region: "region"
    publicNetworkName: "public-network"
    lock: true
    availabilityZone: "az-2"
