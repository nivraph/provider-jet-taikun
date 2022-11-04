
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
  
`domain` (string)`:` The OpenStack domain.
  
`name` (string)`:` The name of the OpenStack cloud credential.
  
`passwordSecretRef` (object)`:` The OpenStack password.

* `key` (string)`:` The key to select.<font color="orange"> (Required)</font>  

* `name` (string)`:` Name of the secret.<font color="orange"> (Required)</font>  

* `namespace` (string)`:` Namespace of the secret.<font color="orange"> (Required)</font>  
  
`projectName` (string)`:` The OpenStack project name.
  
`publicNetworkName` (string)`:` The name of the public OpenStack network to use.
  
`region` (string)`:` The OpenStack region.
  
`url` (string)`:` The OpenStack authentication URL.
  
`user` (string)`:` The OpenStack user.
  

## Optional
  
`availabilityZone` (string)`:` The OpenStack availability zone.
  
`importedNetworkSubnetId` (string)`:` The OpenStack network subnet ID to import a network.
  
`lock` (boolean)`:` Indicates whether to lock the OpenStack cloud credential. Defaults to `false`.
  
`organizationId` (string)`:` The ID of the organization which owns the OpenStack cloud credential.
  
`organizationIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
`volumeTypeName` (string)`:` The OpenStack type of volume.
  
