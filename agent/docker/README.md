# How to install the Datadog agent on Docker, Containerd, Podman and via Docker compose

Docker Containerd Podman
--------
[Official doc](https://docs.datadoghq.com/containers/docker/?tab=standard)

```
$ docker run -d --cgroupns host --pid host --name dd-agent -v /var/run/docker.sock:/var/run/docker.sock:ro -v /proc/:/host/proc/:ro -v /sys/fs/cgroup/:/host/sys/fs/cgroup:ro -e DD_API_KEY=<DATADOG_API_KEY> gcr.io/datadoghq/agent:7
```
