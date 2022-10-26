
alertingprofile.taikun.jet.crossplane.io
========================================


This document has been generated.
  

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
  
`name`: The alerting profile's name.
  
`reminder`: The frequency of notifications: `HalfHour`, `Hourly`, `Daily` or `None`.
  

## Optional
  
`emails`: The list of emails to notify.
  
`integration`: List of alerting integrations.
  
`lock`: Indicates whether to lock the profile. Defaults to `false`.
  
`organizationId`: The ID of the organization which owns the profile.
  
`organizationIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="red"> (Required)</font>  
  
`organizationIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
`slackConfigurationId`: The ID of the Slack configuration to notify. Defaults to `0`.
  
`slackConfigurationIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="red"> (Required)</font>  
  
`slackConfigurationIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
`webhook`: The list of webhooks to notify.
  
