# declarations-api

Random declarations about myself from the Bible or those who are steeped in the Bible.  Currently stored in a text file.

```
# Build the image
docker build -t declarations-api:latest -t localhost:32000/declarations-api:latest .

# Run the image in docker:
docker run -it --rm --name declare -p 8080:8080 declarations-api:latest

# Test it out:
curl http://localhost:8080/api/declaration/random

# Push it up to the k8s registry
docker push localhost:32000/declarations-api:latest

# Tell it to pick up the new one
kubectl delete pod/declarations-api-abcdefg-1234567

# Test the deployment web interface
kubectl port-forward deploy/declarations-api  8080:8080

# Test the service web interface
kubectl port-forward svc/declarations-api  8080:8080

# Test it out
kubectl exec -it $(kubectl get pod | grep declarations-api | head -n1 | awk '{print $1}') -- bash



```
