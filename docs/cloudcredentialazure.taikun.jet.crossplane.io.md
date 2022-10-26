
cloudcredentialazure.taikun.jet.crossplane.io
=============================================


This document has been generated from the CRD.
  

# Example


```yaml

```  

# Schema
  

## Required
  
`availabilityZone`: The Azure availability zone for the location.
  
`clientIdSecretRef`: The Azure client ID.

* `key`: The key to select.<font color="orange"> (Required)</font>  

* `name`: Name of the secret.<font color="orange"> (Required)</font>  

* `namespace`: Namespace of the secret.<font color="orange"> (Required)</font>  
  
`clientSecretSecretRef`: The Azure client secret.

* `key`: The key to select.<font color="orange"> (Required)</font>  

* `name`: Name of the secret.<font color="orange"> (Required)</font>  

* `namespace`: Namespace of the secret.<font color="orange"> (Required)</font>  
  
`location`: The Azure location.
  
`name`: The name of the Azure cloud credential.
  
`subscriptionId`: The Azure subscription ID.
  
`tenantId`: The Azure tenant ID.
  

## Optional
  
`lock`: Indicates whether to lock the Azure cloud credential. Defaults to `false`.
  
`organizationId`: The ID of the organization which owns the Azure cloud credential.
  
`organizationIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
