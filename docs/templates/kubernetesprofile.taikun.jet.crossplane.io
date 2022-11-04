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
