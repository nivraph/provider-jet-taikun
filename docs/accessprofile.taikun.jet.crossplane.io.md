
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
  
`name`: The name of the access profile.
  

## Optional
  
`allowedHost`: List of allowed hosts.
  
`dnsServer`: List of DNS servers.
  
`httpProxy`: HTTP proxy of the access profile.
  
`lock`: Indicates whether to lock the access profile. Defaults to `false`.
  
`ntpServer`: List of NTP servers.
  
`organizationId`: The ID of the organization which owns the access profile.
  
`organizationIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
`sshUser`: List of SSH users.
  
