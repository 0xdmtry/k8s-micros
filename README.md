
# Microservices + Kubernetes
This project contains two microservices: `a-micro` and `b-micro`. 
The services can be built and run using Docker, managed locally with `docker compose`, and deployed to a Kubernetes cluster on AWS using `eksctl` and `kubectl`.

## Docker
### a-micro
Build and run locally:
````
docker build -t a-micro .
docker run -p 8080:8080 a-micro
````

Push to Docker Hub:
````
docker build -t <docker_hub>/a-micro ./a-micro
docker push <docker_hub>/a-micro
````

### b-micro
Build and run locally:
````
docker build -t b-micro .
docker run -p 8081:8081 b-micro
````

Push to Docker Hub:
```
docker build -t b-micro .
docker tag b-micro:latest <docker_hub>/b-micro:v1.0.3
docker push <docker_hub>/b-micro:v1.0.3
```

## Docker Compose
````
docker compose up --build
````

## Kubernetes (EKS)

Create and delete EKS cluster
````
eksctl create cluster --name k8s-micros-002 --region us-east-1 --zones us-east-1a,us-east-1b
eksctl delete cluster --name k8s-micros-001
````

### Kubernetes Deployments
Apply `b-micro` configs
````
kubectl apply -f b.yaml
kubectl apply -f b-micro-deployment.yaml
kubectl apply -f b-micro-service.yaml
kubectl apply -f b-micro-hpa.yaml
````
To delete them:
```
kubectl delete hpa b-micro-hpa
kubectl delete svc b-micro
kubectl delete deployment b-micro
```

Apply `a-micro` configs
````
kubectl apply -f a-micro-deployment.yaml
kubectl apply -f a-micro-service.yaml
kubectl apply -f a-micro-hpa.yaml
````

### Useful `kubectl` Commands

````
kubectl exec -it b-micro-deployment-***********-***** -c **************************************************************** -- /bin/bash
````

````
kubectl describe pod b-micro-deployment-**********-*****
````

````
docker pull <docker_hub>/b-micro
docker run -d -p 8081:8081 <docker_hub>/b-micro
````

````
kubectl logs b-micro-***********-*****
````

````
kubectl exec -it b-micro-deployment-***********-***** -- /bin/sh
````
