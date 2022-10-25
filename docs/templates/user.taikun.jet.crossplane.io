apiVersion: user.taikun.jet.crossplane.io/v1alpha1
kind: User
metadata:
    name: USER
spec:
  forProvider:
    userName: "USER"
    email: "EMAIL"
    role: "ROLE"
    organizationIdRef:
        name: "ORGANIZATION_REF"
  providerConfigRef:
    name: default