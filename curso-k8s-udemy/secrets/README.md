`kubectl create secret generic mysecret --from-file=secrets-files/test.txt`

Not show the content

`kubectl describe secret mysecret`


Security reason

`kubectl get secrets mysecret -o yaml`
```
apiVersion: v1
data:
  test.txt: c2VjcmV0bzE9aG9sYQpzZWNyZXRvMj1hZGlvcw==
kind: Secret
metadata:
  creationTimestamp: "2022-08-29T03:18:05Z"
  name: mysecret
  namespace: default
  resourceVe´rsion: "573682"
  uid: 7f3365ff-e3c8-4063-acba-efa53241dcdc
type: Opaque
```
you can see the basic secret in base64 c2VjcmV0bzE9aG9sYQpzZWNyZXRvMj1hZGlvcw==

`echo c2VjcmV0bzE9aG9sYQpzZWNyZXRvMj1hZGlvcw== | base64 --decode`
```
secreto1=hola
secreto2=adios%   
```


Dont upload in git the secretes with sensitive data, use placeholders
export PASSWORDTEST=dante
export USER=dante
`envsubst < 03-secrets-secure.yaml > tmp.yaml`

Once to run the last command you can apply with tmp.yaml


 