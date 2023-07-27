## Infrastructure Monitoring (Metrics Collection)

1. Deploy rbac.yaml
```
kubectl apply -f rbac.yaml
```

2. Deploy the agent as a sidecar [example-eks-fargate_infra-apm.yaml](https://github.com/wwongpai/Observability/tree/main/agent/eks-fargate/infra)https://github.com/wwongpai/Observability/tree/main/agent/eks-fargate/infra
```
kubectl apply -f example-eks-fargate_infra-apm.yaml
```

3. Datadog recommends you run the Cluster Agent to access features such as events collection, Kubernetes resources view, and cluster checks. [More detail](https://docs.datadoghq.com/integrations/eks_fargate/#running-the-cluster-agent-or-the-cluster-checks-runner)
Deploy the cluster-agent with the following helm chart
```
helm install <release-name> -f <values.yaml file path> datadog/datadog
```

