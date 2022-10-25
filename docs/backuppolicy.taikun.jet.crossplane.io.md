
backuppolicy.taikun.jet.crossplane.io
=====================================


This document has been generated.
  

# Example


```yaml
apiVersion: backuppolicy.taikun.jet.crossplane.io/v1alpha1
kind: Policy
metadata:
  name: BACKUP_POLICY
spec:
  forProvider:
    name: "BACKUP_POLICY"
    cronPeriod: "PERIOD"
    projectIdRef:
        name: "PROJECT_REF"
    includedNamespaces: ["INC_NAMESPACE"]
  providerConfigRef:
    name: default

```  

# Schema
  

## Required
  
`cronPeriod`: Frequency of backups.
  
`name`: The name of the backup policy.
  

## Optional
  
`excludedNamespaces`: Namespaces excluded from the backups.
  
`includedNamespaces`: Namespaces included in the backups.
  
`projectId`: The ID of the project.
  
`projectIdRef`: A Reference to a named object.
  
`projectIdSelector`: A Selector selects an object.
  
`retentionPeriod`: How long to store the backups. Defaults to `720h`.
  
