Keeper API
=====

# ZBuf Node

## Register node

Every ZBuf node will try to add its own information to Keeper at start up.

The new information will replace old one (If anything changed).

Request：

```
PUT /kp/api/v1/zbuf/add/<node_type>/<node_id>/<addr>

{
    // TODO ([]disk_info: health, cap, writable; []extent_info:)
}
```

Response：

```
200 OK
```

## Get node/nodes information

### Get nodes information

Get all nodes information.

Request：

```
GET /kp/api/v1/zbuf

```

Response：

```
200 OK

Content-Type: application/json

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

### Get node information

Get one node information.

Request：

```
GET /kp/api/v1/zbuf/get/<node_id>

```

Response：

```
200 OK

Content-Type: application/json

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
```

## Offline node

Offline ZBuf nodes, the data on it will be migrated async. 

Request：

```
DELETE /kp/api/v1/zbuf/del/<node_id>

```

Response：

```
200 OK
```

# Disk

## Add disk/disks

### Add disks

Add all new disks in ZBuf.

After set up new disks on a ZBuf node, we should notify this node to load these new disks.

Request：

```
PUT /kp/api/v1/zbuf/disk/add/<node_id>
```

Response：

```
200 OK
```

### Add disk

Add a specific disk.

Request：

```
PUT /kp/api/v1/zbuf/disk/add/<node_id>/<disk_id>
```

Response：

```
200 OK
```

## Offline disk

Offline a specific disk.

Request：

```
DELETE /kp/api/v1/zbuf/disk/add/<node_id>/<disk_id>
```

Response：

```
200 OK
```

# Group

## Get

### Get all groups information

Request：

```
GET /kp/api/v1/zbuf/group
```

Response：

```
200 OK

Content-Type: application/json

{
    "groups": [
                { "type": <type>, 
                "group_id": <group_id>, 
                "state": writable,
                "capactiy": 64GiB,
                "writable": 32GiB,
                "extent_count": 2, 
                "extents":
                    ["status":
                    {seq":1,
                    "node_id":<node_id>,
                    "disk_id":<disk_id>,
                    "state": writable
                    },
                    {seq":2,
                     "node_id":<node_id>,
                     "disk_id":<disk_id>,
                     "state": writable
                    }]
                }
    ]
}

```

### Get specific group information

Request：

```
GET /kp/api/v1/zbuf/group/<group_id>
```

Response：

```
200 OK

Content-Type: application/json

{
                { "type": <type>, 
                "group_id": <group_id>, 
                "state": writable,
                "capactiy": 64GiB,
                "writable": 32GiB,
                "extent_count": 2, 
                "extents":
                    ["status":
                    {seq":1,
                    "node_id":<node_id>,
                    "disk_id":<disk_id>,
                    "state": writable
                    },
                    {seq":2,
                     "node_id":<node_id>,
                     "disk_id":<disk_id>,
                     "state": writable
                    }]
                }
}

```

# Job

## Get job/jobs information

### Get all jobs information

Request:

```
GET /kp/api/v1/job
```

Response:

```
{
    "jobs":[
        "job":{"job_id": <job_id>,
                "job_type": <type>,
                "process": {
                    "group_id":<group_id>,
                    "extent_seq":<extent_seq><
                    "from":{"node_id":<node_id>, "disk_id":<disk_id>},
                    "to":{"node_id":<node_id>, "disk_id":<disk_id>},
                    "seg_done":<seg_done>,
                    "seg_undone":<seg_undone>,
                    "seg_size":<seg_size>
                },
            "status": <status>,
            "msg": <msg>
              },
        ...
    ]
}

```

### Get a specific job.

Request:

```
GET /kp/api/v1/job/<job_id>
```

Response:

```
{
    "job_id": <job_id>,
    "job_type": <type>,
    "process": {
        "group_id":<group_id>,
        "extent_seq":<extent_seq><
        "from":{"node_id":<node_id>, "disk_id":<disk_id>},
        "to":{"node_id":<node_id>, "disk_id":<disk_id>},
        "seg_done":<seg_done>,
        "seg_undone":<seg_undone>,
        "seg_size":<seg_size>
    },
    "status": <status>,
    "msg": <msg>
}

```

## Control

### Start

Request:

```
PUT /kp/api/v1/job/start/<job_id>
```

Response:

```
200 OK
```

### Pause

Request:

```
PUT /kp/api/v1/job/pause/<job_id>
```

Response:

```
200 OK
```

### Terminate

Request:

```
PUT /kp/api/v1/job/terminate/<job_id>
```

Response:

```
200 OK
```
