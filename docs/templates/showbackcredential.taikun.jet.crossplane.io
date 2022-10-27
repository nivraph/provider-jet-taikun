apiVersion: showbackcredential.taikun.jet.crossplane.io/v1alpha1
kind: Credential
metadata:
  name: SHOWBACK_CREDENTIAL
spec:
  forProvider:
    name: "SHOWBACK_CREDENTIAL"
    passwordSecretRef:
      key: "password"
      name: "billing-credentials"
      namespace: "crossplane-system"
    url: "URL"
    username: "USER"
    organizationIdRef:
        name: "ORGANIZATION_REF"
    lock: LOCK
  providerConfigRef:
    name: default
