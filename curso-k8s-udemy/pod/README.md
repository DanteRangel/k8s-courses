# to create a pod via command
`kubectl run podtest --image=nginx:alpine`
# to delete
`kubectl delete pod podtest`


# to get all pods
`kubectl get pods`



# describe a pod
`kubectl describe pod podtest`

# logs in the pod
`kubectl logs podtest`

`kubectl logs podtest -f  #to follow the logs`



# to export the yaml pod
`kubectl get pod podtest -o yaml`

# for forward a ip
`kubectl port-forward podtest 7000:80`

# for exec commands into a pod
`kubectl exec -it podtest -- sh`

# to create  resources via manifest


`kubectl apply -f 01-pod.yaml`

# to delete resources with manifest

`kubectl delete -f 01-pod.yaml`


# filter pods by labels
`kubectl get pods -l app=backend`
`kubectl get pods -l app=frontend`
`kubectl get pods -l env=dev`