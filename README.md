# nats streaming chart

## Setup for database

 1. Copy the example yaml to a non-example yaml e.g. `cp values.example.mysql.yml values.mysql.yml`
 2. Setup the connection string in the copied yaml to your needs
 3. Make sure, that the database is seeded with the official seeding query ([MySQL](https://github.com/nats-io/nats-streaming-server/blob/master/mysql.db.sql) / [Postgres](https://github.com/nats-io/nats-streaming-server/blob/master/postgres.db.sql))

## Install the helm chart

```bash
$ helm install --name nats --namespace nats-namespace ./chart -f values.mysql.yml
```

## Test it

Best way to test it locally, port forward a pod:

```bash
$ export POD_NAME=$(kubectl get pods --namespace nats-namespace -l "app=nats-streaming,release=release-name" -o jsonpath="{.items[0].metadata.name}")
$ kubectl port-forward $POD_NAME 4222:4222
```

You can then start consumer and producer:

(consumer)
```bash
$ cd testconsumer
$ go mod tidy
$ go run main.go
```

(producer)
```bash
$ cd tesproducer
$ go mod tidy
$ go run main.go
Enter a message: 
```