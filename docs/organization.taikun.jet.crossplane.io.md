
organization.taikun.jet.crossplane.io
=====================================


This document has been generated from the CRD.
  

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
  
`fullName` (string)`:` Full name.
  
`name` (string)`:` Organization's name.
  

## Optional
  
`address` (string)`:` Address.
  
`billingEmail` (string)`:` Billing email.
  
`city` (string)`:` City.
  
`country` (string)`:` Country.
  
`discountRate` (number)`:` Discount rate, must be between 0 and 100 (included). Defaults to `100`.
  
`email` (string)`:` Email.
  
`lock` (boolean)`:` Indicates whether to lock the organization. Defaults to `false`.
  
`managersCanChangeSubscription` (boolean)`:` Allow subscription to be changed by managers. Defaults to `true`.
  
`phone` (string)`:` Phone number.
  
`vatNumber` (string)`:` VAT number.
  
