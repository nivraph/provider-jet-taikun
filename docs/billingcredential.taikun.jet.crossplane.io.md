
billingcredential.taikun.jet.crossplane.io
==========================================


This document has been generated from the CRD.
  

# Example


```yaml
apiVersion: billingcredential.taikun.jet.crossplane.io/v1alpha1
kind: Credential
metadata:
  name: BILLING_CREDENTIAL
spec:
  forProvider:
    name: "BILLING_CREDENTIAL"
    prometheusPasswordSecretRef:
      key: "password"
      name: "billing-credentials"
      namespace: "crossplane-system"
    prometheusUrl: "URL"
    prometheusUsername: "USERNAME"
    organizationIdRef:
        name: "ORGANIZATION_REF"
    lock: LOCK
  providerConfigRef:
    name: default

```  

# Schema
  

## Required
  
`name` (string)`:` The name of the billing credential.
  
`prometheusPasswordSecretRef` (object)`:` The Prometheus password.

* `key` (string)`:` The key to select.<font color="orange"> (Required)</font>  

* `name` (string)`:` Name of the secret.<font color="orange"> (Required)</font>  

* `namespace` (string)`:` Namespace of the secret.<font color="orange"> (Required)</font>  
  
`prometheusUrl` (string)`:` The Prometheus URL.
  
`prometheusUsername` (string)`:` The Prometheus username.
  

## Optional
  
`lock` (boolean)`:` Indicates whether to lock the billing credential. Defaults to `false`.
  
`organizationId` (string)`:` The ID of the organization which owns the billing credential.
  
`organizationIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
