#! /bin/bash

# exit when any command fails
set -e

ENV_FOLDER="$(pwd)/prototype/.env"
ENV_ACTIVATE_SCRIPT="${ENV_FOLDER}/bin/activate"

cd ./prototype

if [ ! -f "$ENV_ACTIVATE_SCRIPT" ]; then
    python3 -m venv "${ENV_FOLDER}"
fi

. "$ENV_ACTIVATE_SCRIPT"

mkdir -p ../production-schema
python ./generate_schema.py -o ../production-schema/

deactivate

cd ..

docker run --rm -v "${PWD}/production-schema:/local" --security-opt label=disable openapitools/openapi-generator-cli generate \
    -i /local/prototype_openapi.schema.json \
    -g rust-server \
    -o /local/out/rust-hyper