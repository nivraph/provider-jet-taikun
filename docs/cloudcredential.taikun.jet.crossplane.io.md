
cloudcredential.taikun.jet.crossplane.io
========================================


This document has been generated from the CRD.
  

# Example


```yaml
apiVersion: cloudcredential.taikun.jet.crossplane.io/v1alpha1
kind: Credential
metadata:
  name: cloud-credential
spec:
  forProvider:
    name: "cloud-credential"
    type: "openstack"
    user: "user"
    passwordSecretRef:
      key: "password"
      name: "openstack-credentials"
      namespace: "namespace"
    organizationIdRef:
        name: "organization-name"
    url: "url"
    domain: "domain"
    projectName: "project"
    region: "region"
    publicNetworkName: "public-network"
    lock: true
    availabilityZone: "az-2"

```  

# Schema
  

## Required
  
`name` (string)`:` The name of the cloud credential.
  
`type` (string)`:` The type of the cloud credential.
  

## Optional
  
`accessKeyIdSecretRef` (object)`:` The AWS access key ID. Required for AWS.

* `key` (string)`:` The key to select.<font color="orange"> (Required)</font>  

* `name` (string)`:` Name of the secret.<font color="orange"> (Required)</font>  

* `namespace` (string)`:` Namespace of the secret.<font color="orange"> (Required)</font>  
  
`availabilityZone` (string)`:` The availability zone of the cloud credential. Optional for Openstack. Required for AWS and Azure. See `zone` for GCP.
  
`billingAccountId` (string)`:` The ID of the GCP credential's billing account. Conflicts with: `import_project`.
  
`clientIdSecretRef` (object)`:` The Azure client ID. Required for Azure.

* `key` (string)`:` The key to select.<font color="orange"> (Required)</font>  

* `name` (string)`:` Name of the secret.<font color="orange"> (Required)</font>  

* `namespace` (string)`:` Namespace of the secret.<font color="orange"> (Required)</font>  
  
`clientSecretSecretRef` (object)`:` The Azure client secret. Required for Azure.

* `key` (string)`:` The key to select.<font color="orange"> (Required)</font>  

* `name` (string)`:` Name of the secret.<font color="orange"> (Required)</font>  

* `namespace` (string)`:` Namespace of the secret.<font color="orange"> (Required)</font>  
  
`configFile` (string)`:` The path of the GCP credential's configuration file. Required for GCP.
  
`domain` (string)`:` The OpenStack domain. Required for Openstack.
  
`folderId` (string)`:` The folder ID of the GCP credential. Conflicts with: `import_project`.
  
`importProject` (boolean)`:` Whether to import a project or not Defaults to `false`. Conflicts with: `billing_account_id`, `folder_id`.
  
`importedNetworkSubnetId` (string)`:` The OpenStack network subnet ID to import a network.
  
`location` (string)`:` The Azure location. Required for Azure.
  
`lock` (boolean)`:` Indicates whether to lock the cloud credential. Defaults to `false`.
  
`organizationId` (string)`:` The ID of the organization which owns the cloud credential.
  
`organizationIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
`passwordSecretRef` (object)`:` The OpenStack password. Required for Openstack.

* `key` (string)`:` The key to select.<font color="orange"> (Required)</font>  

* `name` (string)`:` Name of the secret.<font color="orange"> (Required)</font>  

* `namespace` (string)`:` Namespace of the secret.<font color="orange"> (Required)</font>  
  
`projectName` (string)`:` The OpenStack project name. Required for Openstack.
  
`publicNetworkName` (string)`:` The name of the public OpenStack network to use. Required for Openstack.
  
`region` (string)`:` The region of the cloud credential. Required for Openstack, AWS and GCP.
  
`secretAccessKeySecretRef` (object)`:` The AWS secret access key. Required for AWS.

* `key` (string)`:` The key to select.<font color="orange"> (Required)</font>  

* `name` (string)`:` Name of the secret.<font color="orange"> (Required)</font>  

* `namespace` (string)`:` Namespace of the secret.<font color="orange"> (Required)</font>  
  
`subscriptionId` (string)`:` The Azure subscription ID. Required for Azure.
  
`tenantId` (string)`:` The Azure tenant ID. Required for Azure.
  
`url` (string)`:` The OpenStack authentication URL. Required for Openstack.
  
`user` (string)`:` The OpenStack user. Required for Openstack.
  
`volumeTypeName` (string)`:` The OpenStack type of volume.
  
`zone` (string)`:` The zone of the GCP credential. Required for GCP.
  
