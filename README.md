# declarations-api

Random declarations about myself from the Bible or someone whose mind is steeped in the Bible.  Currently stored in a text file.

```
# Build the image
docker build -t declarations-api:latest .

# Run the image in docker:
docker run -it --rm --name declarations-api -p 8080:8080 declarations-api:latest

# Test it out:
curl http://localhost:8080/api/declaration/random

# Push it up to the microk8s registry
docker tag declarations-api:latest localhost:32000/declarations-api:latest .
docker push localhost:32000/declarations-api:latest

#############################################################
# Note that my kubectl is aliased to 'microk8s kubectl'
#############################################################

# Tell it to pick up the new one
kubectl delete pod/declarations-api-abcdefg-1234567

# Test the deployment web interface
kubectl port-forward deploy/declarations-api  8080:8080

# Test the service web interface
kubectl port-forward svc/declarations-api  8080:8080

# Test it out
kubectl exec -it $(kubectl get pod | grep declarations-api | head -n1 | awk '{print $1}') -- bash

# Copy a new declarations file up to the persistent volume
# Interesting... my setup doesn't like absolute file paths 
kubectl cp ./static/declarations declarations-api-5994dff79b-jxt5t:./declarations/
```
