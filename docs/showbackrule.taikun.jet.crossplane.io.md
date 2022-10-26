
showbackrule.taikun.jet.crossplane.io
=====================================


This document has been generated from the CRD.
  

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

* `name`: Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
`projectAlertLimit`: Set limit of alerts for one project. Defaults to `0`.
  
`showbackCredentialId`: ID of the showback credential.
  
`showbackCredentialIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="orange"> (Required)</font>  
  
`showbackCredentialIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
