apiVersion: accessprofile.taikun.jet.crossplane.io/v1alpha1
kind: Profile
metadata:
  name: access-profile
spec:
  forProvider:
    name: "access-profile"
    organizationIdRef:
        name: "test-organization"
  providerConfigRef:
    name: default
