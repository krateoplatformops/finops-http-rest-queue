# FinOps HTTP REST Queue
This repository in an HTTP API that listens on port 8080 for requests on the path /upload. The data received is supposed to be a JSON file in the following format:
```
{
  "resourceId": "/subscriptions/<....>/resourcegroups/finops/providers/microsoft.compute/virtualmachines/finops-resource-usage-test",
  "optimization": {
    "resourceName": "Percentage CPU",
    "resourceDelta": -69,
    "typeChange": {
      "cyclic": "day",
      "from": "19:00",
      "to": "7:00"
    }
  }
}
```
The data uploaded is then published on a nats service installed in the same Kubernetes cluster.
The URL should also include the query parameter `topic` to identify the topic.

## How to install
The installation can be performed using HELM:
```sh
$ helm repo add krateo https://charts.krateo.io
$ helm repo update krateo
$ helm install finops-http-rest-queue krateo/finops-http-rest-queue
```