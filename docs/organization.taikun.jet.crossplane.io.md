
organization.taikun.jet.crossplane.io
=====================================


This document has been generated.
  

# Example


```yaml
apiVersion: organization.taikun.jet.crossplane.io/v1alpha1
kind: Organization
metadata:
  name: ORGANIZATION
spec:
  forProvider:
    name: "ORGANIZATION"
    fullName: "DESCRIPTION"
    discountRate: DISCOUNT
    city: "CITY"
    billingEmail: "BILLING_EMAIL"
    email: "EMAIL"
    phone: "PHONE"
  providerConfigRef:
    name: default

```  

# Schema
  

## Required
  
`fullName`: Full name.
  
`name`: Organization's name.
  

## Optional
  
`address`: Address.
  
`billingEmail`: Billing email.
  
`city`: City.
  
`country`: Country.
  
`discountRate`: Discount rate, must be between 0 and 100 (included). Defaults to `100`.
  
`email`: Email.
  
`lock`: Indicates whether to lock the organization. Defaults to `false`.
  
`managersCanChangeSubscription`: Allow subscription to be changed by managers. Defaults to `true`.
  
`phone`: Phone number.
  
`vatNumber`: VAT number.
  
