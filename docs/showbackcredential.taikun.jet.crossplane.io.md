
showbackcredential.taikun.jet.crossplane.io
===========================================


This document has been generated.
  

# Example


```yaml
apiVersion: showbackcredential.taikun.jet.crossplane.io/v1alpha1
kind: Credential
metadata:
  name: SHOWBACK_CREDENTIAL
spec:
  forProvider:
    name: "SHOWBACK_CREDENTIAL"
    passwordSecretRef:
      key: "password"
      name: "billing-credentials"
      namespace: "crossplane-system"
    url: "URL"
    username: "USER"
    organizationIdRef:
        name: "ORGANIZATION_REF"
    lock: LOCK
  providerConfigRef:
    name: default

```  

# Schema
  

## Required
  
`name`: The name of the showback credential.
  
`passwordSecretRef`: The Prometheus password or other credential.
  
`url`: URL of the source.
  
`username`: The Prometheus username or other credential.
  

## Optional
  
`lock`: Indicates whether to lock the showback credential. Defaults to `false`.
  
`organizationId`: The ID of the organization which owns the showback credential.
  
`organizationIdRef`: A Reference to a named object.
  
`organizationIdSelector`: A Selector selects an object.
  
