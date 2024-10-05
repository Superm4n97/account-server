# keycloak

## kind setup
create kind cluster with below config file
```yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
name: hbr-cluster
nodes:
- role: control-plane
  kubeadmConfigPatches:
  - |
    kind: InitConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: "ingress-ready=true"
  extraPortMappings:
  - containerPort: 80
    hostPort: 80
    protocol: TCP
  - containerPort: 443
    hostPort: 443
    protocol: TCP
- role: worker
- role: worker
```
**Make sure you have enough host resources and no other kind cluster is active.** 
```bash
kind create cluster --name kind-a --config=config.yml
```

## install nginx ingress
```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
```

## install keycloak helm chart
```bash
helm install keycloak keycloak
```