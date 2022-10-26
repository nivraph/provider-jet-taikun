
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
  
`name`: The name of the billing credential.
  
`prometheusPasswordSecretRef`: The Prometheus password.

* `key`: The key to select.<font color="orange"> (Required)</font>  

* `name`: Name of the secret.<font color="orange"> (Required)</font>  

* `namespace`: Namespace of the secret.<font color="orange"> (Required)</font>  
  
`prometheusUrl`: The Prometheus URL.
  
`prometheusUsername`: The Prometheus username.
  

## Optional
  
`lock`: Indicates whether to lock the billing credential. Defaults to `false`.
  
`organizationId`: The ID of the organization which owns the billing credential.
  
`organizationIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
