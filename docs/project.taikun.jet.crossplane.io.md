
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
  
`name` (string)`:` Project name.
  

## Optional
  
`accessProfileId` (string)`:` ID of the project's access profile. Defaults to the default access profile of the project's organization.
  
`accessProfileIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`accessProfileIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
`alertingProfileId` (string)`:` ID of the project's alerting profile.
  
`alertingProfileIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`alertingProfileIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
`autoUpgrade` (boolean)`:` If enabled, the Kubespray version will be automatically upgraded when a new version is available. Defaults to `false`.
  
`backupCredentialId` (string)`:` ID of the backup credential. If unspecified, backups are disabled.
  
`backupCredentialIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`backupCredentialIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
`cloudCredentialId` (string)`:` ID of the cloud credential used to create the project's servers.
  
`cloudCredentialIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`cloudCredentialIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
`deleteOnExpiration` (boolean)`:` If enabled, the project will be deleted on the expiration date and it will not be possible to recover it. Defaults to `false`. Required with: `expiration_date`.
  
`expirationDate` (string)`:` Project's expiration date in the format: 'dd/mm/yyyy'.
  
`flavors` (array)`:` List of flavors bound to the project.
  
`images` (array)`:` List of images bound to the project.
  
`kubernetesProfileId` (string)`:` ID of the project's Kubernetes profile. Defaults to the default Kubernetes profile of the project's organization.
  
`kubernetesProfileIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`kubernetesProfileIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
`kubernetesVersion` (string)`:` Kubernetes Version at project creation. Use the meta-argument `ignore_changes` to ignore future upgrades.
  
`lock` (boolean)`:` Indicates whether to lock the project. Defaults to `false`.
  
`monitoring` (boolean)`:` Kubernetes cluster monitoring. Defaults to `false`.
  
`organizationId` (string)`:` ID of the organization which owns the project.
  
`organizationIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
`policyProfileId` (string)`:` ID of the Policy profile. If unspecified, Gatekeeper is disabled.
  
`policyProfileIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`policyProfileIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
`quotaCpuUnits` (number)`:` Maximum CPU units. Defaults to `1000000`.
  
`quotaDiskSize` (number)`:` Maximum disk size in GBs. Defaults to `102400`.
  
`quotaRamSize` (number)`:` Maximum RAM size in GBs. Defaults to `102400`.
  
`quotaVmCpuUnits` (number)`:` Maximum CPU units for standalone VMs. Defaults to `1000000`.
  
`quotaVmRamSize` (number)`:` Maximum RAM size in GBs for standalone VMs. Defaults to `102400`.
  
`quotaVmVolumeSize` (number)`:` Maximum volume size in GBs for standalone VMs. Defaults to `102400`.
  
`routerIdEndRange` (number)`:` Router ID end range (specify only if using OpenStack cloud credentials with Taikun Load Balancer enabled). Required with: `router_id_start_range`, `taikun_lb_flavor`.
  
`routerIdStartRange` (number)`:` Router ID start range (specify only if using OpenStack cloud credentials with Taikun Load Balancer enabled). Required with: `router_id_end_range`, `taikun_lb_flavor`.
  
`serverBastion` (array)`:` Bastion server. Required with: `server_kubemaster`, `server_kubeworker`.
  
`serverKubemaster` (array)`:` Kubemaster server. Required with: `server_bastion`, `server_kubeworker`.
  
`serverKubeworker` (array)`:` Kubeworker server. Required with: `server_bastion`, `server_kubemaster`.
  
`taikunLbFlavor` (string)`:` OpenStack flavor for the Taikun load balancer (specify only if using OpenStack cloud credentials with Taikun Load Balancer enabled). Required with: `router_id_end_range`, `router_id_start_range`.
  
`vm` (array)`:` Virtual machines.
  
