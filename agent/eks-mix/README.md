# Monitor EKS Cluster with a mix of EC2 and Fargate workloads

In this scenario, the EKS cluster hosts workloads running on both EC2 instances and Fargate. The cluster includes multiple namespaces: one for the default settings, one for applications running on Fargate, one for applications on EC2, and another dedicated to Datadog monitoring.

Official document
---------  
- [https://docs.datadoghq.com/containers/kubernetes/installation?tab=helm](https://docs.datadoghq.com/containers/kubernetes/installation?tab=datadogoperator)
- [https://docs.datadoghq.com/integrations/eks_fargate/?tab=datadogoperator](https://docs.datadoghq.com/integrations/eks_fargate/?tab=datadogoperator)


Quick set up
--------
1. Create RBAC for Agent that run as a sidecar in AWS EKS Fargate, Refer to [Getting Started with Datadog ](https://docs.datadoghq.com/integrations/eks_fargate/?tab=datadogoperator#aws-eks-fargate-rbac)
```
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: datadog-agent
rules:
  - apiGroups:
    - ""
    resources:
    - nodes
    - namespaces
    - endpoints
    verbs:
    - get
    - list
  - apiGroups:
      - ""
    resources:
      - nodes/metrics
      - nodes/spec
      - nodes/stats
      - nodes/proxy
      - nodes/pods
      - nodes/healthz
    verbs:
      - get
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: datadog-agent
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: datadog-agent
subjects:
  - kind: ServiceAccount
    name: datadog-agent
    namespace: <FARGATE_NAMESPACE>
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: datadog-agent
  namespace: <FARGATE_NAMESPACE>
```
2. Deploy it with
```
kubectl apply -f rbac.yaml -n <FARGATE_NAMESPACE>
```
3. Create a Kubernetes secret containing your Datadog API key and Cluster Agent token in the Datadog installation and application namespaces
```
kubectl create secret generic datadog-secret -n datadog \
        --from-literal api-key=<YOUR_DATADOG_API_KEY> --from-literal token=<CLUSTER_AGENT_TOKEN>
kubectl create secret generic datadog-secret -n <FARGATE_NAMESPACE> \
        --from-literal api-key=<YOUR_DATADOG_API_KEY> --from-literal token=<CLUSTER_AGENT_TOKEN>
```
<CLUSTER_AGENT_TOKEN> must be 32 alphanumeric characters, for example, "abcdabcdabcdabcdabcdabcdabcdabcd"

4. Install the Datadog Operator
```
helm repo add datadog https://helm.datadoghq.com
helm install datadog-operator datadog/datadog-operator
```

5. Configure datadog-agent.yaml
```
apiVersion: datadoghq.com/v2alpha1
kind: DatadogAgent
metadata:
  name: datadog-operator
spec:
  global:
    # Required in case the Agent cannot resolve the cluster name through IMDS. See the note below.
    clusterName: <EKS_CLUSTER_NAME> 
    clusterAgentToken: <CLUSTER_AGENT_TOKEN> 
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
      hostPortConfig:
        enabled: true
      unixDomainSocketConfig:
        enabled: false
    dogstatsd:
      hostPortConfig:
        enabled: true
      unixDomainSocketConfig:
        enabled: false
    npm:
      enabled: true
    admissionController:
      enabled: true
      agentCommunicationMode: hostip
      agentSidecarInjection:
        enabled: true
        provider: fargate
        selectors:
        - objectSelector:
            matchLabels:
              "app": my-fargate-nginx # this label for custom selector to target workload pods instead of updating the workload to add agent.datadoghq.com/sidecar:fargate labels.
        profiles:
        - env:
          - name: DD_PROCESS_AGENT_PROCESS_COLLECTION_ENABLED
            value: "true"

  override:
    nodeAgent:
      env:
        - name: DD_LOGS_CONFIG_AUTO_MULTI_LINE_DETECTION
          value: "true"
```
> [!NOTE]Things to be aware of
>  - If you set the name of your DatadogAgent to be datadog and the Fargate RBAC resources in your rbac.yaml to be datadog-agent, the Operator will override the Fargate RBAC resources with its own RBAC called <datadogagent_name>-agent due to conflicting names hence causing issues. You want to either change the name of the DatadogAgent to be something else other than datadog OR change the RBAC resource names. In the operator.yaml I named the DatadogAgent datadog-operator so the RBAC doesn’t conflict.
>  - Use the spec.features.admissionController.agentSidecarInjection.selectors property to configure a custom selector to target workload pods instead of updating the workload to add agent.datadoghq.com/sidecar:fargate labels.

6. Deploy Agent with the above configuration file
```
kubectl apply -f datadog-agent.yaml -n datadog
```
The Admission Controller does not mutate pods that are already created. In case you have deployed agent after current workload have been running, you need to scale or rolling update the workloads.

7. In application deployment manifest, you need to add a service account
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: fargate-nginx-deployment
  namespace: crc-fargate-apps
spec:
  replicas: 2
  selector:
    matchLabels:
      app: my-fargate-nginx
  template:
    metadata:
      labels:
        app: my-fargate-nginx
      #  agent.datadoghq.com/sidecar: fargate <-- required when not define spec.features.admissionController.agentSidecarInjection.selectors
    spec:
      serviceAccountName: datadog-agent # service account with required RBAC
      containers:
      - name: my-fargate-nginx
        image: nginx
        ports:
        - containerPort: 80

```
** Bind this [RBAC (Getting Started with Datadog )](https://docs.datadoghq.com/integrations/eks_fargate/?tab=datadogoperator#aws-eks-fargate-rbac) to application pod by setting Service Account name. In this case, I created a new one called datadog-agent as you see in the first step. If application deployment already has a service account in place, you can use that service account instead of creating a datadog-agent one by adding the following permission to your existing ClusterRole
