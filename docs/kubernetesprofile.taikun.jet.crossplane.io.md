
kubernetesprofile.taikun.jet.crossplane.io
==========================================


This document has been generated from the CRD.
  

# Example


```yaml
apiVersion: kubernetesprofile.taikun.jet.crossplane.io/v1alpha1
kind: Profile
metadata:
  name: KUBERNETES_PROFILE
spec:
  forProvider:
    name: "KUBERNETES_PROFILE"
    organizationIdRef:
        name: "ORGANIZATION_REF"
    lock: false
    loadBalancingSolution: "Octavia"
    uniqueClusterName: true
  providerConfigRef:
    name: default

```  

# Schema
  

## Required
  
`name` (string)`:` The name of the Kubernetes profile.
  

## Optional
  
`bastionProxy` (boolean)`:` Whether to expose the Service on each Node's IP at a static port, the NodePort. You'll be able to contact the NodePort Service, from outside the cluster, by requesting `<NodeIP>:<NodePort>`. Defaults to `false`.
  
`loadBalancingSolution` (string)`:` The load-balancing solution: `None`, `Octavia` or `Taikun`. `Octavia` and `Taikun` are only available for OpenStack cloud. Defaults to `Octavia`.
  
`lock` (boolean)`:` Indicates whether to lock the Kubernetes profile. Defaults to `false`.
  
`organizationId` (string)`:` The ID of the organization which owns the Kubernetes profile.
  
`organizationIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
`scheduleOnMaster` (boolean)`:` When enabled, the workload will also run on master nodes (not recommended). Defaults to `false`.
  
`uniqueClusterName` (boolean)`:` If not enabled, the cluster name will be cluster.local. Defaults to `true`.
  
