# Distributed voting system on Kubernetes

A distributed system that ingests a high-volume stream of votes, queues it,
persists it to two different stores and visualizes it in real time — deployed
as microservices on Kubernetes.

Built as a load-handling exercise: the point was not the voting itself, but
what happens when thousands of requests arrive at once and nothing is allowed
to be lost.

**Stack:** Go · gRPC · Kafka · Redis · MongoDB · Kubernetes · Grafana · Locust · Node.js · Vue · Cloud Run

---

## How the data flows

```
Locust (load generator)
      │  thousands of simulated votes
      ▼
gRPC client  ──▶  gRPC server          (Go)
                      │
                      ▼
                  Kafka topic           queues and decouples ingestion
                      │
                      ▼
                  Consumer              (Go)
                   ┌──┴──┐
                   ▼     ▼
                 Redis  MongoDB         fast counters / durable records
                   │      │
                   ▼      ▼
                Grafana   Node.js API ──▶ Vue web app   (Cloud Run)
```

**Why Kafka in the middle.** Writing straight from the gRPC server to the
databases means the ingestion rate is capped by the slowest write. The queue
decouples them: the server only has to publish, and the consumer drains at its
own pace without dropping votes under a spike.

**Why two stores.** Redis holds the live counters that the dashboards read
constantly; MongoDB keeps the durable record that the API queries later. Each
one does what it is good at.

---

## What's in here

| Path | What it is |
|---|---|
| `Proyecto2/grpc/` | gRPC client and server in Go, with generated protobuf code |
| `Proyecto2/consumer/` | Go service that reads from Kafka and writes to Redis and MongoDB |
| `Proyecto2/kafka/` | Kafka deployment manifests |
| `Proyecto2/redis/` · `mongo/` | Store deployments and services |
| `Proyecto2/locust/` | Python load generator that simulates the vote traffic |
| `Proyecto2/cloudRun/` | Node.js API and Vue front end deployed to Cloud Run |
| `Proyecto2/ingress/` | ingress-nginx routing |
| `Proyecto2/deploy.yaml` | Full cluster deployment |

Every service ships with its own `Dockerfile` and Kubernetes manifest.

---

## Running it

Requires a Kubernetes cluster with `kubectl` configured.

```bash
kubectl apply -f Proyecto2/namespace.yml
kubectl apply -f Proyecto2/kafka/
kubectl apply -f Proyecto2/redis/ -f Proyecto2/mongo/
kubectl apply -f Proyecto2/deploy.yaml
kubectl apply -f Proyecto2/ingress/

# generate traffic
locust -f Proyecto2/locust/traffic.py
```

Configuration lives in Kubernetes secrets (`Proyecto2/secrets/`) — ports, Kafka
broker address and topic name. No credentials are committed.

---

## Context

Operating Systems 1 course project, Computer Science and Systems Engineering —
Universidad de San Carlos de Guatemala.
