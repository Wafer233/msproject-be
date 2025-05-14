Write-Output "starting namespace..."
kubectl apply -f namespace.yaml
Write-Output "namespace success..."

Write-Output "starting configmap..."
kubectl apply -f configmaps.yaml
Write-Output "configmap success..."

Write-Output "starting deploys & services..."
kubectl apply -f deployments.yaml
Write-Output "success..."
