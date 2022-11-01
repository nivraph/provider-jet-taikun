
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
  
`name`: The name of the cloud credential.
  
`type`: The type of the cloud credential.
  

## Optional
  
`accessKeyIdSecretRef`: The AWS access key ID. Required for AWS.

* `key`: The key to select.<font color="orange"> (Required)</font>  

* `name`: Name of the secret.<font color="orange"> (Required)</font>  

* `namespace`: Namespace of the secret.<font color="orange"> (Required)</font>  
  
`availabilityZone`: The availability zone of the cloud credential. Optional for Openstack. Required for AWS and Azure. See `zone` for GCP.
  
`billingAccountId`: The ID of the GCP credential's billing account. Conflicts with: `import_project`.
  
`clientIdSecretRef`: The Azure client ID. Required for Azure.

* `key`: The key to select.<font color="orange"> (Required)</font>  

* `name`: Name of the secret.<font color="orange"> (Required)</font>  

* `namespace`: Namespace of the secret.<font color="orange"> (Required)</font>  
  
`clientSecretSecretRef`: The Azure client secret. Required for Azure.

* `key`: The key to select.<font color="orange"> (Required)</font>  

* `name`: Name of the secret.<font color="orange"> (Required)</font>  

* `namespace`: Namespace of the secret.<font color="orange"> (Required)</font>  
  
`configFile`: The path of the GCP credential's configuration file. Required for GCP.
  
`domain`: The OpenStack domain. Required for Openstack.
  
`folderId`: The folder ID of the GCP credential. Conflicts with: `import_project`.
  
`importProject`: Whether to import a project or not Defaults to `false`. Conflicts with: `billing_account_id`, `folder_id`.
  
`importedNetworkSubnetId`: The OpenStack network subnet ID to import a network.
  
`location`: The Azure location. Required for Azure.
  
`lock`: Indicates whether to lock the cloud credential. Defaults to `false`.
  
`organizationId`: The ID of the organization which owns the cloud credential.
  
`organizationIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
`passwordSecretRef`: The OpenStack password. Required for Openstack.

* `key`: The key to select.<font color="orange"> (Required)</font>  

* `name`: Name of the secret.<font color="orange"> (Required)</font>  

* `namespace`: Namespace of the secret.<font color="orange"> (Required)</font>  
  
`projectName`: The OpenStack project name. Required for Openstack.
  
`publicNetworkName`: The name of the public OpenStack network to use. Required for Openstack.
  
`region`: The region of the cloud credential. Required for Openstack, AWS and GCP.
  
`secretAccessKeySecretRef`: The AWS secret access key. Required for AWS.

* `key`: The key to select.<font color="orange"> (Required)</font>  

* `name`: Name of the secret.<font color="orange"> (Required)</font>  

* `namespace`: Namespace of the secret.<font color="orange"> (Required)</font>  
  
`subscriptionId`: The Azure subscription ID. Required for Azure.
  
`tenantId`: The Azure tenant ID. Required for Azure.
  
`url`: The OpenStack authentication URL. Required for Openstack.
  
`user`: The OpenStack user. Required for Openstack.
  
`volumeTypeName`: The OpenStack type of volume.
  
`zone`: The zone of the GCP credential. Required for GCP.
  
