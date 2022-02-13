kubectl config set-context docker-desktop

# Install ingress-nginx
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.1.1/deploy/static/provider/cloud/deploy.yaml

# Create sample namespace
# For some reason it doesn't work if we let this be created by .yaml
kubectl create namespace go-skaffold
