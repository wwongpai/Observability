## Infrastructure Monitoring (Metrics Collection)

1. Deploy rbac.yaml
```
kubectl apply -f rbac.yaml
```

2. Deploy the agent as a sidecar [example-eks-fargate_infra-apm.yaml](https://github.com/wwongpai/Observability/tree/main/agent/eks-fargate/infra)
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
- DD_CLUSTER_AGENT_URL is equivalent to https://<helm-install-name>-datadog-cluster-agent.<ddogca-install-namespace>.svc.cluster.local:5005
For example, if when they installed the helm chart, the release name was ddogca and they installed the ddogca in the namespace datadog-fargate their URL will look like https://ddogca-datadog-cluster-agent.datadog-fargate.svc.cluster.local:5005


5. We recommend to use cluster check runners to run cluster checks in an EKS Fargate environment. The Cluster Agent can dispatch out two types of checks: endpoint checks and cluster checks. The checks are slightly different. Endpoint checks are dispatched specifically to the regular Datadog Agent on the same node as the application pod endpoints. Executing endpoint checks on the same node as the application endpoint allows proper tagging of the metrics. Cluster checks monitor internal Kubernetes services, as well as external services like managed databases and network devices, and can be dispatched much more freely. Using Cluster Check Runners is optional. When you use Cluster Check Runners, a small, dedicated set of Agents runs the cluster checks, leaving the endpoint checks to the normal Agent. This strategy can be beneficial to control the dispatching of cluster checks, especially when the scale of your cluster checks increases. To enable this, add the parameter clusterChecksRunner.enabled set to true to the helm chart from above
```
clusterChecksRunner
  enabled: true
```

For example, let’s say I wanted to run the kube_apiserver_metrics check as a cluster check. I would define it under the parameter clusterAgent.confd in my helm chart like so:
```
clusterAgent:
[...]
  confd: 
    kube_apiserver_metrics.yaml: |-
      cluster_check: true
      instances:
        - prometheus_url: https://%%env_KUBERNETES_SERVICE_HOST%%:443/metrics
          bearer_token_auth: true
          ssl_verify: false
```

