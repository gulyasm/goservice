# Example Go service

## Build
To build the service:
```
make
```

## Test
To run tests:
```
make test
```
To get test coverage:
```
make coverage-test
```

## Running
The service uses environmental variables for configuration. To build and run:
```
make && LISTENING_ADDRESS=:12500 HEARTBEAT_ADDRESS=:12501 ./goservice
```

To test the running service:
```
curl localhost:12500/service/ip?domain=enbrite.ly
```

To test the heartbeat service:
```
curl localhost:12501/heartbeat
```

## UI
The UI is available at [http://localhost:12500](http://localhost:12500)


