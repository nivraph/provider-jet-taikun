
cloudcredentialgcp.taikun.jet.crossplane.io
===========================================


This document has been generated from the CRD.
  

# Example


```yaml
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

```  

# Schema
  

## Required
  
`configFile`: The path of the GCP credential's configuration file.
  
`name`: The name of the GCP credential.
  
`region`: The region of the GCP credential.
  
`zone`: The zone of the GCP credential.
  

## Optional
  
`billingAccountId`: The ID of the GCP credential's billing account. Conflicts with: `import_project`.
  
`billingAccountIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="orange"> (Required)</font>  
  
`billingAccountIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
`folderId`: The folder ID of the GCP credential. Conflicts with: `import_project`.
  
`importProject`: Whether to import a project or not Defaults to `false`. Conflicts with: `billing_account_id`, `folder_id`.
  
`lock`: Indicates whether to lock the GCP cloud credential. Defaults to `false`.
  
`organizationId`: The ID of the organization which owns the GCP credential.
  
`organizationIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
