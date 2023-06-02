# How to configure integrations Autodiscovery with Docker

Official document
--------
https://docs.datadoghq.com/containers/docker/integrations/?tab=dockeradv1
Autodiscovery (AD) is a feature that allows the Datadog agent to run checks against containers that spawn dynamically in a containerized infrastructure.

Example 1
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

Example 2
--------
This example shows step to configure the integrations discovery for postgres db running as part of compose application.

example [docker-compose.yaml](https://github.com/wwongpai/Observability/blob/main/integrations%20autodiscovery/docker/docker-compose-postgres-example.yaml)

Adding Autodiscovery labels to the labels block of the db service. These Autodiscovery labels tell the Agent to run the postgres check on this container, and provide the credentials for querying metrics.
```
com.datadoghq.ad.check_names: '["postgres"]'
com.datadoghq.ad.init_configs: '[{}]'
com.datadoghq.ad.instances: '[{"host":"%%host%%", "port":5432,"username":"xxx","password":"xxx"}]'
```

To tells Datadog to use the PostgreSQL integration's log pipeline to parse this service's logs more intelligently, and to tag the log lines with service:database.
```
com.datadoghq.ad.logs: '[{"source": "postgresql", "service": "database"}]'
```

Finally, you need to add an environment variable to the agent service. This allows the Agent to accept APM traces from other containers. You'll see later how APM traces PostgreSQL through instrumented applications that connect to the database.
```
- DD_APM_NON_LOCAL_TRAFFIC=true
```

[Full docker-compose.yaml](https://github.com/wwongpai/Observability/blob/main/integrations%20autodiscovery/docker/docker-compose-postgres-example-ad.yaml) after adding above autodiscovery labels

Start application:
```
docker-compose up -d
```

Check the Datadog Agent status by running the following command:
```
docker-compose exec datadog agent status
```
