
backuppolicy.taikun.jet.crossplane.io
=====================================


This document has been generated from the CRD.
  

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
  
`cronPeriod` (string)`:` Frequency of backups.
  
`name` (string)`:` The name of the backup policy.
  

## Optional
  
`excludedNamespaces` (array)`:` Namespaces excluded from the backups.
  
`includedNamespaces` (array)`:` Namespaces included in the backups.
  
`projectId` (string)`:` The ID of the project.
  
`projectIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`projectIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
`retentionPeriod` (string)`:` How long to store the backups. Defaults to `720h`.
  
