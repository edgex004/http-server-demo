#! /bin/bash

# exit when any command fails
set -e
set -x

ENV_FOLDER="$(pwd)/prototype/.env"
ENV_ACTIVATE_SCRIPT="${ENV_FOLDER}/bin/activate"
ROOT_DIR="$(git rev-parse --show-toplevel)"

cd "${ROOT_DIR}/prototype"

if [ ! -f "$ENV_ACTIVATE_SCRIPT" ]; then
    python3 -m venv "${ENV_FOLDER}"
fi

. "$ENV_ACTIVATE_SCRIPT"

pip install -r ./requirements.txt
mkdir -p "${ROOT_DIR}/production-schema"
python ./generate_schema.py -o "${ROOT_DIR}/production-schema"

deactivate

cd "${ROOT_DIR}"

docker run --rm -v "${ROOT_DIR}/production-schema:/local" --security-opt label=disable openapitools/openapi-generator-cli generate \
    -i /local/prototype_openapi.schema.json \
    -g rust-server \
    -o /local/out/rust-hyper

cd "${ROOT_DIR}/production-schema" && npm install && npm run generate-go-types

cd "${ROOT_DIR}"

cp "${ROOT_DIR}/production-schema/out/logging_schema.go" "${ROOT_DIR}/robot-client/logging/logging_schema.go"
