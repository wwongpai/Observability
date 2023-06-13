# How to Instrumenting Node.js Serverless Applications using Serverless Framework and Datadog plugin

1. For this example, I will use a simple serverless node application from [Serverless Framework’s repository](https://github.com/serverless/examples).

2. Install node, npm, and the Serverless Framework CLI
[Installing Node.js® and NPM on Mac](https://treehouse.github.io/installation-guides/mac/node-mac.html)
[Setting Up Serverless Framework With AWS](https://www.serverless.com/framework/docs/getting-started/)
[AWS Credentials](https://www.serverless.com/framework/docs/providers/aws/guide/credentials/)

3. Deploy an example serverless app
Using a simple node app, [more detail](https://github.com/serverless/examples/tree/v3/aws-node-express-api)

4. The Datadog Serverless Plugin automatically configures your functions to send metrics, traces, and logs to Datadog through the Datadog Lambda Extension ([link](https://docs.datadoghq.com/serverless/installation/nodejs/?tab=serverlessframework) to an official document).

Install the Datadog Serverless Plugin:
```
$ npm install serverless-plugin-datadog --save
```

Update your serverless.yml:
```
custom:
  datadog:
    site: <DATADOG_SITE>
    apiKey: <DATADOG_API_KEY_SECRET_ARN>
```

5. Check your package.json and node module
$ vi package.json
You should see "serverless-plugin-datadog" as a dependency

$ ls -l node_modules | grep serverless-plugin-datadog

Refer to this [link](https://docs.datadoghq.com/serverless/libraries_integrations/plugin/) for configuration option of serverless-plugin-datadog
