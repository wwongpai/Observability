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

example docker-compose.yaml
```
version: '3'
services:
  datadog:
    image: 'datadog/agent:7.31.1'
    environment:
      - DD_API_KEY
      - DD_HOSTNAME=dd101-sre-host
      - DD_LOGS_ENABLED=true
      - DD_LOGS_CONFIG_CONTAINER_COLLECT_ALL=true
      - DD_PROCESS_AGENT_ENABLED=true
      - DD_DOCKER_LABELS_AS_TAGS={"my.custom.label.team":"team"}
      - DD_TAGS='env:dd101-sre'
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro
      - /proc/:/host/proc/:ro
      - /sys/fs/cgroup/:/host/sys/fs/cgroup:ro
  discounts:
    environment:
      - FLASK_APP=discounts.py
      - FLASK_DEBUG=1
      - POSTGRES_PASSWORD
      - POSTGRES_USER
      - POSTGRES_HOST=db
    image: 'public.ecr.aws/x2b9z2t7/ddtraining/discounts-fixed:2.2.0'
    ports:
      - '5001:5001'
    depends_on:
      - datadog
      - db
    labels:
      com.datadoghq.tags.env: 'dd101-sre'
      com.datadoghq.tags.service: 'discounts-service'
      com.datadoghq.tags.version: '2.2.0'
      my.custom.label.team: 'discounts'
  frontend:
    image: 'public.ecr.aws/x2b9z2t7/ddtraining/storefront-fixed:2.2.0'
    ports:
      - '3000:3000'
    depends_on:
      - datadog
      - discounts
      - advertisements
    labels:
      com.datadoghq.tags.env: 'dd101-sre'
      com.datadoghq.tags.service: 'store-frontend'
      com.datadoghq.tags.version: '2.2.0'
      my.custom.label.team: 'frontend'
  advertisements:
    environment:
      - FLASK_APP=ads.py
      - FLASK_DEBUG=1
      - POSTGRES_PASSWORD
      - POSTGRES_USER
      - POSTGRES_HOST=db
    image: 'public.ecr.aws/x2b9z2t7/ddtraining/advertisements-fixed:2.2.0'
    ports:
      - '5002:5002'
    depends_on:
      - datadog
      - db
    labels:
      com.datadoghq.tags.env: 'dd101-sre'
      com.datadoghq.tags.service: 'advertisements-service'
      com.datadoghq.tags.version: '2.2.0'
      my.custom.label.team: 'advertisements'
  db:
    image: postgres:11-alpine
    restart: always
    environment:
      - POSTGRES_PASSWORD
      - POSTGRES_USER
    ports:
      - '5432:5432'
    labels:
      com.datadoghq.tags.env: 'dd101-sre'
      com.datadoghq.tags.service: 'database'
      com.datadoghq.tags.version: '11'
      my.custom.label.team: 'database'
    volumes:
      - /root/postgres:/var/lib/postgresql/data
      - /root/dd_agent.sql:/docker-entrypoint-initdb.d/dd_agent.sql
  puppeteer:
    image: buildkite/puppeteer:10.0.0
    volumes:
      - /root/puppeteer-mobile.js:/puppeteer.js
      - /root/puppeteer.sh:/puppeteer.sh
    environment:
      - STOREDOG_URL
      - PUPPETEER_TIMEOUT
    depends_on:
      - frontend
    command: bash puppeteer.sh
```

Adding Autodiscovery labels to the labels block of the db service,
```
com.datadoghq.ad.check_names: '["postgres"]'
com.datadoghq.ad.init_configs: '[{}]'
com.datadoghq.ad.instances: '[{"host":"%%host%%", "port":5432,"username":"xxx","password":"xxx"}]'
```

Full docker-compose.yaml after adding above autodiscovery labels
```
version: '3'
services:
  datadog:
    image: 'datadog/agent:7.31.1'
    environment:
      - DD_API_KEY
      - DD_HOSTNAME=dd101-sre-host
      - DD_LOGS_ENABLED=true
      - DD_LOGS_CONFIG_CONTAINER_COLLECT_ALL=true
      - DD_PROCESS_AGENT_ENABLED=true
      - DD_DOCKER_LABELS_AS_TAGS={"my.custom.label.team":"team"}
      - DD_TAGS='env:dd101-sre'
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro
      - /proc/:/host/proc/:ro
      - /sys/fs/cgroup/:/host/sys/fs/cgroup:ro
  discounts:
    environment:
      - FLASK_APP=discounts.py
      - FLASK_DEBUG=1
      - POSTGRES_PASSWORD
      - POSTGRES_USER
      - POSTGRES_HOST=db
    image: 'public.ecr.aws/x2b9z2t7/ddtraining/discounts-fixed:2.2.0'
    ports:
      - '5001:5001'
    depends_on:
      - datadog
      - db
    labels:
      com.datadoghq.tags.env: 'dd101-sre'
      com.datadoghq.tags.service: 'discounts-service'
      com.datadoghq.tags.version: '2.2.0'
      my.custom.label.team: 'discounts'
  frontend:
    image: 'public.ecr.aws/x2b9z2t7/ddtraining/storefront-fixed:2.2.0'
    ports:
      - '3000:3000'
    depends_on:
      - datadog
      - discounts
      - advertisements
    labels:
      com.datadoghq.tags.env: 'dd101-sre'
      com.datadoghq.tags.service: 'store-frontend'
      com.datadoghq.tags.version: '2.2.0'
      my.custom.label.team: 'frontend'
  advertisements:
    environment:
      - FLASK_APP=ads.py
      - FLASK_DEBUG=1
      - POSTGRES_PASSWORD
      - POSTGRES_USER
      - POSTGRES_HOST=db
    image: 'public.ecr.aws/x2b9z2t7/ddtraining/advertisements-fixed:2.2.0'
    ports:
      - '5002:5002'
    depends_on:
      - datadog
      - db
    labels:
      com.datadoghq.tags.env: 'dd101-sre'
      com.datadoghq.tags.service: 'advertisements-service'
      com.datadoghq.tags.version: '2.2.0'
      my.custom.label.team: 'advertisements'
  db:
    image: postgres:11-alpine
    restart: always
    environment:
      - POSTGRES_PASSWORD
      - POSTGRES_USER
    ports:
      - '5432:5432'
    labels:
      com.datadoghq.tags.env: 'dd101-sre'
      com.datadoghq.tags.service: 'database'
      com.datadoghq.tags.version: '11'
      my.custom.label.team: 'database'
      com.datadoghq.ad.check_names: '["postgres"]'
      com.datadoghq.ad.init_configs: '[{}]'
      com.datadoghq.ad.instances: '[{"host":"%%host%%", "port":5432,"username":"xxx","password":"xxx"}]'
    volumes:
      - /root/postgres:/var/lib/postgresql/data
      - /root/dd_agent.sql:/docker-entrypoint-initdb.d/dd_agent.sql
  puppeteer:
    image: buildkite/puppeteer:10.0.0
    volumes:
      - /root/puppeteer-mobile.js:/puppeteer.js
      - /root/puppeteer.sh:/puppeteer.sh
    environment:
      - STOREDOG_URL
      - PUPPETEER_TIMEOUT
    depends_on:
      - frontend
    command: bash puppeteer.sh

```
