
kubeconfig.taikun.jet.crossplane.io
===================================


This document has been generated from the CRD.
  

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
  
`accessScope` (string)`:` Who can use the kubeconfig: `personal` (only you), `managers` (managers only) or `all` (all users with access to this project).
  
`name` (string)`:` The kubeconfig's name.
  
`role` (string)`:` The kubeconfig's role: `cluster-admin`, `admin`, `edit` or `view`.
  

## Optional
  
`namespace` (string)`:` The kubeconfig's namespace.
  
`projectId` (string)`:` ID of the kubeconfig's project.
  
`projectIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`projectIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
`userId` (string)`:` ID of the kubeconfig's user, if the kubeconfig is personal.
  
`userIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`userIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
`validityPeriod` (number)`:` The kubeconfig's validity period in minutes. Unlimited (-1) by default. Defaults to `-1`.
  
