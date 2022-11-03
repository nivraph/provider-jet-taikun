
slackconfiguration.taikun.jet.crossplane.io
===========================================


This document has been generated from the CRD.
  

# Example


```yaml
apiVersion: slackconfiguration.taikun.jet.crossplane.io/v1alpha1
kind: Configuration
metadata:
  name: SLACK_CONFIGURATION
spec:
  forProvider:
    name: "SLACK_CONFIGURATION"
    organizationIdRef:
        name: "ORGANIZATION_REF"
    channel: "CHANNEL"
    type: "TYPE"
    url: "URL"
  providerConfigRef:
    name: default
```  

# Schema
  

## Required
  
`channel` (string)`:` Slack channel for notifications.
  
`name` (string)`:` The Slack configuration's name.
  
`type` (string)`:` The type of notifications to receive: `Alert` (only alert-type notifications) or `General` (all notifications).
  
`url` (string)`:` Webhook URL from Slack app.
  

## Optional
  
`organizationId` (string)`:` The ID of the organization which owns the Slack configuration.
  
`organizationIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`organizationIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
