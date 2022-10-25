
user.taikun.jet.crossplane.io
=============================


This document has been generated.
  

# Example


```yaml
apiVersion: user.taikun.jet.crossplane.io/v1alpha1
kind: User
metadata:
    name: USER
spec:
  forProvider:
    userName: "USER"
    email: "EMAIL"
    role: "ROLE"
    organizationIdRef:
        name: "ORGANIZATION_REF"
  providerConfigRef:
    name: default
```  

# Schema
  

## Required
  
`email`: The email of the user.
  
`role`: The role of the user: `Manager` or `User`.
  
`userName`: The name of the user.
  

## Optional
  
`displayName`: The user's display name. Defaults to ` `.
  
`organizationId`: The ID of the user's organization.
  
`organizationIdRef`: A Reference to a named object.
  
`organizationIdSelector`: A Selector selects an object.
  
