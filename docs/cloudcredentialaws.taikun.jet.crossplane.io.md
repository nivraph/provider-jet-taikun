
cloudcredentialaws.taikun.jet.crossplane.io
===========================================


This document has been generated from the CRD.
  

# Example


```yaml
apiVersion: cloudcredentialaws.taikun.jet.crossplane.io/v1alpha1
kind: CredentialAws
metadata:
  name: cred-aws
spec:
  forProvider:
    name: "cred-aws"
    accessKeyIdSecretRef:
      key: "access_key_id"
      name: "aws-credentials"
      namespace: "crossplane-system"
    az_count: 2
    region: "eu-west-3"
    secretAccessKeySecretRef:
      key: "secret_access_key"
      name: "aws-credentials"
      namespace: "crossplane-system"
    organizationId: "15687"
    lock: false

```  

# Schema
  

## Required
  
`accessKeyIdSecretRef`: The AWS access key ID.

* `key`: The key to select.<font color="orange"> (Required)</font>  

* `name`: Name of the secret.<font color="orange"> (Required)</font>  

* `namespace`: Namespace of the secret.<font color="orange"> (Required)</font>  
  
`availabilityZone`: The AWS availability zone for the region.
  
`name`: The name of the AWS cloud credential.
  
`region`: The AWS region.
  
`secretAccessKeySecretRef`: The AWS secret access key.

* `key`: The key to select.<font color="orange"> (Required)</font>  

* `name`: Name of the secret.<font color="orange"> (Required)</font>  

* `namespace`: Namespace of the secret.<font color="orange"> (Required)</font>  
  

## Optional
  
`lock`: Indicates whether to lock the AWS cloud credential. Defaults to `false`.
  
`organizationId`: The ID of the organization which owns the AWS cloud credential.
  
`organizationIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
