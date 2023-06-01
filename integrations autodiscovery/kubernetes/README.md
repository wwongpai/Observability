# How to configure integrations Autodiscovery with Kubernetes.

Official document
--------
https://docs.datadoghq.com/containers/kubernetes/integrations/?tab=kubernetesadv1


Quick set up
--------
In a non-containerized deployment you would go to the Agent’s integration configurations in the /etc/datadog-agent/conf.d folder and modify the configs to generally access a “local” instance of that integration running on the same host. In a containerized world this happens a little differently as

- The applications the Agent will monitor may vary depending on where Kubernetes schedules the other application pods

- The endpoints (IP addresses) the Agent uses need to evaluate to the dynamic IP address of the pod/container it is supposed to monitor

To do this we have our Autodiscovery components the Agent can use to automatically setup these integration configs pointed at the right endpoints. Refer to [document](https://docs.datadoghq.com/containers/kubernetes/integrations/?tab=kubernetesadv1) there are several ways to achieve this, for the sake of a quick setup, example will shows 2 different approaches:
- Annotations: the check will be configured on the pod’s side
- ConfigMap: the check will be configured on the Agent’s side

Annotations
--------
To see how annotations work, we only need to have the Agent running on the cluster and another pod where the integration we are looking for is running.

For this example, we are going to configure annotations on a redis pods to collect redis metrics. You can use this [redis.yaml](https://github.com/wwongpai/Observability/blob/main/integrations%20autodiscovery/kubernetes/redis.yaml) file to deploy it:
```
apiVersion: v1
kind: Pod
metadata:
  name: redis
  annotations:
    ad.datadoghq.com/redis.check_names: '["redisdb"]'
    ad.datadoghq.com/redis.init_configs: '[{}]'
    ad.datadoghq.com/redis.instances: |
      [
        {
          "host": "%%host%%",
          "port":"6379"
        }
      ]      
  labels:
    name: redis
spec:
  containers:
    - name: redis
      image: redis
      ports:
        - containerPort: 6379
```
