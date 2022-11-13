import argparse
import os

parser = argparse.ArgumentParser(
                    prog = 'Generate Schema',
                    description = 'Generate JSON Schema from prototype server')

parser.add_argument('-o', '--output_folder', required=True)
args = parser.parse_args()
open_api_schema_file = os.path.join( args.output_folder, "prototype_openapi.schema.json")
json_schema_file = os.path.join( args.output_folder, "prototype_json_schema.schema.json")

import main
import json

open_api_dict = main.app.openapi()

f = open(open_api_schema_file, "w")
stringified_schema=json.dumps(open_api_dict, indent=2)
f.write(stringified_schema)
f.close()


# Define a recursive function to walk all dictionary entries and redirect references from "components/schemas" to "$defs"
def redirect_references(x: object, max_depth: int = 50):
    if isinstance(x, list):
        for value in x:
            redirect_references(value,max_depth-1)
    elif isinstance(x, dict):
        for attr, value in x.items():
            if isinstance(value, str):
                if attr == "$ref":
                    x[attr]=value.replace("components/schemas","$defs")
            else:
                redirect_references(value,max_depth-1)



if "components" in open_api_dict and "schemas" in open_api_dict["components"]:
    json_schema_dict = {
        "description": "Container for prototype schema definitions.",
        "type": "object"
        }
    json_schema_dict["$defs"] = open_api_dict["components"]["schemas"]
    json_schema_dict["properties"] = {}
    for attr, value in json_schema_dict["$defs"].items():
        json_schema_dict["properties"][attr] = {"$ref":"#/$defs/" + attr}
    redirect_references(json_schema_dict["$defs"])
    f = open(json_schema_file, "w")
    stringified_schema = json.dumps(json_schema_dict, indent=2)
    f.write(stringified_schema)
    f.close()