
backupcredential.taikun.jet.crossplane.io
=========================================


This document has been generated from the CRD.
  

# Example


```yaml
apiVersion: backupcredential.taikun.jet.crossplane.io/v1alpha1
kind: Credential
metadata:
  name: BACKUP_CREDENTIAL
spec:
  forProvider:
    name: "BACKUP_CREDENTIAL"
    s3AccessKeyId: ACCESS_KEY_ID
    s3Endpoint: ENDPOINT
    s3Region: REGION
    s3SecretAccessKeySecretRef:
      key: "password"
      name: "backup-credentials"
      namespace: "crossplane-system"
    organizationIdRef:
        name: "ORGANIZATION_REF"
    lock: LOCK
  providerConfigRef:
    name: default

```  

# Schema
  

## Required
  
`name` (string)`:` The name of the backup credential.
  
`s3AccessKeyId` (string)`:` The S3 access key ID.
  
`s3Endpoint` (string)`:` The S3 endpoint URL.
  
`s3Region` (string)`:` The S3 region.
  
`s3SecretAccessKeySecretRef` (object)`:` The S3 secret access key.

* `key` (string)`:` The key to select.<font color="orange"> (Required)</font>  

* `name` (string)`:` Name of the secret.<font color="orange"> (Required)</font>  

* `namespace` (string)`:` Namespace of the secret.<font color="orange"> (Required)</font>  
  

## Optional
  
`lock` (boolean)`:` Indicates whether to lock the backup credential. Defaults to `false`.
  
`organizationId` (string)`:` The ID of the organization which owns the backup credential.
  
`organizationIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
