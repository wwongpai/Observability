# How to configure integrations Autodiscovery with Kubernetes.

Official document
--------
https://docs.datadoghq.com/containers/kubernetes/integrations/?tab=kubernetesadv1


Quick set up
--------
In a non-containerized deployment you would go to the Agent’s integration configurations in the /etc/datadog-agent/conf.d folder and modify the configs to generally access a “local” instance of that integration running on the same host. In a containerized world this happens a little differently as

The applications the Agent will monitor may vary depending on where Kubernetes schedules the other application pods

The endpoints (IP addresses) the Agent uses need to evaluate to the dynamic IP address of the pod/container it is supposed to monitor

To do this we have our Autodiscovery components the Agent can use to automatically setup these integration configs pointed at the right endpoints.
```
$ helm repo add datadog https://helm.datadoghq.com
$ helm repo update
```
