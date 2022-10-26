
cloudcredentialaws.taikun.jet.crossplane.io
===========================================


This document has been generated.
  

# Example


```yaml

```  

# Schema
  

## Required
  
`accessKeyIdSecretRef`: The AWS access key ID.

* `key`: The key to select.<font color="red"> (Required)</font>  

* `name`: Name of the secret.<font color="red"> (Required)</font>  

* `namespace`: Namespace of the secret.<font color="red"> (Required)</font>  
  
`availabilityZone`: The AWS availability zone for the region.
  
`name`: The name of the AWS cloud credential.
  
`region`: The AWS region.
  
`secretAccessKeySecretRef`: The AWS secret access key.

* `key`: The key to select.<font color="red"> (Required)</font>  

* `name`: Name of the secret.<font color="red"> (Required)</font>  

* `namespace`: Namespace of the secret.<font color="red"> (Required)</font>  
  

## Optional
  
`lock`: Indicates whether to lock the AWS cloud credential. Defaults to `false`.
  
`organizationId`: The ID of the organization which owns the AWS cloud credential.
  
`organizationIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="red"> (Required)</font>  
  
`organizationIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
