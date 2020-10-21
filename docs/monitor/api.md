Monitor API
===

You can use the following two types of interfaces to monitor the Zai cluster state:

- [The state interface](#use-the-state-interface): this interface uses the HTTP interface to get the component basic information.
- [The metrics interface](#use-the-metrics-interface): this interface uses Prometheus to record the detailed information of the various operations in components and views these metrics using Grafana.

## Use the state interface

The state interface monitors the basic information of a specific component in the Zai cluster. It can also act as the monitor interface for Keepalive messages. In addition, the state interface for Keeper can get the details of the entire Zai cluster.

### Zai 

Zai will be embedded in other application, the basic information will only include version in present.

- Zai API address: `http://${host}:${port}`
- Default port: `7000`

The following example uses `http://${host}:${port}/status` to get the current state of the Zai server and to determine whether the server is alive. The result is returned in JSON format.

```bash
curl http://127.0.0.1:7000/status
{
    version: {
        version: "1.0.0",  # The Zai version number.
        git_hash: "26fb1766bda5b90f61d57a0170f5886e4cb17ece"  # The Git Hash of the current Zai code.
        git_branch: "master" # The Git branch of current Zai code.
    }
}
```

### ZMap

ZMap will be embedded in keeper, the basic information will only include version in present.

### Keeper server

- Keeper API address: `http://${host}:${port}/kp/api/v1/${api_name}`
- Default port: `7001`

The Keeper interface provides the state of all the ZBuf servers and the cluster information. See the following example for the information about a single-node ZBuf cluster:

```bash
curl http://127.0.0.1:7001/kp/api/v1/zbuf
{
  "count": 1,  # The number of ZBuf nodes.
  "zbuf": [  # The list of ZBuf nodes.
    # The details about the single ZBuf node.
    {
      "node": {
        "id": 1,
        "address": "127.0.0.1:7002",
        "version": "1.0.0",
        "state": "up"
      },
      "status": {
        "capacity": "4096 GiB",  # The total capacity.
        "available": "1024 GiB",  # The available capacity.
        "extent_count": 512,
        "start_ts": "2020-10-20T15:29:12+08:00",  # The starting timestamp.
        "last_heartbeat_ts": "2020-10-20T15:29:10.961171958+08:00",  # The timestamp of the last heartbeat.
        "uptime": "20.961171958s"
      }
    }
  ]
}
```

## Use the metrics interface

The metrics interface monitors the state and performance of the entire Zai cluster.
