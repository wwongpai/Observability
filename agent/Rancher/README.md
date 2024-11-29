# How to install the Datadog Agent on Rancher Cluster with Datadog Operator

Official document
--------
https://docs.datadoghq.com/containers/kubernetes/distributions/?tab=datadogoperator#Rancher

Rancher installations are similar to vanilla Kubernetes installations, requiring only some minor configuration:
- Tolerations are required to schedule the Node Agent on controlplane and etcd nodes.
- The cluster name should be set as it cannot be retrieved automatically from the cloud provider.

Quick set up Datadog Operator on Rancher
[https://docs.datadoghq.com/containers/kubernetes/installation?tab=helm](https://docs.datadoghq.com/containers/kubernetes/installation/?tab=datadogoperator)

1. Install the Datadog Operator with Helm:
```
helm repo add datadog https://helm.datadoghq.com
helm install datadog-operator datadog/datadog-operator
```

2. Create a Kubernetes secret with your API key:
Create API key and App key in Datadog following [this link](https://docs.datadoghq.com/account_management/api-app-keys)
```
kubectl create secret generic datadog-secret --from-literal api-key=<DATADOG_API_KEY>
```

3. Create a datadog-agent.yaml file with the spec of your DatadogAgent deployment configuration. The following sample configuration enables metrics, logs, events, npm, apm and enabling multi-line log detection:
- Change a cluster name to the name you want
- For more configuration option, please check [this link](https://github.com/DataDog/datadog-operator/blob/main/docs/configuration.v2alpha1.md)
```
apiVersion: datadoghq.com/v2alpha1
kind: DatadogAgent
metadata:
  name: datadog
spec:
  global:
    clusterName: <Cluster Name>
    credentials:
      apiSecret:
        secretName: datadog-secret
        keyName: api-key
    kubelet:
      tlsVerify: false
  features:
    logCollection:
      enabled: true
      containerCollectAll: true
    liveProcessCollection:
      enabled: true
    apm:
      enabled: true
    cspm:
      enabled: true
    npm:
      enabled: true
    usm:
      enabled: true
    admissionController:
      enabled: true
  override:
    clusterAgent:
      image:
        name: gcr.io/datadoghq/cluster-agent:latest
    nodeAgent:
      image:
        name: gcr.io/datadoghq/agent:latest
      tolerations:
        - key: node-role.kubernetes.io/controlplane
          operator: Exists
          effect: NoSchedule
        - key: node-role.kubernetes.io/etcd
          operator: Exists
          effect: NoExecute
      env:
        - name: DD_LOGS_CONFIG_AUTO_MULTI_LINE_DETECTION
          value: "true"
        - name: DD_CONTAINER_EXCLUDE
          value: "kube_namespace:kube-system"
```


4. Deploy the Datadog Agent:
```
kubectl apply -f /path/to/your/datadog-agent.yaml
```

5. Checking if daemonset and deployment work properly:
```
kubectl get ds
kubectl get deploy
kubectl get pods
```

