# How to Instrumenting Node.js Serverless Applications using Serverless Framework and Datadog plugin

Steps
--------
1. For this example, I will use a simple serverless node application from [Serverless Framework’s repository](https://github.com/serverless/examples).


2. Install node, npm, and the Serverless Framework CLI
- [Installing Node.js® and NPM on Mac](https://treehouse.github.io/installation-guides/mac/node-mac.html)
- [Setting Up Serverless Framework With AWS](https://www.serverless.com/framework/docs/getting-started/)
- [AWS Credentials](https://www.serverless.com/framework/docs/providers/aws/guide/credentials/)


3. Deploy an example serverless app
Using a simple node app, [more detail](https://github.com/serverless/examples/tree/v3/aws-node-express-api)


4. The Datadog Serverless Plugin automatically configures your functions to send metrics, traces, and logs to Datadog through the Datadog Lambda Extension ([link](https://docs.datadoghq.com/serverless/installation/nodejs/?tab=serverlessframework) to an official document).

Install the Datadog Serverless Plugin:
```
$ npm install serverless-plugin-datadog --save
```

Update your [serverless.yml](https://github.com/wwongpai/Observability/blob/main/apm/serverless/nodejs/serverless.yml):
```
custom:
  datadog:
    site: <DATADOG_SITE>
    apiKey: <DATADOG_API_KEY>
```

5. Check your package.json and node module
```
$ vi package.json
```

You should see "serverless-plugin-datadog" as a dependency
```
$ ls -l node_modules | grep serverless-plugin-datadog
```

Refer to this [link](https://docs.datadoghq.com/serverless/libraries_integrations/plugin/) for configuration option of serverless-plugin-datadog


6. Add custom logging inside the handler.js file. For example, you can add:
```
console.log("A log message under root.");
...
console.log("A log message under path.");
```

7. Deploy your serverless application
```
$ sls deploy --verbose
```

8. Remove application
```
$ sls remove
```

Outcomes
--------
Enhanced metric:

![enhanced-metrics](https://p-qkfgo2.t2.n0.cdn.getcloudapp.com/items/OAulxZlL/e0535850-7a85-41a0-b375-ce6e8dd97011.jpg?source=viewer&v=3790348d58b8629c6ea98fce46bb7bac)


Serverless App:

![serverless](https://p-qkfgo2.t2.n0.cdn.getcloudapp.com/items/z8ubwoz4/a560c2e9-45c2-42da-8068-1773ab05b17c.jpg?source=viewer&v=cd0c9684571387c9f79c900de6ad6ea3)


Traces & Logs are collected and connected each other (Correlations):

![trace-log](https://p-qkfgo2.t2.n0.cdn.getcloudapp.com/items/jkuRyp6G/b0fbc090-230b-4583-9338-b12c16e1c254.jpg?source=viewer&v=4cbaef80560d2d42eda4c3194718a931)
