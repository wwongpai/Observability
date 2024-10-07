# How to install the Datadog Agent on AWS EKS Cluster with Datadog Operator and Instrumenting application using local lib injection

Official document
--------
[https://docs.datadoghq.com/containers/kubernetes/installation?tab=helm](https://docs.datadoghq.com/containers/kubernetes/installation/?tab=datadogoperator)
[https://docs.datadoghq.com/tracing/trace_collection/library_injection_local/?tab=kubernetes](https://docs.datadoghq.com/tracing/trace_collection/library_injection_local/?tab=kubernetes)



Quick set up
--------
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
- Change <Cluster Name> to the name you want
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
    nodeAgent:
      env:
        - name: DD_LOGS_CONFIG_AUTO_MULTI_LINE_DETECTION
          value: "true"
```

4. Deploy the Datadog Agent:
```
kubectl apply -f /path/to/your/datadog-agent.yaml
```






