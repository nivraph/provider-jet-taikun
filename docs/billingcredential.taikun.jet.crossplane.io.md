
billingcredential.taikun.jet.crossplane.io
==========================================


This document has been generated.
  

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

* `key`: The key to select.<font color="red"> (Required)</font>  

* `name`: Name of the secret.<font color="red"> (Required)</font>  

* `namespace`: Namespace of the secret.<font color="red"> (Required)</font>  
  
`prometheusUrl`: The Prometheus URL.
  
`prometheusUsername`: The Prometheus username.
  

## Optional
  
`lock`: Indicates whether to lock the billing credential. Defaults to `false`.
  
`organizationId`: The ID of the organization which owns the billing credential.
  
`organizationIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="red"> (Required)</font>  
  
`organizationIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
