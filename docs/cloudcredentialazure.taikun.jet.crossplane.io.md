
cloudcredentialazure.taikun.jet.crossplane.io
=============================================


This document has been generated from the CRD.
  

# Example


```yaml
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

```  

# Schema
  

## Required
  
`availabilityZone` (string)`:` The Azure availability zone for the location.
  
`clientIdSecretRef` (object)`:` The Azure client ID.

* `key` (string)`:` The key to select.<font color="orange"> (Required)</font>  

* `name` (string)`:` Name of the secret.<font color="orange"> (Required)</font>  

* `namespace` (string)`:` Namespace of the secret.<font color="orange"> (Required)</font>  
  
`clientSecretSecretRef` (object)`:` The Azure client secret.

* `key` (string)`:` The key to select.<font color="orange"> (Required)</font>  

* `name` (string)`:` Name of the secret.<font color="orange"> (Required)</font>  

* `namespace` (string)`:` Namespace of the secret.<font color="orange"> (Required)</font>  
  
`location` (string)`:` The Azure location.
  
`name` (string)`:` The name of the Azure cloud credential.
  
`subscriptionId` (string)`:` The Azure subscription ID.
  
`tenantId` (string)`:` The Azure tenant ID.
  

## Optional
  
`lock` (boolean)`:` Indicates whether to lock the Azure cloud credential. Defaults to `false`.
  
`organizationId` (string)`:` The ID of the organization which owns the Azure cloud credential.
  
`organizationIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
