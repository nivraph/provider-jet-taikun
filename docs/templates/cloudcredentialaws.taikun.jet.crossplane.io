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
