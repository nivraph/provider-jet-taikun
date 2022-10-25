
organizationbillingruleattachment.taikun.jet.crossplane.io
==========================================================


This document has been generated.
  

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
  
`billingRuleIdSelector`: A Selector selects an object.
  
`discountRate`: Discount rate in percents (0-100 %). Defaults to `100`.
  
`organizationId`: ID of the organisation.
  
`organizationIdRef`: A Reference to a named object.
  
`organizationIdSelector`: A Selector selects an object.
  
