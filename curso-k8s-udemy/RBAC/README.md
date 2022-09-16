


`openssl genrsa -out dante.key 2048`

`openssl req -new -key dante.key -out dante.csr -subj "/CN=dante/O=dev"`

`kubectl apply -f 01-dante-pods.yaml`

`kubectl get roles`

`kubectl describe role pod-reader`

`kubectl get rolebindings`


`kubectl describe rolebindings read-pods`