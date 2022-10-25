
cloudcredentialgcp.taikun.jet.crossplane.io
===========================================


This document has been generated.
  

# Example


```yaml

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
  
`billingAccountIdSelector`: A Selector selects an object.
  
`folderId`: The folder ID of the GCP credential. Conflicts with: `import_project`.
  
`importProject`: Whether to import a project or not Defaults to `false`. Conflicts with: `billing_account_id`, `folder_id`.
  
`lock`: Indicates whether to lock the GCP cloud credential. Defaults to `false`.
  
`organizationId`: The ID of the organization which owns the GCP credential.
  
`organizationIdRef`: A Reference to a named object.
  
`organizationIdSelector`: A Selector selects an object.
  
