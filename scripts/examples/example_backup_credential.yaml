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
