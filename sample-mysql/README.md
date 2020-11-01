# Using this sample

[blaqkube/blaqkode](https://github.com/blaqkube/blaqkode) contains an example
in the `mysql-sample` directory that assumes already installed the Prometheus
operator and created a `Prometheus` Kind in the `prometheus` namespace. To
ensure, it can manage the `PodMonitor` you create in the `mysql` namespace
2 conditions must be met:

- The `mysql` namespace must have a label that is `monitoring=enabled`. To
  enable it, you an simply run the command below:

```shell
kubectl label ns/mysql monitoring=enabled
kubectl get ns/mysql --show-labels
```

- You must patch the `prometheus` Prometheus resource with the command
  below

```shell
kubectl patch prometheus/prometheus -n prometheus \
--type merge --patch "spec:
  podMonitorSelector:
    matchLabels:
      team: database
  podMonitorNamespaceSelector:
    matchLabels:
      monitoring: enabled
"
```


