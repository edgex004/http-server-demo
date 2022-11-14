# http-server-demo

Messing around with http server frameworks.

## Prototype server

Written using FastAPI python http framework.

See [the prototype readme](prototype/README.md) for usage.

## Production server stub

Generate a stub to begin writing a production server by running 

```
./generate_production_server.sh
```

Schema files will generate to `production-schema/prototype_json_schema.schema.json` and `production-schema/prototype_openapi.schema.json`.

A stub Rust server as well as an example client will be written to `production-schema/out/rust-hyper`.
