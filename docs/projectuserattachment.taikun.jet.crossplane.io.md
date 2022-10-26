
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

* `name`: Name of the referenced object.<font color="red"> (Required)</font>  
  
`projectIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
`userId`: ID of the user.
  
`userIdRef`: A Reference to a named object.

* `name`: Name of the referenced object.<font color="red"> (Required)</font>  
  
`userIdSelector`: A Selector selects an object.

* `matchControllerRef`: MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels`: MatchLabels ensures an object with matching labels is selected.  
  
