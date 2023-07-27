Running the Cluster Agent or the Cluster Checks Runner
Datadog recommends you run the Cluster Agent to access features such as events collection, Kubernetes resources view, and cluster checks.

When using EKS Fargate, there are two possible scenarios depending on whether or not the EKS cluster is running mixed workloads (Fargate/non-Fargate).

If the EKS cluster runs Fargate and non-Fargate workloads, and you want to monitor the non-Fargate workload through Node Agent DaemonSet, add the Cluster Agent/Cluster Checks Runner to this deployment. For more information, see the Cluster Agent Setup.

The Cluster Agent token must be reachable from the Fargate tasks you want to monitor. If you are using the Helm Chart or Datadog Operator, this is not reachable by default because a secret in the target namespace is created.

You have two options for this to work properly:

Use an hardcoded token value (clusterAgent.token in Helm, credentials.token in the Datadog Operator); convenient, but less secure.
Use a manually-created secret (clusterAgent.tokenExistingSecret in Helm, not available in the Datadog Operator) and replicate it in all namespaces where Fargate tasks need to be monitored; secure, but requires extra operations.
If the EKS cluster runs only Fargate workloads, you need a standalone Cluster Agent deployment. And, as described above, choose one of the two options for making the token reachable.

Use the following Helm values.yaml:
--------
[Official doc](https://docs.datadoghq.com/integrations/eks_fargate/#running-the-cluster-agent-or-the-cluster-checks-runner)

```
datadog:
  apiKey: <YOUR_DATADOG_API_KEY>
  clusterName: <CLUSTER_NAME>
agents:
  enabled: false
clusterAgent:
  enabled: true
  replicas: 2
```

  In both cases, you need to change the Datadog Agent sidecar manifest in order to allow communication with the Cluster Agent:
  --------
  ```
         env:
        - name: DD_CLUSTER_AGENT_ENABLED
          value: "true"
        - name: DD_CLUSTER_AGENT_AUTH_TOKEN
          value: <hardcoded token value> # Use valueFrom: if you're using a secret
        - name: DD_CLUSTER_AGENT_URL
          value: https://<CLUSTER_AGENT_SERVICE_NAME>.<CLUSTER_AGENT_SERVICE_NAMESPACE>.svc.cluster.local:5005
        - name: DD_ORCHESTRATOR_EXPLORER_ENABLED # Required to get Kubernetes resources view
          value: "true"
        - name: DD_CLUSTER_NAME
          value: <CLUSTER_NAME>
```
