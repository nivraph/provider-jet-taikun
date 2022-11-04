
standaloneprofile.taikun.jet.crossplane.io
==========================================


This document has been generated from the CRD.
  

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
  
`name` (string)`:` The name of the standalone profile.
  
`publicKey` (string)`:` The public key of the standalone profile.
  

## Optional
  
`lock` (boolean)`:` Indicates whether to lock the standalone profile. Defaults to `false`.
  
`organizationId` (string)`:` The ID of the organization which owns the standalone profile.
  
`organizationIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
`securityGroup` (array)`:` List of security groups.
  
