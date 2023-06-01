# How to install the Datadog Agent on AWS EKS Cluster

Official document
--------
[Installing the necessary components\](https://docs.datadoghq.com/containers/kubernetes/installation?tab=helm)


Quick set up
--------
add the Helm Datadog repo:
```
$ helm repo add datadog https://helm.datadoghq.com
$ helm repo update
```
Create values.yaml, please refer the following link or using values.yaml in this repo
```
helm charts value - https://github.com/DataDog/helm-charts/blob/main/charts/datadog/values.yaml
basic value example - https://github.com/DataDog/helm-charts/blob/main/examples/datadog/agent_basic_values.yaml
specfic to EKS - https://docs.datadoghq.com/containers/kubernetes/distributions/?tab=helm#EKS
```

Start the agent with this command:
```
$ helm install datadog -f values.yaml datadog/datadog
```



