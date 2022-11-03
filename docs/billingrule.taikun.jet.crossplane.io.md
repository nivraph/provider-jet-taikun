
billingrule.taikun.jet.crossplane.io
====================================


This document has been generated from the CRD.
  

# Example


```yaml

apiVersion: billingrule.taikun.jet.crossplane.io/v1alpha1
kind: Rule
metadata:
  name: BILLING_RULE
spec:
  forProvider:
    name: "BILLING_RULE"
    billingCredentialIdRef:
        name: "BILLING_CRED_REF"
    label:
      - key: "label"
        value: "value"
    metricName: "coredns_forward_request_duration_seconds"
    type: "TYPE"
    price: PRICE
  providerConfigRef:
    name: default

```  

# Schema
  

## Required
  
`label` (array)`:` Labels linked to the billing rule.
  
`metricName` (string)`:` The name of the Prometheus metric (e.g. volumes, flavors, networks) to bill.
  
`name` (string)`:` The name of the billing rule.
  
`price` (number)`:` The price in CZK per selected unit.
  
`type` (string)`:` The type of billing rule: `Count` (calculate package as unit) or `Sum` (calculate per quantity).
  

## Optional
  
`billingCredentialId` (string)`:` ID of the billing credential.
  
`billingCredentialIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`billingCredentialIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
