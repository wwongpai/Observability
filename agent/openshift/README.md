# How to install the Datadog Agent on Openshift

Official document
--------
https://docs.datadoghq.com/containers/kubernetes/installation?tab=helm
https://docs.datadoghq.com/containers/kubernetes/distributions/?tab=helm#Openshift

OpenShift comes with hardened security by default (SELinux, SecurityContextConstraints), thus requiring some specific configuration:
- Create SCC for Node Agent and Cluster Agent
- Specific CRI socket path as OpenShift uses CRI-O container runtime
- Kubelet API certificates may not always be signed by cluster CA
- Tolerations are required to schedule the Node Agent on master and infra nodes
- Cluster name should be set as it cannot be retrieved automatically from cloud provider

Quick set up
--------
For Datadog, OpenShift differs from the standard Kubernetes deployment through the permissions allowed to the agent. These permissions need to be managed through OpenShift. That said, we make it a point to explain that the agent's Daemonset will typically need elevated privileges to run all features of Datadog: https://docs.datadoghq.com/integrations/openshift/#configuration. As explained in the doc, these privileges are typically allowed through Security Context Constraints (or SCCs).

Add the Helm Datadog repo:
```
$ helm repo add datadog https://helm.datadoghq.com
$ helm repo update
```
Create API key and App key in Datadog following [this link](https://docs.datadoghq.com/account_management/api-app-keys)

Create secret
```
oc create secret generic datadog-app-key -n datadog --from-literal=app-key=<ADD_YOUR_APP_KEY_HERE>
```
```
oc create secret generic datadog-api-key -n datadog --from-literal=api-key=<ADD_YOUR_API_KEY_HERE>
```

Create values.yaml, please refer the following link or using this [example values.yaml](https://github.com/wwongpai/Observability/blob/main/agent/openshift/values.yaml)
```
helm charts value - https://github.com/DataDog/helm-charts/blob/main/charts/datadog/values.yaml
basic value example - https://github.com/DataDog/helm-charts/blob/main/examples/datadog/agent_basic_values.yaml
specfic to Openshift - https://docs.datadoghq.com/containers/kubernetes/distributions/?tab=helm#Openshift
```

The key parts of this configuration are:
- This creates the Security Context Constraints (SCCs) for both the Agent and Cluster Agent. These SCCs ensure the Agent has the right permissions in this OpenShift cluster that it would ordinarily be blocked. This means you do not need to manually deploy the SCC in this cluster.
- Tolerations are added to ensure the Agent can be deployed on the tainted node
- Kubelet TLS Verification is disabled

For reference you can see the SCC template files here:
- [helm-charts/cluster-agent-scc.yaml at main · DataDog/helm-charts ](https://github.com/DataDog/helm-charts/blob/main/charts/datadog/templates/cluster-agent-scc.yaml)
- [helm-charts/agent-scc.yaml at main · DataDog/helm-charts ](https://github.com/DataDog/helm-charts/blob/main/charts/datadog/templates/agent-scc.yaml)

:wave: When installing the chart, we recommend to install on a non-default namespace. Due to existing SecurityContextConstraints in the default namespace. To do so first create your desired namespace (can name whatever):
```
oc new-project datadog-openshift
```

Then run the install command like normal, adding --namespace <Your Namespace> (or -n for short)
```
$ helm install datadog -f values.yaml -n datadog-openshift datadog/datadog
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




