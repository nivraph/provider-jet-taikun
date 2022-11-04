apiVersion: alertingprofile.taikun.jet.crossplane.io/v1alpha1
kind: Profile
metadata:
  name: alerting-profile
spec:
  forProvider:
    name: "test-alerting-profile"
    emails:
      - "foo@bar.com"
    lock: false
    organizationIdRef:
      name: "test-organization"
    reminder: "Daily"
  providerConfigRef:
    name: default
