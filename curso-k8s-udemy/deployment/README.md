# Deployment

### API resource
`apps/v1`
`kubectl api-resources | grep Deployment`

### COMMANDS
`kubectl api-resources | grep Deployment`

 ### to check deployment with manifest
`kubectl rollout status -f 01-deployment.yaml` 

 ### to check deployment via name
`kubectl rollout status deployment <deployment-test>` 

 ### list pods with labels
`kubectl get pods --show-labels`

 ### Describe deployment
`kubectl describe deployment deployment-test`

 ### get the manifest
`kubectl get deployment deployment-test -o yaml`

 ### to run a rollout
`kubectl rollout status deployment deployment-test`

### to add a cause to the deployment
`kubectl apply -f 01-deployment.yaml  --record`
or Add `annotations:` with `kubernetes.io/change-cause: "Change port to 100"` as value

### to check a specific revision
`kubectl rollout history deployment deployment-test --revision=3`

### to do a rollback

`kubectl rollout undo deployment deployment-test --to-revision=3`
