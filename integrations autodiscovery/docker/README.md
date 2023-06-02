# How to configure integrations Autodiscovery with Docker

Official document
--------
https://docs.datadoghq.com/containers/docker/integrations/?tab=dockeradv1
Autodiscovery (AD) is a feature that allows the Datadog agent to run checks against containers that spawn dynamically in a containerized infrastructure.

Example
--------
Docker labels should be placed on the application container. The agent will search through each containers' labels to see if there are any AD templates. Examples for dockerfile, docker compose, docker run, and docker swarm labels can be found [here](https://docs.datadoghq.com/containers/docker/integrations/?tab=dockeradv1#configuration). Examples for ECS and Fargate can be found [here](https://docs.datadoghq.com/integrations/faq/integration-setup-ecs-fargate/?tab=rediswebui#examples).

```
labels:
  com.datadoghq.ad.check_names: '["nginx"]'
  com.datadoghq.ad.init_configs: '[{}]'
  com.datadoghq.ad.instances: '[{"nginx_status_url": "http://%%host%%/nginx_status"}]'
```
👋 Notice:
- The integration name in the check_names label should exactly match the agent check's integration name. Double check the naming with [the integrations-core repo](https://github.com/DataDog/integrations-core/tree/master).
- init_configs is needed even if empty. If this isn't present, the AD configuration will not be applied for this check, and the check won't run.
- Instances should be in JSON format. Use a YAML to JSON (or vice versa) formatter to double check that the syntax is correct. Each integration has an example yaml file to compare against (integrations, core checks).
- In ECS and Fargate, quotations should NOT be escaped when adding labels via the Web UI as AWS will automatically insert escape characters into the JSON. When adding labels via the AWS CLI or directly in the task definition's JSON, quotations NEED to be escaped. For examples, see our public docs.

