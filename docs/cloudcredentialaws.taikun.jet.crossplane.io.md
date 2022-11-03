
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
  
`accessKeyIdSecretRef` (object)`:` The AWS access key ID.

* `key` (string)`:` The key to select.<font color="orange"> (Required)</font>  

* `name` (string)`:` Name of the secret.<font color="orange"> (Required)</font>  

* `namespace` (string)`:` Namespace of the secret.<font color="orange"> (Required)</font>  
  
`availabilityZone` (string)`:` The AWS availability zone for the region.
  
`name` (string)`:` The name of the AWS cloud credential.
  
`region` (string)`:` The AWS region.
  
`secretAccessKeySecretRef` (object)`:` The AWS secret access key.

* `key` (string)`:` The key to select.<font color="orange"> (Required)</font>  

* `name` (string)`:` Name of the secret.<font color="orange"> (Required)</font>  

* `namespace` (string)`:` Namespace of the secret.<font color="orange"> (Required)</font>  
  

## Optional
  
`lock` (boolean)`:` Indicates whether to lock the AWS cloud credential. Defaults to `false`.
  
`organizationId` (string)`:` The ID of the organization which owns the AWS cloud credential.
  
`organizationIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
