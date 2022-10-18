apiVersion: billingcredential.taikun.jet.crossplane.io/v1alpha1
kind: Credential
metadata:
  name: BILLING_CREDENTIAL
spec:
  forProvider:
    name: "BILLING_CREDENTIAL"
    prometheusPasswordSecretRef:
      key: "password"
      name: "billing-credentials"
      namespace: "crossplane-system"
    prometheusUrl: "URL"
    prometheusUsername: "USERNAME"
    organizationIdRef:
        name: "ORGANIZATION_REF"
    lock: LOCK
  providerConfigRef:
    name: default
