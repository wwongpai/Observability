CREATE A DATADOG ACCOUNT
Your team can create an account at https://www.datadoghq.com/. Click the “FREE TRIAL” link.

GETTING ACCESS TO DATADOG
Once your team has an account, an existing member can invite others from this page:
https://app.datadoghq.com/account/team.

There are three user types––admin, standard, and read-only. After users have access to a Datadog
account, they can see all dashboards, graphs, and alerts that are included in the account, but only
admins and standard users can make modifications.

For more information on user roles and permissions, see User Management.

DATADOG AGENT DETAILS
For details about installing the agent, see Getting started with the Agent.

The Datadog Agent is open source, installed onto your hosts, and is push-only.

The Datadog Agent comes with many integrations for data collection across a wide
variety of sources, see Integrations.

Run the Datadog Agent in your Kubernetes cluster in order to start collecting your cluster and applications metrics, traces, and logs. See, Kubernetes.

FIREWALL AND PROXY CONSIDERATIONS
Datadog will never reach directly to the agent for any information. Traffic is always initiated by the
Datadog Agent to our service. The Datadog Agent will only send outbound traffic and will send all
traffic over SSL via 443 TCP.

For more information about agent endpoints and whitelisting, see: Network Traffic.
For a list of ports required for local host communication, see: Open Ports.

If your network restricts outbound traffic, you can route all agent traffic by proxy. For detailed information about setting up proxies, see: Agent Proxy Configuration.

BASIC AGENT USAGE
Supported OS versions
Network Performance Monitoring supported platforms
Agent Commands
Agent Configuration Files
Adding the agent package to an internal repository
