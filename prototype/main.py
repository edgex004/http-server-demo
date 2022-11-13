from typing import Union

from fastapi import FastAPI, Query
from pydantic import BaseModel

description = """
Test server to store and retrieve objects from volitile memory.

## Items

You can:
* **create items**.
* **read items**.
"""

tags_metadata = [
    {
        "name": "items",
        "description": "Manage items.",
    },
]

app = FastAPI(
    version="0.0.1",
    description=description,
    openapi_tags=tags_metadata
    )

Items = {}

class Item(BaseModel):
    name: str
    value: int

class ItemResponse(BaseModel):
    id: int
    item: Item

@app.get("/items/", response_model=list[ItemResponse], tags=["items"])
def read_item(item_ids: Union[list[int], None] = Query(default=None)):
    ret = []
    if item_ids:
        for item_id in item_ids:
            if item_id in Items:
                ret.append({"id": item_id, "item":Items[item_id]})
    return ret

@app.put("/items/{item_id}", response_model=ItemResponse, tags=["items"])
def update_item(item_id: int, item: Item):
    Items[item_id] = item
    return {"id": item_id, "item":Items[item_id]}
