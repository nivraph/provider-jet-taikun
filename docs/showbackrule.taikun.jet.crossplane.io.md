
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
  
`kind` (string)`:` The kind of showback rule: `General` (data source is Taikun) or `External` (data source is external, see `showback_credential_id`).
  
`metricName` (string)`:` The metric name.
  
`name` (string)`:` The name of the showback rule.
  
`price` (number)`:` Billing in CZK per selected unit.
  
`type` (string)`:` The type of showback rule: `Count` (calculate package as unit) or `Sum` (calculate per quantity).
  

## Optional
  
`globalAlertLimit` (number)`:` Set limit of alerts for all projects. Defaults to `0`.
  
`label` (array)`:` Labels linked to this showback rule.
  
`organizationId` (string)`:` The ID of the organization which owns the showback rule.
  
`organizationIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
`projectAlertLimit` (number)`:` Set limit of alerts for one project. Defaults to `0`.
  
`showbackCredentialId` (string)`:` ID of the showback credential.
  
`showbackCredentialIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`showbackCredentialIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
