# http-server-demo

Messing around with http server frameworks.

## Prototype server

Written using FastAPI python http framework.

See [the prototype readme](prototype/README.md) for usage.


## Robot Client

Robot test client is written in golang. It runs utilizing generated code from [the code generators](#code-generators).

See [the robot readme](robot-client/README.md) for usage.

## Code Generators

All code generators can be run using a top level script: 

```
./generate_code.sh
```

this will generate the following:

1. Robot alarms table accessors

    This is generated from `robot-client/alarms/alarm_info.csv` and creates an interface for the csv data in `robot-client/alarms/alarm_info.go`.

2. Robot client-server interface code

    Using the json schema types from `./prototype` server code, golang bindings are generated and placed at `robot-client/logging/logging_schema.go`.

3. Production server stub

    Schema files will generate to `production-schema/prototype_json_schema.schema.json` and `production-schema/prototype_openapi.schema.json`.

    A stub Rust server as well as an example client will be written to `production-schema/out/rust-hyper`.
