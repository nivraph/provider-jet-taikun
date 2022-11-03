
showbackcredential.taikun.jet.crossplane.io
===========================================


This document has been generated from the CRD.
  

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
  
`name` (string)`:` The name of the showback credential.
  
`passwordSecretRef` (object)`:` The Prometheus password or other credential.

* `key` (string)`:` The key to select.<font color="orange"> (Required)</font>  

* `name` (string)`:` Name of the secret.<font color="orange"> (Required)</font>  

* `namespace` (string)`:` Namespace of the secret.<font color="orange"> (Required)</font>  
  
`url` (string)`:` URL of the source.
  
`username` (string)`:` The Prometheus username or other credential.
  

## Optional
  
`lock` (boolean)`:` Indicates whether to lock the showback credential. Defaults to `false`.
  
`organizationId` (string)`:` The ID of the organization which owns the showback credential.
  
`organizationIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
