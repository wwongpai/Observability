## Infrastructure Monitoring (Metrics Collection)

1. Deploy rbac.yaml
```
kubectl apply -f rbac.yaml
```

2. Deploy the agent as a sidecar [example-eks-fargate_infra-apm.yaml](https://github.com/wwongpai/Observability/tree/main/agent/eks-fargate/infra)https://github.com/wwongpai/Observability/tree/main/agent/eks-fargate/infra
```
kubectl apply -f example-eks-fargate_infra-apm.yaml
```

3. Datadog recommends you run the Cluster Agent to access features such as events collection, Kubernetes resources view, and cluster checks. Deploy the cluster-agent with the following helm chart. Find more [detail](https://docs.datadoghq.com/integrations/eks_fargate/#running-the-cluster-agent-or-the-cluster-checks-runner)
```
helm install <release-name> -f <values.yaml file path> datadog/datadog
```
To be noticed:
- agents.enabled is set to false because we do not want the node agent daemonset deployed because daemonsets are not supported in EKS Fargate, and the node agent is already being deployed as a sidecar to their application
- clusterAgent.token must be a hardcoded value to avoid issues when running helm upgrade
- The cluster agent must match the EKS Fargate profile which uses namespace. The namespace where the cluster agent will be deployed can be specified in the helm install command, For example,
```
helm install datadog-release -f values.yaml datadog/datadog -n custom-namespace
```

4. Add these env to your application pod/deployment file (which already had a sidecar agent)
```
kubectl apply -f example-eks-fargate_infra-apm-clusteragent-connection.yaml
```
To be noticed:
- DD_ORCHESTRATOR_EXPLORER_ENABLED set to true will enable the orchestrator explorer, so pods/depoloyments/etc will be able on the Live Containers page
- DD_CLUSTER_AGENT_AUTH_TOKEN must match the hardcoded token set in the helm chart
- DD_CLUSTER_AGENT_URL is equivalent to https://<helm-install-name>-datadog-cluster-agent.<dca-install-namespace>.svc.cluster.local:5005
For example, if when they installed the helm chart, the release name was dca and they installed the dca in the namespace datadog-fargate their URL will look like https://dca-datadog-cluster-agent.datadog-fargate.svc.cluster.local:5005


5.

