
kubeconfig.taikun.jet.crossplane.io
===================================


This document has been generated.
  

# Example


```yaml
apiVersion: kubeconfig.taikun.jet.crossplane.io/v1alpha1
kind: Kubeconfig
metadata:
    name: KUBECONFIG
spec:
  forProvider:
    accessScope: "SCOPE"
    name: "KUBECONFIG"
    role: "ROLE"
    projectId: "PROJECT"
    namespace: "NAMESPACE"
    userId: "USER"
    validityPeriod: VALIDITY
  providerConfigRef:
    name: default

```  

# Schema
  

## Required
  
`accessScope`: Who can use the kubeconfig: `personal` (only you), `managers` (managers only) or `all` (all users with access to this project).
  
`name`: The kubeconfig's name.
  
`role`: The kubeconfig's role: `cluster-admin`, `admin`, `edit` or `view`.
  

## Optional
  
`namespace`: The kubeconfig's namespace.
  
`projectId`: ID of the kubeconfig's project.
  
`projectIdRef`: A Reference to a named object.
  
`projectIdSelector`: A Selector selects an object.
  
`userId`: ID of the kubeconfig's user, if the kubeconfig is personal.
  
`userIdRef`: A Reference to a named object.
  
`userIdSelector`: A Selector selects an object.
  
`validityPeriod`: The kubeconfig's validity period in minutes. Unlimited (-1) by default. Defaults to `-1`.
  
