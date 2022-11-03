
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
  
`configFile` (string)`:` The path of the GCP credential's configuration file.
  
`name` (string)`:` The name of the GCP credential.
  
`region` (string)`:` The region of the GCP credential.
  
`zone` (string)`:` The zone of the GCP credential.
  

## Optional
  
`billingAccountId` (string)`:` The ID of the GCP credential's billing account. Conflicts with: `import_project`.
  
`billingAccountIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`billingAccountIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
`folderId` (string)`:` The folder ID of the GCP credential. Conflicts with: `import_project`.
  
`importProject` (boolean)`:` Whether to import a project or not Defaults to `false`. Conflicts with: `billing_account_id`, `folder_id`.
  
`lock` (boolean)`:` Indicates whether to lock the GCP cloud credential. Defaults to `false`.
  
`organizationId` (string)`:` The ID of the organization which owns the GCP credential.
  
`organizationIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
