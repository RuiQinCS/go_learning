make docker
kubectl apply -f k8s-webook-deployment.yaml
kubectl apply -f k8s-webook-service.yaml

# kubectl get deployments
# kubectl get services
# kubectl get pods
# kubectl get namespaces
# 测试地址 : http://localhost:81/ping