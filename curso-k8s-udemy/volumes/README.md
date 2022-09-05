`EmptyDir` store the data to pod level, 
When a container is dead, the persistense is store in the pod level
but when the pod is dead the volume too

`HostPath` store the date to node level(in the Cluster)


Aprovicionamiento Dinamico
We can create a Clain and with the storage class we can create it dinamicaly
