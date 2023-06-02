# How to install the Datadog agent on Docker, Containerd, Podman and via Docker compose

Docker Containerd Podman
--------
[Official doc](https://docs.datadoghq.com/containers/docker/?tab=standard)

```
$ docker run -d --cgroupns host --pid host --name dd-agent -v /var/run/docker.sock:/var/run/docker.sock:ro -v /proc/:/host/proc/:ro -v /sys/fs/cgroup/:/host/sys/fs/cgroup:ro -e DD_API_KEY=<DATADOG_API_KEY> gcr.io/datadoghq/agent:7
```

Compose and the Datadog agent
--------
[Official doc](https://docs.datadoghq.com/agent/guide/compose-and-the-datadog-agent)

The following is an example of how you can monitor a Redis container using Compose. The file structure is:
```
|- docker-compose.yml
|- datadog
  |- Dockerfile
  |- conf.d
    |-redisdb.yaml
```
The docker-compose.yml file describes how your containers work together and sets some of the configuration details for the containers.
```
version: '3'
services:
  redis:
    image: redis
  datadog:
    build: datadog
    pid: host
    environment:
     - DD_API_KEY=${DD_API_KEY}
     - DD_SITE=datadoghq.com
    volumes:
     - /var/run/docker.sock:/var/run/docker.sock
     - /proc/:/host/proc/:ro
     - /sys/fs/cgroup:/host/sys/fs/cgroup:ro
```
The redisdb.yaml is patterned after the [redisdb.yaml.example](https://github.com/DataDog/integrations-core/blob/master/redisdb/datadog_checks/redisdb/data/conf.yaml.example) file and tells the Datadog Agent to look for Redis on the host named redis (defined in docker-compose.yaml above) and to use the standard Redis port:
```
init_config:

instances:
    - host: redis
      port: 6379
```
The Dockerfile is used to instruct Docker compose to build a Datadog Agent image including the redisdb.yaml file at the right location:
```
FROM gcr.io/datadoghq/agent:latest
ADD conf.d/redisdb.yaml /etc/datadog-agent/conf.d/redisdb.yaml
```

To start the application stack, run:
```
docker-compose up -d
```

Once the containers are running, run the Agent status command:
```
docker-compose exec datadog agent status
```

To see the Agent’s configuration, run the Agent config command:
```
docker-compose exec datadog agent config
```

