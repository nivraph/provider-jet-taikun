
projectuserattachment.taikun.jet.crossplane.io
==============================================


This document has been generated.
  

# Example


```yaml
apiVersion: projectuserattachment.taikun.jet.crossplane.io/v1alpha1
kind: UserAttachment
metadata:
  name: USER_ATTACH
spec:
  forProvider:
    userId: "USER_REF"
    projectId: "PROJECT_REF"
```  

# Schema
  

## Optional
  
`projectId`: ID of the project.
  
`projectIdRef`: A Reference to a named object.
  
`projectIdSelector`: A Selector selects an object.
  
`userId`: ID of the user.
  
`userIdRef`: A Reference to a named object.
  
`userIdSelector`: A Selector selects an object.
  
