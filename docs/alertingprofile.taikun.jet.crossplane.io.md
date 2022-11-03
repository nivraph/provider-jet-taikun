
alertingprofile.taikun.jet.crossplane.io
========================================


This document has been generated from the CRD.
  

# Example


```yaml
apiVersion: alertingprofile.taikun.jet.crossplane.io/v1alpha1
kind: Profile
metadata:
  name: alerting-profile
spec:
  forProvider:
    name: "test-alerting-profile"
    emails:
      - "foo@bar.com"
    lock: false
    organizationIdRef:
      name: "test-organization"
    reminder: "Daily"
  providerConfigRef:
    name: default

```  

# Schema
  

## Required
  
`name` (string)`:` The alerting profile's name.
  
`reminder` (string)`:` The frequency of notifications: `HalfHour`, `Hourly`, `Daily` or `None`.
  

## Optional
  
`emails` (array)`:` The list of emails to notify.
  
`integration` (array)`:` List of alerting integrations.
  
`lock` (boolean)`:` Indicates whether to lock the profile. Defaults to `false`.
  
`organizationId` (string)`:` The ID of the organization which owns the profile.
  
`organizationIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
`slackConfigurationId` (string)`:` The ID of the Slack configuration to notify. Defaults to `0`.
  
`slackConfigurationIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`slackConfigurationIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
`webhook` (array)`:` The list of webhooks to notify.
  
