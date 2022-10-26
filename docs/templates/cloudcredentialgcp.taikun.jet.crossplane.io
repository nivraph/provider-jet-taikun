apiVersion: cloudcredentialgcp.taikun.jet.crossplane.io/v1alpha1
kind: CredentialGCP
metadata:
  name: cred-gcp
spec:
  forProvider:
    name: "cred-gcp"
    configFile: "path-to-the-json-gcp-config-file"
    region: "europe-central2"
    az_count: 3
    organizationId: "15727"
    lock: false
    importProject: true
