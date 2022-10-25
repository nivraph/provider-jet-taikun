
policyprofile.taikun.jet.crossplane.io
======================================


This document has been generated.
  

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
  
`name`: The name of the Policy profile.
  

## Optional
  
`allowedRepos`: Requires container images to begin with a string from the specified list.
  
`forbidHttpIngress`: Requires Ingress resources to be HTTPS only. Defaults to `false`.
  
`forbidNodePort`: Disallows all Services with type NodePort. Defaults to `false`.
  
`forbiddenTags`: Container images must have an image tag different from the ones in the list.
  
`ingressWhitelist`: List of allowed Ingress rule hosts.
  
`lock`: Indicates whether to lock the Policy profile. Defaults to `false`.
  
`organizationId`: The ID of the organization which owns the Policy profile.
  
`organizationIdRef`: A Reference to a named object.
  
`organizationIdSelector`: A Selector selects an object.
  
`requireProbe`: Requires Pods to have readiness and liveness probes. Defaults to `false`.
  
`uniqueIngress`: Requires all Ingress rule hosts to be unique. Defaults to `false`.
  
`uniqueServiceSelector`: Whether services must have globally unique service selectors or not. Defaults to `false`.
  
