
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
  
`label`: Labels linked to the billing rule.
  
`metricName`: The name of the Prometheus metric (e.g. volumes, flavors, networks) to bill.
  
`name`: The name of the billing rule.
  
`price`: The price in CZK per selected unit.
  
`type`: The type of billing rule: `Count` (calculate package as unit) or `Sum` (calculate per quantity).
  

## Optional
  
`billingCredentialId`: ID of the billing credential.
  
`billingCredentialIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="orange"> (Required)</font>  
  
`billingCredentialIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
