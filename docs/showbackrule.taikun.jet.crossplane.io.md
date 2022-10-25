
showbackrule.taikun.jet.crossplane.io
=====================================


This document has been generated.
  

# Example


```yaml
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

```  

# Schema
  

## Required
  
`kind`: The kind of showback rule: `General` (data source is Taikun) or `External` (data source is external, see `showback_credential_id`).
  
`metricName`: The metric name.
  
`name`: The name of the showback rule.
  
`price`: Billing in CZK per selected unit.
  
`type`: The type of showback rule: `Count` (calculate package as unit) or `Sum` (calculate per quantity).
  

## Optional
  
`globalAlertLimit`: Set limit of alerts for all projects. Defaults to `0`.
  
`label`: Labels linked to this showback rule.
  
`organizationId`: The ID of the organization which owns the showback rule.
  
`organizationIdRef`: A Reference to a named object.
  
`organizationIdSelector`: A Selector selects an object.
  
`projectAlertLimit`: Set limit of alerts for one project. Defaults to `0`.
  
`showbackCredentialId`: ID of the showback credential.
  
`showbackCredentialIdRef`: A Reference to a named object.
  
`showbackCredentialIdSelector`: A Selector selects an object.
  
