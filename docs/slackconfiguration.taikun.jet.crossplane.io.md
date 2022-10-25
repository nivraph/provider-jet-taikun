
slackconfiguration.taikun.jet.crossplane.io
===========================================


This document has been generated.
  

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
  
`channel`: Slack channel for notifications.
  
`name`: The Slack configuration's name.
  
`type`: The type of notifications to receive: `Alert` (only alert-type notifications) or `General` (all notifications).
  
`url`: Webhook URL from Slack app.
  

## Optional
  
`organizationId`: The ID of the organization which owns the Slack configuration.
  
`organizationIdRef`: A Reference to a named object.
  
`organizationIdSelector`: A Selector selects an object.
  
