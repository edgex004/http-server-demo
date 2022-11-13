# Prototype for an http server

Must be easy to read/understand and maintain. After prototyping work is complete, a schema should be able to be generated which can serve as a contract between client(s) and server.

This initial prototype utilizes [FastAPI](https://fastapi.tiangolo.com/).

## Dependancies

Recommend first creating and activating a virtual environment:

```
python -m venv ./.env && . ./.env/bin/activate
```

Then install dependancies:

```
pip install -r ./requirements.txt
```

## Running the prototype server

Launch the server with:

```
uvicorn main:app --reload
```

Reach auto-generated documentation (generated using [Swagger](https://swagger.io/)) by pointing a browser to:

```
http://127.0.0.1:8000/docs#/
```

Swagger's hosted documentation publishes what the expectations of the server are (request/response types) as well as hosts a UI that allows potential client developers to easily craft and run test cases against the live server.

## Generating OpenAPI schema for the server

The OpenAPI schema serves as a contract between the server and client. This is an important step when server and client will be developed in parallel. Not only does it serve to scope the work for server and client, but it can also be used to automatically generate client code, server stubs, and type parsers for a variety of languages using [Quicktype](https://quicktype.io/).

To generate a schema for the server:

```
python generate_schema.py
```

