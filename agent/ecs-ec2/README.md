# Datadog Agent with ECS on Fargate

The Datadog Agent in ECS should be deployed as a container once on every EC2 instance in your ECS cluster. This is done by creating a Task Definition for the Datadog Agent container and deploying it as a Daemon service. Each Datadog Agent container then monitors the other containers on their respective EC2 instances.

[Refer link](https://docs.datadoghq.com/containers/amazon_ecs/?tab=awscli)https://docs.datadoghq.com/containers/amazon_ecs/?tab=awscli
