
accessprofile.taikun.jet.crossplane.io
======================================


This document has been generated from the CRD.
  

# Example


```yaml
apiVersion: accessprofile.taikun.jet.crossplane.io/v1alpha1
kind: Profile
metadata:
  name: access-profile
spec:
  forProvider:
    name: "access-profile"
    organizationIdRef:
        name: "test-organization"
  providerConfigRef:
    name: default

```  

# Schema
  

## Required
  
`name` (string)`:` The name of the access profile.
  

## Optional
  
`allowedHost` (array)`:` List of allowed hosts.
  
`dnsServer` (array)`:` List of DNS servers.
  
`httpProxy` (string)`:` HTTP proxy of the access profile.
  
`lock` (boolean)`:` Indicates whether to lock the access profile. Defaults to `false`.
  
`ntpServer` (array)`:` List of NTP servers.
  
`organizationId` (string)`:` The ID of the organization which owns the access profile.
  
`organizationIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
`sshUser` (array)`:` List of SSH users.
  
