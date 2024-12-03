
# How to send custom metrics with DogStatsD

Official document
--------
https://docs.datadoghq.com/developers/dogstatsd/?tab=hostagent
https://docs.datadoghq.com/metrics/custom_metrics/dogstatsd_metrics_submission/
https://docs.datadoghq.com/containers/amazon_ecs/?tab=awscli#dogstatsd
https://docs.datadoghq.com/integrations/ecs_fargate/?tab=webui

Step
--------
To monitor your ECS Fargate tasks with Datadog, run the Agent as a container in same task definition as your application container. To collect custom metrics with Datadog, each task definition should include a Datadog Agent container with DogStatD enabled in addition to the application containers. Follow these setup steps:

1. Adding the Datadog Agent with DogStatD enabled alongside the application container in the same task definition.
```
{
    "containerDefinitions": [
        {
            "name": "custom-metrics-app",
            "image": "wwongpai/dd-custom-metrics:latest",
            "cpu": 256,
            "memory": 512,
            "portMappings": [],
            "essential": true,
            "environment": [
                {
                    "name": "ENV",
                    "value": "dev"
                }
            ],
            "mountPoints": [],
            "volumesFrom": [],
            "systemControls": []
        },
        {
            "name": "datadog-agent",
            "image": "public.ecr.aws/datadog/agent:latest",
            "cpu": 256,
            "memoryReservation": 512,
            "portMappings": [
                {
                    "name": "datadog-agent-8126-tcp",
                    "containerPort": 8126,
                    "hostPort": 8126,
                    "protocol": "tcp"
                },
                {
                    "containerPort": 8125,
                    "hostPort": 8125,
                    "protocol": "udp"
                }
            ],
            "essential": true,
            "environment": [
                {
                    "name": "DD_API_KEY",
                    "value": "XXXXXXXXXXXXXXX"
                },
                {
                    "name": "DD_CONTAINER_EXCLUDE",
                    "value": "image:^aws-fargate-pause$ image:^public.ecr.aws/datadog/agent$"
                },
                {
                    "name": "DD_PROCESS_AGENT_PROCESS_COLLECTION_ENABLED",
                    "value": "true"
                },
                {
                    "name": "DD_SITE",
                    "value": "datadoghq.com"
                },
                {
                    "name": "DD_DOCKER_LABELS_AS_TAGS",
                    "value": "{\"com.docker.service.name\":\"service_name\"}"
                },
                {
                    "name": "ECS_FARGATE",
                    "value": "true"
                },
                {
                    "name": "DD_APM_ENABLED",
                    "value": "true"
                },
                {
                    "name": "DD_DOGSTATSD_NON_LOCAL_TRAFFIC",
                    "value": "true"
                }
            ],
            "mountPoints": [],
            "volumesFrom": [],
            "systemControls": []
        }
    ],
    "family": "custom-metrics-app",
    "taskRoleArn": "arn:aws:iam::XXXXXX:role/ecsExecTaskRole",
    "executionRoleArn": "arn:aws:iam::XXXXX:role/ecsTaskExecutionRole",
    "networkMode": "awsvpc",
    "volumes": [],
    "placementConstraints": [],
    "requiresCompatibilities": [
        "FARGATE"
    ],
    "cpu": "1024",
    "memory": "3072",
    "pidMode": "task",
    "runtimePlatform": {
        "cpuArchitecture": "X86_64",
        "operatingSystemFamily": "LINUX"
    },
    "tags": [
        {
            "key": "Name",
            "value": "custom-metrics-app"
        }
    ]
}
```

2. In this example, I am using the pre-built image wwongpai/dd-custom-metrics:latest. If you prefer to build your own, you can do so from the app directory.
I use an example from https://docs.datadoghq.com/metrics/custom_metrics/dogstatsd_metrics_submission/#gauge to showcase submitting GAUGE metrics to Datadog
```
pip install datadog
```

```
from datadog import initialize, statsd
import time

options = {
    'statsd_host': '127.0.0.1',
    'statsd_port': 8125
}

initialize(**options)

i = 0

while True:
    i += 1
    statsd.gauge('example_metric.gauge', i, tags=["environment:dev"])
    print(f"Sent metric: example_metric.gauge with value {i}")
    time.sleep(10)
```
COUNT, GAUGE, and SET metric types are familiar to StatsD users. TIMER from StatsD is a sub-set of HISTOGRAM in DogStatsD. Additionally, you can submit HISTOGRAM and DISTRIBUTION metric types using DogStatsD. You can learn more example from https://docs.datadoghq.com/metrics/custom_metrics/dogstatsd_metrics_submission/
Note: Depending on the submission method used, the actual metric type stored within Datadog might differ from the submission metric type. When submitting a RATE metric type through DogStatsD, the metric appears as a GAUGE in-app to ensure relevant comparison across different Agents.
