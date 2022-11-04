apiVersion: backuppolicy.taikun.jet.crossplane.io/v1alpha1
kind: Policy
metadata:
  name: BACKUP_POLICY
spec:
  forProvider:
    name: "BACKUP_POLICY"
    cronPeriod: "PERIOD"
    projectIdRef:
        name: "PROJECT_REF"
    includedNamespaces: ["INC_NAMESPACE"]
  providerConfigRef:
    name: default
