
kubernetesprofile.taikun.jet.crossplane.io
==========================================


This document has been generated.
  

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
  
`name`: The name of the Kubernetes profile.
  

## Optional
  
`bastionProxy`: Whether to expose the Service on each Node's IP at a static port, the NodePort. You'll be able to contact the NodePort Service, from outside the cluster, by requesting `<NodeIP>:<NodePort>`. Defaults to `false`.
  
`loadBalancingSolution`: The load-balancing solution: `None`, `Octavia` or `Taikun`. `Octavia` and `Taikun` are only available for OpenStack cloud. Defaults to `Octavia`.
  
`lock`: Indicates whether to lock the Kubernetes profile. Defaults to `false`.
  
`organizationId`: The ID of the organization which owns the Kubernetes profile.
  
`organizationIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="red"> (Required)</font>  
  
`organizationIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
`scheduleOnMaster`: When enabled, the workload will also run on master nodes (not recommended). Defaults to `false`.
  
`uniqueClusterName`: If not enabled, the cluster name will be cluster.local. Defaults to `true`.
  
