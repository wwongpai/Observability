# How to install the Datadog Agent on GKE Cluster Standard (You manage the cluster’s underlying infrastructure, giving you node configuration flexibilit)

Official document
--------
[https://docs.datadoghq.com/containers/kubernetes/installation?tab=helm](https://docs.datadoghq.com/containers/kubernetes/distributions/?tab=datadogoperator#standard)


Quick set up
--------
Add the Helm Datadog repo:
```
$ helm repo add datadog https://helm.datadoghq.com
$ helm repo update
```
Create API key and App key in Datadog following [this link](https://docs.datadoghq.com/account_management/api-app-keys)

Create secret
```
kubectl create secret generic datadog-api-secret --from-literal api-key=$DD_API_KEY
```
```
kubectl create secret generic datadog-app-secret --from-literal app-key=$DD_APP_KEY
```

Create values.yaml, please refer the following link or using this [example values.yaml](https://github.com/wwongpai/Observability/blob/main/agent/eks/value.yaml)
```
helm charts value - https://github.com/DataDog/helm-charts/blob/main/charts/datadog/values.yaml
basic value example - https://github.com/DataDog/helm-charts/blob/main/examples/datadog/agent_basic_values.yaml
specfic to EKS - https://docs.datadoghq.com/containers/kubernetes/distributions/?tab=helm#EKS
```

Start the agent with this command:
```
$ helm install datadog -f values.yaml datadog/datadog
```

Start using and enjoy the ride:
- [Unify kubernetes insight with the Kubernetes overveiw page](https://www.datadoghq.com/blog/unify-kubernetes-insights-with-the-kubernetes-overview-page)
- [Live container](https://docs.datadoghq.com/infrastructure/livecontainers)

Need more visibilities? please jump into the area you are looking for:
- [Integrations & Discovery](https://docs.datadoghq.com/containers/kubernetes/integrations/?tab=kubernetesadv1)
- [APM](https://docs.datadoghq.com/containers/kubernetes/apm/?tab=helm)
- [Log Collection](https://docs.datadoghq.com/containers/kubernetes/log/?tab=helm)
- [Prometheus & OpoenMetrics: Collect your exposed Prometheus and OpenMetrics metrics from your application running inside Kubernetes by using the Datadog Agent, and the Datadog-OpenMetrics or Datadog-Prometheus integrations](https://docs.datadoghq.com/containers/kubernetes/prometheus/?tab=kubernetesadv2)
- [Control plane monitoring](https://docs.datadoghq.com/containers/kubernetes/control_plane/?tab=helm)

Appendix 1 - Helm comnmands
--------
Install a chart:
```
helm install <RELEASE_NAME> -f values.yaml --set datadog.apiKey=<DATADOG_API_KEY> datadog/datadog --set targetSystem=<TARGET_SYSTEM>
```
Upgrade a chart (for changes after install):
```
helm upgrade <RELEASE_NAME> -f values.yaml --set datadog.apiKey=<DATADOG_API_KEY> datadog/datadog --set targetSystem=<TARGET_SYSTEM>
```
Uninstall a chart:
```
helm uninstall <RELEASE_NAME>
```
List all your charts:
```
helm list
```
Get values
```
helm get values <RELEASE_NAME>
```
You can append the -n NAMESPACE flag on any of these commands too, to run these with respect to a given Kubernetes Namespace.

You can find more commands about Helm here: https://helm.sh/docs/helm/helm/




