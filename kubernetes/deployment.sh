# web 服务器部署
make docker
kubectl apply -f k8s-webook-deployment.yaml
kubectl apply -f k8s-webook-service.yaml

# kubectl get deployments
# kubectl get services
# kubectl get pods
# kubectl get namespaces
# 测试地址 : http://localhost:81/ping


# redis 部署
kubectl apply -f k8s-redis-deployment.yaml
kubectl apply -f k8s-redis-service.yaml

# 测试 : redis-cli -p 31379
# 查看pod日志 kubectl logs -f {pod 名称}  # pod 名称可以通过 kubectl get pods 获取
# 删除 :  kubectl delete service webook-redis kubectl delete deployment webook-redis


# mysql 部署
kubectl apply -f k8s-mysql-deployment.yaml
kubectl apply -f k8s-mysql-service.yaml
kubectl apply -f k8s-mysql-pvc.yaml
kubectl apply -f k8s-mysql-pv.yaml


# nginx 部署
# 修改 /etc/hosts, 添加 : 127.0.0.1       practice.webook.com
kubectl apply -f k8s-ingress-nginx.yaml

# kubectl get ingress
# kubectl delete ingress webook-ingress
