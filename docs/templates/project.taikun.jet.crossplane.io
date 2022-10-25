apiVersion: project.taikun.jet.crossplane.io/v1alpha1
kind: Project
metadata:
  name: PROJECT_NAME
spec:
  forProvider:
    name: "PROJECT_NAME"
    cloudCredentialId: "CLOUD_REF"
    organizationIdRef:
        name: "ORGANIZATION_REF"
    quotaDiskSize: QDISK_SIZE
    quotaRamSize: QRAM_SIZE
    lock: false
    monitoring: false
    kubernetesProfileIdRef:
        name: "KP_REF"
    backupCredentialIdRef:
        name: "BACK_REF"
    flavors:
      - "PROJECT_FLAVOR"
    images:
      - "IMG"
    vm:
      - flavor: "VM_FLAVOR"
        imageId: "VM_IMAGE_ID"
        name: "VM_NAME"
        standaloneProfileId: "VM_SA"
        volumeSize: VM_VOLUME_SIZE
        username: "VM_USER"
        publicIp: true
  providerConfigRef:
    name: default
