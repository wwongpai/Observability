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
Several important things to note:

- You can notice that init_configs and instances are actually the different sections we use to configure each integration on a regular host (non-containerized environment), [here](https://github.com/DataDog/integrations-core/blob/master/redisdb/datadog_checks/redisdb/data/conf.yaml.example) for example for Redis. This means that you can use all the parameters available in the conf.yaml.example file in these annotations. 

- The check_names is just the name of the check (here redisdb). To be sure about which name you can use, you can go to [the integrations-core repository](https://github.com/DataDog/integrations-core)

- The container_identifier that you can see right before check_names, init_configs and instances which is redis in this example. This value must match the value in spec.containers.name

Now that our Redis pod is ready we can deploy it with the following command:
```
$ kubectl apply -f redis.yaml
```
After a few minutes, you can run the Agent status command on the Agent pod to see if this redis check works:
```
$ kubectl exec -it <AGENT_POD_NAME> agent status
```
And if it works, you should see this output:
![example-redis-from-agent-status](https://p-qkfgo2.t2.n0.cdn.getcloudapp.com/items/eDuEzez6/0941781c-276c-4b79-95c5-8f4482c47ee5.jpg?v=bb7c808becfc1d49c757ff79089988ea)

ConfigMap via Helm
--------
Let’s now configure the Redis check the other way around: the configuration will be on the Agent side and not on the application pod side.

For this, you can keep [the Redis pod](https://github.com/wwongpai/Observability/blob/main/integrations%20autodiscovery/kubernetes/redis-no-annotation.yaml) without the annotations:
```
apiVersion: v1
kind: Pod
metadata:
  name: redis    
  labels:
    name: redis
spec:
  containers:
    - name: redis
      image: redis
      ports:
        - containerPort: 6379
```
This method requires a file similar to that redisdb.d/auto_conf.yaml file that the Agent has by default. 

To create this file in Kubernetes you can create a ConfigMap containing the file contents, volume mapping that ConfigMap into the pod, and volumeMount mapping that into the agent container. That being said, Helm can do this all the legwork of this for you through [the datadog.confd section](https://github.com/DataDog/helm-charts/blob/main/charts/datadog/values.yaml#L499).
```
datadog: 
  #(...)
  confd:
    redisdb.yaml: |-
      ad_identifiers:
        - redis
      init_config:
      instances:
        - host: "%%host%%"
          port: "6379"
```
