# Example of instrumenting Go app 
using net/http to create web app and using Datadog Go tracer to instrument this library. 
Auto-instrument Go app means replacing go library with datadog version, please find additional information here - https://pkg.go.dev/gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http

**Step for testing **
1. Install datadog agent on your host and enable APM agent
2. run
```
go run main.go
```
3. Go to a web browser - http://localhost:8080
4. See the result on Datadog UI
