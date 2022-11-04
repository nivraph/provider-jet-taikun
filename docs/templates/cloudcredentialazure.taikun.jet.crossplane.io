apiVersion: cloudcredentialazure.taikun.jet.crossplane.io/v1alpha1
kind: CredentialAzure
metadata:
  name: cred-azure
spec:
  forProvider:
    name: "cred-azure"
    az_count: 1
    clientIdSecretRef:
      key: "client_id"
      name: "azure-credentials"
      namespace: "crossplane-system"
    clientSecretSecretRef:
      key: "client_secret"
      name: "azure-credentials"
      namespace: "crossplane-system"
    location: "westeurope"
    subscriptionId: "azure-subscription-id"
    tenantId: "azure-tenant-id"
    organizationId: "15687"
    lock: false
