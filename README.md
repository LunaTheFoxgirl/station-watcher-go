# station-watcher-go
A single-purpose application that monitors ShoutCast and IceCast instances and converts changes in song or listener data into URI webhooks. 

## Running in Docker

```
docker-compose build
docker-compose run go go-wrapper install
```

### Example URIs

#### `shoutcast2` adapter
```
http://docker.local:8000/statistics?json=1
```

#### `icecast` adapter
```
http://admin:password@docker.local:8000/admin/stats
```

## About this project

This application is a single-purpose utility script used to convert changes in Shoutcast or Icecast radio metadata (namely, listeners or the currently playing song) into webhook notifications, allowing for metadata updates that much more closely match the actual changes in songs than is otherwise possible.

This application is built to be a part of the [AzuraCast](https://github.com/AzuraCast/AzuraCast) radio management suite.
