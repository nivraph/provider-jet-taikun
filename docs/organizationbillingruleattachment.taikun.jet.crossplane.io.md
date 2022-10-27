
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
  
`billingRuleId`: ID of the billing rule.
  
`billingRuleIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="orange"> (Required)</font>  
  
`billingRuleIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
`discountRate`: Discount rate in percents (0-100 %). Defaults to `100`.
  
`organizationId`: ID of the organisation.
  
`organizationIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
