
project.taikun.jet.crossplane.io
================================


This document has been generated from the CRD.
  

# Example


```yaml
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

```  

# Schema
  

## Required
  
`name`: Project name.
  

## Optional
  
`accessProfileId`: ID of the project's access profile. Defaults to the default access profile of the project's organization.
  
`accessProfileIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="orange"> (Required)</font>  
  
`accessProfileIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
`alertingProfileId`: ID of the project's alerting profile.
  
`alertingProfileIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="orange"> (Required)</font>  
  
`alertingProfileIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
`autoUpgrade`: If enabled, the Kubespray version will be automatically upgraded when a new version is available. Defaults to `false`.
  
`backupCredentialId`: ID of the backup credential. If unspecified, backups are disabled.
  
`backupCredentialIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="orange"> (Required)</font>  
  
`backupCredentialIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
`cloudCredentialId`: ID of the cloud credential used to create the project's servers.
  
`cloudCredentialIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="orange"> (Required)</font>  
  
`cloudCredentialIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
`deleteOnExpiration`: If enabled, the project will be deleted on the expiration date and it will not be possible to recover it. Defaults to `false`. Required with: `expiration_date`.
  
`expirationDate`: Project's expiration date in the format: 'dd/mm/yyyy'.
  
`flavors`: List of flavors bound to the project.
  
`images`: List of images bound to the project.
  
`kubernetesProfileId`: ID of the project's Kubernetes profile. Defaults to the default Kubernetes profile of the project's organization.
  
`kubernetesProfileIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="orange"> (Required)</font>  
  
`kubernetesProfileIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
`kubernetesVersion`: Kubernetes Version at project creation. Use the meta-argument `ignore_changes` to ignore future upgrades.
  
`lock`: Indicates whether to lock the project. Defaults to `false`.
  
`monitoring`: Kubernetes cluster monitoring. Defaults to `false`.
  
`organizationId`: ID of the organization which owns the project.
  
`organizationIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
`policyProfileId`: ID of the Policy profile. If unspecified, Gatekeeper is disabled.
  
`policyProfileIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="orange"> (Required)</font>  
  
`policyProfileIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
`quotaCpuUnits`: Maximum CPU units. Defaults to `1000000`.
  
`quotaDiskSize`: Maximum disk size in GBs. Defaults to `102400`.
  
`quotaRamSize`: Maximum RAM size in GBs. Defaults to `102400`.
  
`quotaVmCpuUnits`: Maximum CPU units for standalone VMs. Defaults to `1000000`.
  
`quotaVmRamSize`: Maximum RAM size in GBs for standalone VMs. Defaults to `102400`.
  
`quotaVmVolumeSize`: Maximum volume size in GBs for standalone VMs. Defaults to `102400`.
  
`routerIdEndRange`: Router ID end range (specify only if using OpenStack cloud credentials with Taikun Load Balancer enabled). Required with: `router_id_start_range`, `taikun_lb_flavor`.
  
`routerIdStartRange`: Router ID start range (specify only if using OpenStack cloud credentials with Taikun Load Balancer enabled). Required with: `router_id_end_range`, `taikun_lb_flavor`.
  
`serverBastion`: Bastion server. Required with: `server_kubemaster`, `server_kubeworker`.
  
`serverKubemaster`: Kubemaster server. Required with: `server_bastion`, `server_kubeworker`.
  
`serverKubeworker`: Kubeworker server. Required with: `server_bastion`, `server_kubemaster`.
  
`taikunLbFlavor`: OpenStack flavor for the Taikun load balancer (specify only if using OpenStack cloud credentials with Taikun Load Balancer enabled). Required with: `router_id_end_range`, `router_id_start_range`.
  
`vm`: Virtual machines.
  
