apiVersion: showbackrule.taikun.jet.crossplane.io/v1alpha1
kind: Rule
metadata:
  name: SHOWBACK_RULE
spec:
  forProvider:
    name: "SHOWBACK_RULE"
    kind: "KIND"
    label:
      - key: "key"
        value: "value"
    metricName: "metricname"
    type: "TYPE"
    price: PRICE
    globalAlertLimit: GLOBAL_ALERT
    projectAlertLimit: PROJECT_ALERT
    organizationIdRef:
        name: "ORGANIZATION_REF"
    showbackCredentialIdRef:
        name: "SHOWBACK_CREDENTIAL"
  providerConfigRef:
    name: default
