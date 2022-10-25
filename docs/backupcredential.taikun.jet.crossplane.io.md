
backupcredential.taikun.jet.crossplane.io
=========================================


This document has been generated.
  

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
  
`name`: The name of the backup credential.
  
`s3AccessKeyId`: The S3 access key ID.
  
`s3Endpoint`: The S3 endpoint URL.
  
`s3Region`: The S3 region.
  
`s3SecretAccessKeySecretRef`: The S3 secret access key.
  

## Optional
  
`lock`: Indicates whether to lock the backup credential. Defaults to `false`.
  
`organizationId`: The ID of the organization which owns the backup credential.
  
`organizationIdRef`: A Reference to a named object.
  
`organizationIdSelector`: A Selector selects an object.
  
