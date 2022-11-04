
projectuserattachment.taikun.jet.crossplane.io
==============================================


This document has been generated from the CRD.
  

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
  
`projectId` (string)`:` ID of the project.
  
`projectIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`projectIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
`userId` (string)`:` ID of the user.
  
`userIdRef` (object)`:` A Reference to a named object.

* `name` (string)`:` Name of the referenced object.<font color="orange"> (Required)</font>  
  
`userIdSelector` (object)`:` A Selector selects an object.

* `matchControllerRef` (boolean)`:` MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.  

* `matchLabels` (object)`:` MatchLabels ensures an object with matching labels is selected.  
  
