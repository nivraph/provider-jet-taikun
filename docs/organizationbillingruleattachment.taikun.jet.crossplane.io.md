
organizationbillingruleattachment.taikun.jet.crossplane.io
==========================================================


This document has been generated from the CRD.
  

# Example


```yaml
apiVersion: organizationbillingruleattachment.taikun.jet.crossplane.io/v1alpha1
kind: BillingRuleAttachment
metadata:
  name: ORG_BILL_RULE_ATTACH
spec:
  forProvider:
    billingRuleIdRef:
        name: "BILL_RULE_REF"
    organizationIdRef:
        name: "ORGANIZATION_REF"
  providerConfigRef:
    name: default
```  

# Schema
  

## Optional
  
`billingRuleId` (string)`:` ID of the billing rule.
  
`billingRuleIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`billingRuleIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
`discountRate` (number)`:` Discount rate in percents (0-100 %). Defaults to `100`.
  
`organizationId` (string)`:` ID of the organisation.
  
`organizationIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
