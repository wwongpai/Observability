# How to configure an Integration on Host

Configure Agent Checks
--------
The Datadog Agent runs a core set of checks by default. These include cpu, disk, memory, uptime, and more. When you run the datadog-agent status command, these appear in the Collector > Running Checks section of the output.

Each of these checks has a corresponding configuration file located in a subdirectory of datadog-agent/conf.d/.

To see a concise list of the checks the Agent is running and their configuration file paths, run the following command in the terminal.
```
datadog-agent configcheck
```
Notice that the configuration files for these checks all end in .default, for example, datadog-agent/conf.d/cpu.d/conf.yaml.default


Example configure an integration for Postgres
--------
- [Metric collection](https://docs.datadoghq.com/integrations/postgres/?tab=host#metric-collection)
- [Log collection](https://docs.datadoghq.com/integrations/postgres/?tab=host#metric-collection)
- [Instrument your application that makes requests to Postgres](https://docs.datadoghq.com/integrations/postgres/?tab=host#trace-collection)
