
providerconfigs.taikun.jet.crossplane.io
========================================


This document has been generated from the CRD.
  

# Example


```yaml
apiVersion: taikun.jet.crossplane.io/v1alpha1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: example-creds
      namespace: crossplane-system
      key: credentials

```  

# Schema
  

## Required
  
`credentials`: Credentials required to authenticate to this provider.

* `source`: Source of the provider credentials.<font color="orange"> (Required)</font>  

* `env`: Env is a reference to an environment variable that contains credentials that must be used to connect to the provider.
  * `name`: Name is the name of an environment variable.<font color="orange"> (Required)</font>  
  

* `fs`: Fs is a reference to a filesystem location that contains credentials that must be used to connect to the provider.
  * `path`: Path is a filesystem path.<font color="orange"> (Required)</font>  
  

* `secretRef`: A SecretRef is a reference to a secret key that contains the credentials that must be used to connect to the provider.
  * `key`: The key to select.<font color="orange"> (Required)</font>  

  * `name`: Name of the secret.<font color="orange"> (Required)</font>  

  * `namespace`: Namespace of the secret.<font color="orange"> (Required)</font>  
  
  
