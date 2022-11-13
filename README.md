# http-server-demo

Messing around with http server frameworks.

## Prototype server

Written using FastAPI python http framework.

Generate a stub to begin production server by running 

```
./generate_production_server.sh
```

Schema files will generate to `production-schema/prototype_json_schema.schema.json` and `production-schema/prototype_openapi.schema.json`.

A stub Rust server as well as an example client will be written to `production-schema/out/rust-hyper`.
