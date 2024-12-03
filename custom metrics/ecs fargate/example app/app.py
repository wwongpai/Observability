from datadog import initialize, statsd
import time

options = {
    'statsd_host': '127.0.0.1',
    'statsd_port': 8125
}

initialize(**options)

i = 0

while True:
    i += 1
    statsd.gauge('example_metric.gauge', i, tags=["environment:dev"])
    print(f"Sent metric: example_metric.gauge with value {i}")
    time.sleep(10)
