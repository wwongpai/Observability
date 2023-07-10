# Logs Collection

Method 1: Collecting logs from EKS on Fargate with Fluent Bit [Official Document Guide](https://docs.datadoghq.com/integrations/eks_fargate/#log-collection)
--------

Monitor EKS Fargate logs by using Fluent Bit to route EKS logs to CloudWatch Logs and the Datadog Forwarder to route logs to Datadog.
To configure Fluent Bit to send logs to CloudWatch, create a Kubernetes ConfigMap that specifies CloudWatch Logs as its output. The ConfigMap specifies the log group, region, prefix string, and whether to automatically create the log group.

```
 kind: ConfigMap
 apiVersion: v1
 metadata:
   name: aws-logging
   namespace: aws-observability
 data:
   output.conf: |
     [OUTPUT]
         Name cloudwatch_logs
         Match   *
         region us-east-1
         log_group_name awslogs-https
         log_stream_prefix awslogs-firelens-example
         auto_create_group true
```

Use the Datadog Forwarder to collect logs from Cloudwatch and send them to Datadog.


Method 2: Send AWS EKS Fargate Logs with Kinesis Data Firehose [Official Document Guide](https://docs.datadoghq.com/logs/guide/aws-eks-fargate-logs-with-kinesis-data-firehose/)
--------
