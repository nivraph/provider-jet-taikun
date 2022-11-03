
policyprofile.taikun.jet.crossplane.io
======================================


This document has been generated from the CRD.
  

# Example


```yaml
apiVersion: policyprofile.taikun.jet.crossplane.io/v1alpha1
kind: Profile
metadata:
  name: POLICY_PROFILE
spec:
  forProvider:
    name: "POLICY_PROFILE"
    organizationIdRef:
        name: "ORGANIZATION_REF"
    allowedRepos:
      - "REPO"
    forbiddenTags:
      - "TAG"
    lock: LOCK
    forbidHttpIngress: HTTP
    forbidNodePort: NODE
    requireProbe: PROBE
    uniqueIngress: INGRESS
    uniqueServiceSelector: SELECTOR
  providerConfigRef:
    name: default

```  

# Schema
  

## Required
  
`name` (string)`:` The name of the Policy profile.
  

## Optional
  
`allowedRepos` (array)`:` Requires container images to begin with a string from the specified list.
  
`forbidHttpIngress` (boolean)`:` Requires Ingress resources to be HTTPS only. Defaults to `false`.
  
`forbidNodePort` (boolean)`:` Disallows all Services with type NodePort. Defaults to `false`.
  
`forbiddenTags` (array)`:` Container images must have an image tag different from the ones in the list.
  
`ingressWhitelist` (array)`:` List of allowed Ingress rule hosts.
  
`lock` (boolean)`:` Indicates whether to lock the Policy profile. Defaults to `false`.
  
`organizationId` (string)`:` The ID of the organization which owns the Policy profile.
  
`organizationIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
`requireProbe` (boolean)`:` Requires Pods to have readiness and liveness probes. Defaults to `false`.
  
`uniqueIngress` (boolean)`:` Requires all Ingress rule hosts to be unique. Defaults to `false`.
  
`uniqueServiceSelector` (boolean)`:` Whether services must have globally unique service selectors or not. Defaults to `false`.
  
