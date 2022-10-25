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
