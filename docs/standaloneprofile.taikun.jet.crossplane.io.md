
standaloneprofile.taikun.jet.crossplane.io
==========================================


This document has been generated.
  

# Example


```yaml
apiVersion: standaloneprofile.taikun.jet.crossplane.io/v1alpha1
kind: Profile
metadata:
  name: STANDALONE
spec:
  forProvider:
    name: "STANDALONE"
    publicKey: "PUBLIC_KEY"
    organizationIdRef:
        name: "ORGANIZATION_REF"
  providerConfigRef:
    name: default

```  

# Schema
  

## Required
  
`name`: The name of the standalone profile.
  
`publicKey`: The public key of the standalone profile.
  

## Optional
  
`lock`: Indicates whether to lock the standalone profile. Defaults to `false`.
  
`organizationId`: The ID of the organization which owns the standalone profile.
  
`organizationIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="red"> (Required)</font>  
  
`organizationIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
`securityGroup`: List of security groups.
  
