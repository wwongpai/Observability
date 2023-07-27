EKS Fargate - Infra monitoring
Reference [Official Document Guide](https://docs.datadoghq.com/integrations/eks_fargate/#configuration)

Monitor EKS Fargate
--------
Infrastucture monitoring - 
Log Collection
Process Collection
Autodiscovery Integrations

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
