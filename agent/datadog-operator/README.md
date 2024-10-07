# How to install the Datadog Agent on AWS EKS Cluster with Datadog Operator and Instrumenting apps using local lib injection

Official document
--------
[https://docs.datadoghq.com/containers/kubernetes/installation?tab=helm](https://docs.datadoghq.com/containers/kubernetes/installation/?tab=datadogoperator)
[https://docs.datadoghq.com/tracing/trace_collection/library_injection_local/?tab=kubernetes](https://docs.datadoghq.com/tracing/trace_collection/library_injection_local/?tab=kubernetes)


Quick set up Datadog Operator on EKS
--------
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

5. Checking if daemonset and deployment work properly:
```
kubectl get ds
kubectl get deploy
kubectl get pods
```

APM Instrumentation with lib injection
--------
[https://docs.datadoghq.com/tracing/trace_collection/library_injection_local/?tab=kubernetes](https://docs.datadoghq.com/tracing/trace_collection/library_injection_local/?tab=kubernetes)

1. Enable Datadog Admission Controller to mutate app pods and also annotate them for lib injection
```
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    # (...)
spec:
  template:
    metadata:
      labels:
        admission.datadoghq.com/enabled: "true" # Enable Admission Controller to mutate new pods in this deployment
      annotations:
        admission.datadoghq.com/java-lib.version: "<CONTAINER IMAGE TAG>"
    spec:
      containers:
        - # (...)
```

2. Tag your app pods in deployment and pod spec with [Unified Service Tags](https://docs.datadoghq.com/getting_started/tagging/unified_service_tagging/?tab=kubernetes)
```
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    tags.datadoghq.com/env: "prod" # Unified service tag - Deployment Env tag
    tags.datadoghq.com/service: "my-service" # Unified service tag - Deployment Service tag
    tags.datadoghq.com/version: "1.1" # Unified service tag - Deployment Version tag
  # (...)
spec:
  template:
    metadata:
      labels:
        tags.datadoghq.com/env: "prod" # Unified service tag - Pod Env tag
        tags.datadoghq.com/service: "my-service" # Unified service tag - Pod Service tag
        tags.datadoghq.com/version: "1.1" # Unified service tag - Pod Version tag
        admission.datadoghq.com/enabled: "true" # Enable Admission Controller to mutate new pods part of this deployment
      annotations:
        admission.datadoghq.com/java-lib.version: "<CONTAINER IMAGE TAG>"
    spec:
      containers:
        - # (...)
```

3. Inject trace_id and span_id into your logs to correlate with trace, you need to add DD_LOGS_INJECTION through the environment variable.
```
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    tags.datadoghq.com/env: "prod" # Unified service tag - Deployment Env tag
    tags.datadoghq.com/service: "my-service" # Unified service tag - Deployment Service tag
    tags.datadoghq.com/version: "1.1" # Unified service tag - Deployment Version tag
  # (...)
spec:
  template:
    metadata:
      labels:
        tags.datadoghq.com/env: "prod" # Unified service tag - Pod Env tag
        tags.datadoghq.com/service: "my-service" # Unified service tag - Pod Service tag
        tags.datadoghq.com/version: "1.1" # Unified service tag - Pod Version tag
        admission.datadoghq.com/enabled: "true" # Enable Admission Controller to mutate new pods part of this deployment
      annotations:
        admission.datadoghq.com/java-lib.version: "<CONTAINER IMAGE TAG>"
    spec:
      containers:
        - name: <CONTAINER_NAME>
          image: <CONTAINER_IMAGE>/<TAG>
          env:
            - name: DD_LOGS_INJECTION
              value: "true"
```

4. Apply the configuration.
```
kubectl apply -f deployment.yaml
```

Learn more example from [this link](https://github.com/wwongpai/Observability/tree/main/Injecting%20Libraries/kubernetes)
