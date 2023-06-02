# How to install the Datadog agent on Host

Datadog Agent Details
--------
For details about installing the agent, see [Getting started with the Agent](https://docs.datadoghq.com/getting_started/agent/).

The Datadog Agent is open source, installed onto your hosts, and is push-only.

The Datadog Agent comes with many integrations for data collection across a wide variety of sources, see [Integrations](https://docs.datadoghq.com/integrations/).

- [Supported OS versions](https://docs.datadoghq.com/agent/basic_agent_usage/?tab=agentv6v7#supported-platforms)
- [Agent Commands](https://docs.datadoghq.com/agent/guide/agent-commands/?tab=agentv6v7)
- [Agent Configuration Files](https://docs.datadoghq.com/agent/guide/agent-configuration-files/?tab=agentv6v7)
- [Adding the agent package to an internal repository](https://docs.datadoghq.com/agent/guide/installing-the-agent-on-a-server-with-limited-internet-connectivity/#pagetitle)


Firewall and Proxy Considerations
--------
Datadog will never reach directly to the agent for any information. Traffic is always initiated by the
Datadog Agent to our service. The Datadog Agent will only send outbound traffic and will send all
traffic over SSL via 443 TCP.

For more information about agent endpoints and whitelisting, see: [Network Traffic](https://docs.datadoghq.com/agent/guide/network/?tab=agentv6v7).
For a list of ports required for local host communication, see: [Open Ports](https://docs.datadoghq.com/agent/guide/network/?tab=agentv6v7#open-ports).

If your network restricts outbound traffic, you can route all agent traffic by proxy. For detailed information about setting up proxies, see: [Agent Proxy Configuration](https://docs.datadoghq.com/agent/proxy/?tab=agentv6v7).


Configure Agent Checks
--------
The Datadog Agent runs a core set of checks by default. These include cpu, disk, memory, uptime, and more. When you run the datadog-agent status command, these appear in the Collector > Running Checks section of the output.

Each of these checks has a corresponding configuration file located in a subdirectory of datadog-agent/conf.d/.

To see a concise list of the checks the Agent is running and their configuration file paths, run the following command in the terminal.
```
datadog-agent configcheck
```
