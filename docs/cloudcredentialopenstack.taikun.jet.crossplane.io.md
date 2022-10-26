
cloudcredentialopenstack.taikun.jet.crossplane.io
=================================================


This document has been generated from the CRD.
  

# Example


```yaml
apiVersion: cloudcredentialopenstack.taikun.jet.crossplane.io/v1alpha1
kind: CredentialOpenstack
metadata:
  name: OPENSTACK
spec:
  forProvider:
    name: "OPENSTACK"
    user: "USER"
    passwordSecretRef:
      key: "password"
      name: "openstack-credentials"
      namespace: "crossplane-system"
    organizationIdRef:
        name: "ORGANIZATION_REF"
    url: "URL"
    domain: "DOMAIN"
    projectName: "PROJECT"
    region: "REGION"
    publicNetworkName: "PUBLIC_NETWORK"
    lock: LOCK
    availabilityZone: "AVAILABILITY"
```  

# Schema
  

## Required
  
`domain`: The OpenStack domain.
  
`name`: The name of the OpenStack cloud credential.
  
`passwordSecretRef`: The OpenStack password.

* `key`: The key to select.<font color="orange"> (Required)</font>  

* `name`: Name of the secret.<font color="orange"> (Required)</font>  

* `namespace`: Namespace of the secret.<font color="orange"> (Required)</font>  
  
`projectName`: The OpenStack project name.
  
`publicNetworkName`: The name of the public OpenStack network to use.
  
`region`: The OpenStack region.
  
`url`: The OpenStack authentication URL.
  
`user`: The OpenStack user.
  

## Optional
  
`availabilityZone`: The OpenStack availability zone.
  
`importedNetworkSubnetId`: The OpenStack network subnet ID to import a network.
  
`lock`: Indicates whether to lock the OpenStack cloud credential. Defaults to `false`.
  
`organizationId`: The ID of the organization which owns the OpenStack cloud credential.
  
`organizationIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
`volumeTypeName`: The OpenStack type of volume.
  
